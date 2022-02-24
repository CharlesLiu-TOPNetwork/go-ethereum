pragma solidity 0.6.4;

contract System {
	bool public alreadyInit;

	uint8 public constant TRANSFER_IN_HANDLE = 0x01;
	uint8 public constant TRANSFER_OUT_HANDLE = 0x02;

	address public constant CROSS_CHAIN_CONTRACT_ADDR = 0x0000000000000000000000000000000000000100;
	address public constant TOKEN_EXCHANGE_CONTRACT_ADDR = 0x0000000000000000000000000000000000000200;

	modifier onlyInit() {
		require(alreadyInit, "Contract not init yet");
		_;
	}
	modifier onlyNotInit() {
		require(!alreadyInit, "Contract not init yet");
		_;
	}
}
