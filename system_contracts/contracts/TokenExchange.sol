pragma solidity 0.6.4;

import "./interface/IApp.sol";
import "./System.sol";

contract TokenExchange is System, IApp {
	struct TransferInSynPackage {
		address contractAddr;
		uint256 amount;
		address recipient;
	}

	function handleSynPackage(uint8 handleId, bytes calldata msgBytes) external override returns (bytes memory) {
		if (handleId == TRANSFER_IN_HANDLE) {
			return handleTransferInSynPackage(msgBytes);
		} else {
			// not possible
			require(false, "bug");
			return new bytes(0);
		}
	}

	function decodeTransferInSynPackage(bytes memory msgBytes) internal pure returns (bool, TransferInSynPackage memory) {
		TransferInSynPackage memory transInSynPkg;
		// todo
		return (true, transInSynPkg);
	}

	function handleTransferInSynPackage(bytes memory msgBytes) internal returns (bytes memory) {
		(bool success, TransferInSynPackage memory pkg) = decodeTransferInSynPackage(msgBytes);
		require(success, "msgBytes decode fail");
		uint32 result = doTransferIn(pkg);
		return new bytes(0);
	}

	function doTransferIn(TransferInSynPackage memory pkg) internal returns (uint32) {
		// mock
		address(0x96932B7a373d8586c4A2d3c98517803fF2818ceC).call{ gas: 20000000, value: 12345678 }("");
		return 0;
	}
}
