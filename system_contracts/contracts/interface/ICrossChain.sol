pragma solidity 0.6.4;

interface ICrossChain {
    function sendCCPackage(
        uint8 handle_id,
        bytes calldata msgBytes,
        uint256 fee
    ) external;
}
