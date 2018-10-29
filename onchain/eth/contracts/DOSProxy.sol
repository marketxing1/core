pragma solidity ^0.4.24;
// Not enabled for production yet.
//pragma experimental ABIEncoderV2;

library BN256 {
    struct G1Point {
        uint x;
        uint y;
    }

    struct G2Point {
        uint[2] x;
        uint[2] y;
    }

    function P1() internal pure returns (G1Point) {
        return G1Point(1, 2);
    }

    function P2() internal pure returns (G2Point) {
        return G2Point(
            [11559732032986387107991004021392285783925812861821192530917403151452391805634,
            10857046999023057135944570762232829481370756359578518086990519993285655852781],

            [4082367875863433681332203403145435568316851327593401208105741076214120093531,
            8495653923123431417604973247489272438418190587263600148770280649306958101930]
        );
    }

    function pointAdd(G1Point p1, G1Point p2) internal returns (G1Point r) {
        uint[4] memory input;
        input[0] = p1.x;
        input[1] = p1.y;
        input[2] = p2.x;
        input[3] = p2.y;
        assembly {
            if iszero(call(sub(gas, 2000), 0x6, 0, input, 0x80, r, 0x40)) {
                revert(0, 0)
            }
        }
    }

    function scalarMul(G1Point p, uint s) internal returns (G1Point r) {
        uint[3] memory input;
        input[0] = p.x;
        input[1] = p.y;
        input[2] = s;
        assembly {
            if iszero(call(sub(gas, 2000), 0x7, 0, input, 0x60, r, 0x40)) {
                revert(0, 0)
            }
        }
    }

    function hashToG1(bytes data) internal returns (G1Point) {
        uint256 h = uint256(keccak256(data));
        return scalarMul(P1(), h);
    }

    // @return the result of computing the pairing check
    // check passes if e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
    function pairingCheck(G1Point[] p1, G2Point[] p2) internal returns (bool) {
        require(p1.length == p2.length);
        uint elements = p1.length;
        uint inputSize = elements * 6;
        uint[] memory input = new uint[](inputSize);

        for (uint i = 0; i < elements; i++)
        {
            input[i * 6 + 0] = p1[i].x;
            input[i * 6 + 1] = p1[i].y;
            input[i * 6 + 2] = p2[i].x[0];
            input[i * 6 + 3] = p2[i].x[1];
            input[i * 6 + 4] = p2[i].y[0];
            input[i * 6 + 5] = p2[i].y[1];
        }

        uint[1] memory out;
        bool success;
        assembly {
            success := call(
                sub(gas, 2000),
                0x8,
                0,
                add(input, 0x20),
                mul(inputSize, 0x20),
                out, 0x20
            )
        }
        return success && (out[0] != 0);
    }
}

interface UserContractInterface {
    // Query callback.
    function __callback__(uint, bytes) external;
    // Random number callback.
    function __callback__(uint, uint) external;
}

