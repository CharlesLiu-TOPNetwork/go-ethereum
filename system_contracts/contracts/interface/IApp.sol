pragma solidity 0.6.4;

interface IApp {
	function handleSynPackage(uint8 handleId, bytes calldata msgBytes) external returns (bytes memory responsePayload);
}
