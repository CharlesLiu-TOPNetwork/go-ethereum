pragma solidity 0.6.4;

import "./System.sol";
import "./lib/Memory.sol";
import "./interface/IApp.sol";

contract CrossChain is System {
	uint8 public constant SYN_PACKAGE = 0x00;
	uint8 public constant ACK_PACKAGE = 0x01;
	uint8 public constant FAIL_ACK_PACKAGE = 0x02;

	mapping(uint8 => address) public handleContractMap;

	event packageDecodeError(uint64 indexed packageSequence, uint8 indexed handleId, bytes payload);
	event receivePackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed handleId);

	function init() external onlyNotInit {
		handleContractMap[TRANSFER_IN_HANDLE] = TOKEN_EXCHANGE_CONTRACT_ADDR;

		alreadyInit = true;
	}

	// |- type -|- relayFee -|- package -|
	// | 1 byte |  32 bytes  |  bytes... |
	function decodePayload(bytes memory payload)
		internal
		pure
		returns (
			bool,
			uint8,
			uint256,
			bytes memory
		)
	{
		// mock
		return (true, SYN_PACKAGE, 0, new bytes(0));
		if (payload.length < 33) {
			return (false, 0, 0, new bytes(0));
		}

		uint256 ptr;
		assembly {
			ptr := payload
		}

		uint8 packageType;
		ptr += 1;
		assembly {
			packageType := mload(ptr)
		}

		uint256 relayFee;
		ptr += 32;
		assembly {
			relayFee := mload(ptr)
		}

		ptr += 32;

		bytes memory msgBytes = new bytes(payload.length - 33);
		(uint256 dst, ) = Memory.fromBytes(msgBytes);
		Memory.copy(ptr, dst, payload.length - 33);

		return (true, packageType, relayFee, msgBytes);
	}

	function handleCCPackage(
		bytes calldata payload,
		// bytes calldata proof,
		// uint64 height,
		uint64 packageSequence,
		uint8 handleId
	) external onlyInit {
		// todo requires

		// bytes memory payloadLocal

		(bool success, uint8 packageType, uint256 relayFee, bytes memory msgBytes) = decodePayload(payload);

		if (!success) {
			emit packageDecodeError(packageSequence, handleId, payload);
			return;
		}
		emit receivePackage(packageType, packageSequence, handleId);

		if (packageType == SYN_PACKAGE) {
			address handleContract = handleContractMap[handleId];
			try IApp(handleContract).handleSynPackage(handleId, msgBytes) returns (bytes memory responsePayload) {
				// emit
				// send ack
			} catch Error(string memory reason) {
				// emit
				// send failed ack
			}
		}
	}
}