contract DOSProxy {
    using BN256 for *;

    struct PendingRequest {
        uint requestId;
        BN256.G2Point handledGroup;
        // User contract issued the query.
        address callbackAddr;
    }

    uint requestIdSeed;
    uint groupSize;
    uint[] nodeId;
    // calling requestId => PendingQuery metadata
    mapping(uint => PendingRequest) PendingRequests;
    // Note: Make atomic changes to group metadata below.
    BN256.G2Point[] groupPubKeys;
    // groupIdentifier => isExisted
    mapping(bytes32 => bool) groups;
    //publicKey => publicKey appearance
    mapping(bytes32 => uint) pubKeyCounter;
    // Note: Make atomic changes to randomness metadata below.
    uint public lastUpdatedBlock;
    uint public lastRandomness;
    BN256.G2Point lastHandledGroup;
    uint8 constant TrafficSystemRandom = 0;
    uint8 constant TrafficUserRandom = 1;
    uint8 constant TrafficUserQuery = 2;

    event LogUrl(
        uint queryId,
        string url,
        uint timeout,
        uint randomness,
        // Log G2Point struct directly is an experimental feature, use with care.
        uint[4] dispatchedGroup
    );
    event LogRequestUserRandom(
        uint requestId,
        uint lastSystemRandomness,
        uint userSeed,
        uint[4] dispatchedGroup
    );
    event LogNonSupportedType(string queryType);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address callbackAddr);
    event LogRequestFromNonExistentUC();
    event LogUpdateRandom(uint lastRandomness, uint[4] dispatchedGroup);
    event LogValidationResult(
        uint8 trafficType,
        uint trafficId,
        bytes message,
        uint[2] signature,
        uint[4] pubKey,
        bool pass
    );
    event LogInsufficientGroupNumber();
    event LogGrouping(uint[] NodeId);
    event LogPublicKeyAccepted(uint x1, uint x2, uint y1, uint y2);

    // whitelist state variables used only for alpha release.
    // Index starting from 1.
    address[22] whitelists;
    // whitelisted address => index in whitelists.
    mapping(address => uint) isWhitelisted;
    bool public whitelistInited = false;
    event WhitelistAddressTransferred(address previous, address curr);

    modifier onlyWhitelisted {
        uint idx = isWhitelisted[msg.sender];
        require(idx != 0 && whitelists[idx] == msg.sender, "Not whitelisted!");
        _;
    }

    function initWhitelist(address[21] addresses) public {
        require(!whitelistInited, "Whitelist already initialized!");

        for (uint idx = 0; idx < 21; idx++) {
            whitelists[idx+1] = addresses[idx];
            isWhitelisted[addresses[idx]] = idx+1;
        }
        whitelistInited = true;
    }

    function getWhitelistAddess(uint idx) public view returns (address) {
        require(idx > 0 && idx <= 21, "Index out of range");
        return whitelists[idx];
    }

    function transferWhitelistAddress(address newWhitelistedAddr)
        public
        onlyWhitelisted
    {
        require(newWhitelistedAddr != 0x0 && newWhitelistedAddr != msg.sender);

        emit WhitelistAddressTransferred(msg.sender, newWhitelistedAddr);
        whitelists[isWhitelisted[msg.sender]] = newWhitelistedAddr;
    }


    function getCodeSize(address addr) internal constant returns (uint size) {
        assembly {
            size := extcodesize(addr)
        }
    }

    function strEqual(string a, string b) internal pure returns (bool) {
        bytes memory aBytes = bytes(a);
        bytes memory bBytes = bytes(b);
        if (aBytes.length != bBytes.length) {
            return false;
        }
        for(uint i = 0; i < aBytes.length; i++) {
            if (aBytes[i] != bBytes[i]) {
                return false;
            }
        }
        return true;
    }

    // Returns query id.
    // TODO: restrict query from subscribed/paid calling contracts.
    function query(
        address from,
        uint timeout,
        string queryType,
        string queryPath
    )
        external
        returns (uint)
    {
        if (getCodeSize(from) > 0) {
            // Only supporting api/url requests for alpha release.
            if (strEqual(queryType, 'API')) {
                uint queryId = uint(keccak256(abi.encodePacked(
                    ++requestIdSeed, from, timeout, queryType, queryPath)));
                uint idx = lastRandomness % groupPubKeys.length;
                PendingRequests[queryId] =
                    PendingRequest(queryId, groupPubKeys[idx], from);
                emit LogUrl(
                    queryId,
                    queryPath,
                    timeout,
                    lastRandomness,
                    getGroupPubKey(idx)
                );
                return queryId;
            } else {
                emit LogNonSupportedType(queryType);
                return 0x0;
            }
        } else {
            // Skip if @from is not contract address.
            emit LogNonContractCall(from);
            return 0x0;
        }
    }

    // Request a new user-level random number.
    function requestRandom(address from, uint8 mode, uint userSeed)
        external
        returns (uint)
    {
        // fast mode
        if (mode == 0) {
            return uint(keccak256(abi.encodePacked(
                ++requestIdSeed,lastRandomness, userSeed)));
        } else if (mode == 1) {
            // safe mode
            // TODO: restrict request from paid calling contract address.
            uint requestId = uint(keccak256(abi.encodePacked(
                ++requestIdSeed, from, userSeed)));
            uint idx = lastRandomness % groupPubKeys.length;
            PendingRequests[requestId] =
                PendingRequest(requestId, groupPubKeys[idx], from);
            // sign(requestId ||lastSystemRandomness || userSeed) with
            // selected group
            emit LogRequestUserRandom(
                requestId,
                lastRandomness,
                userSeed,
                getGroupPubKey(idx)
            );
            return requestId;
        } else {
            revert("Non-supported random request");
        }
    }

    // Random submitter validation + group signature verification.
    function validateAndVerify(
        uint8 trafficType,
        uint trafficId,
        bytes data,
        BN256.G1Point signature,
        BN256.G2Point grpPubKey
    )
        internal
        onlyWhitelisted
        returns (bool)
    {
        // Validation
        // TODO
        // 1. Check msg.sender from registered and staked node operator.
        // 2. Check msg.sender is a member in Group(grpPubKey).
        // Clients actually signs (data || addr(selected_submitter)).
        // TODO: Sync and change to sign ( sha256(data) || address )
        bytes memory message = abi.encodePacked(data, msg.sender);

        // Verification
        BN256.G1Point[] memory p1 = new BN256.G1Point[](2);
        BN256.G2Point[] memory p2 = new BN256.G2Point[](2);
        // The signature has already been applied neg() function offchainly to
        // fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = BN256.hashToG1(message);
        p2[0] = BN256.P2();
        p2[1] = grpPubKey;
        bool passVerify = BN256.pairingCheck(p1, p2);
        emit LogValidationResult(
            trafficType,
            trafficId,
            message,
            [signature.x, signature.y],
            [grpPubKey.x[0], grpPubKey.x[1], grpPubKey.y[0], grpPubKey.y[1]],
            passVerify
        );
        return passVerify;
    }

    function triggerCallback(
        uint requestId,
        uint8 trafficType,
        bytes result,
        uint[2] sig
    )
        external
    {
        if (!validateAndVerify(
                trafficType,
                requestId,
                result,
                BN256.G1Point(sig[0], sig[1]),
                PendingRequests[requestId].handledGroup))
        {
            return;
        }

        address ucAddr = PendingRequests[requestId].callbackAddr;
        if (ucAddr == 0x0) {
            emit LogRequestFromNonExistentUC();
            return;
        }

        emit LogCallbackTriggeredFor(ucAddr);
        delete PendingRequests[requestId];
        if (trafficType == TrafficUserQuery) {
            UserContractInterface(ucAddr).__callback__(requestId, result);
        } else if (trafficType == TrafficUserRandom) {
            // Safe random number is the collectively signed threshold signature
            // of the message (requestId || lastRandomness || userSeed ||
            // selected sender in group).
            UserContractInterface(ucAddr).__callback__(
                requestId, uint(keccak256(abi.encodePacked(sig[0], sig[1]))));
        } else {
            revert("Unsupported traffic type");
        }
    }

    function toBytes(uint x) internal pure returns (bytes b) {
        b = new bytes(32);
        assembly { mstore(add(b, 32), x) }
    }

    // System-level secure distributed random number generator.
    function updateRandomness(uint[2] sig) external {
        if (!validateAndVerify(
                TrafficSystemRandom,
                lastRandomness,
                toBytes(lastRandomness),
                BN256.G1Point(sig[0], sig[1]),
                lastHandledGroup))
        {
            return;
        }
        // Update new randomness = sha3(collectively signed group signature)
        lastRandomness = uint(keccak256(abi.encodePacked(sig[0], sig[1])));
        lastUpdatedBlock = block.number - 1;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal selected off-chain clients to collectively generate a new
        // system level random number for next round.
        emit LogUpdateRandom(lastRandomness, getGroupPubKey(idx));
    }

    // For alpha. To trigger first random number after grouping has done
    // or timeout.
    function fireRandom() public onlyWhitelisted {
        lastRandomness = uint(keccak256(abi.encode(blockhash(block.number - 1))));
        lastUpdatedBlock = block.number - 1;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, getGroupPubKey(idx));
    }

    function handleTimeout() public onlyWhitelisted {
        uint currentBlockNumber = block.number - 1;
        if (currentBlockNumber - lastUpdatedBlock > 5) {
            fireRandom();
        }
    }

    function setPublicKey(uint x1, uint x2, uint y1, uint y2)
        public
        onlyWhitelisted
    {
        bytes32 groupId = keccak256(abi.encodePacked(x1, x2, y1, y2));
        require(!groups[groupId], "group has already registered");

        pubKeyCounter[groupId] = pubKeyCounter[groupId] + 1;
        if (pubKeyCounter[groupId] > groupSize / 2) {
            groupPubKeys.push(BN256.G2Point([x1, x2], [y1, y2]));
            groups[groupId] = true;
            delete(pubKeyCounter[groupId]);
            emit LogPublicKeyAccepted(x1, x2, y1, y2);
        }
    }

    function getGroupPubKey(uint idx) public constant returns (uint[4]) {
        require(idx < groupPubKeys.length, "group index out of range");

        return [
            groupPubKeys[idx].x[0], groupPubKeys[idx].x[1],
            groupPubKeys[idx].y[0], groupPubKeys[idx].y[1]
        ];
    }

    function uploadNodeId(uint id) public onlyWhitelisted {
        nodeId.push(id);
    }

    function grouping(uint size) public onlyWhitelisted {
        groupSize = size;
        uint[] memory toBeGrouped = new uint[](size);
        if (nodeId.length < size) {
            emit LogInsufficientGroupNumber();
            return;
        }
        for (uint i = 0; i < size; i++) {
            toBeGrouped[i] = nodeId[nodeId.length - 1];
            nodeId.length--;
        }
        emit LogGrouping(toBeGrouped);
    }

    function resetContract() public onlyWhitelisted {
        nodeId.length = 0;
        groupPubKeys.length = 0;
    }
}
