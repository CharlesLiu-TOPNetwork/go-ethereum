pragma solidity 0.6.4;

library Memory {
	// Size of a word, in bytes.
	uint256 internal constant WORD_SIZE = 32;

	// Copy 'len' bytes from memory address 'src', to address 'dest'.
	// This function does not check the or destination, it only copies
	// the bytes.
	function copy(
		uint256 src,
		uint256 dest,
		uint256 len
	) internal pure {
		// Copy word-length chunks while possible
		for (; len >= WORD_SIZE; len -= WORD_SIZE) {
			assembly {
				mstore(dest, mload(src))
			}
			dest += WORD_SIZE;
			src += WORD_SIZE;
		}

		// Copy remaining bytes
		uint256 mask = 256**(WORD_SIZE - len) - 1;
		assembly {
			let srcpart := and(mload(src), not(mask))
			let destpart := and(mload(dest), mask)
			mstore(dest, or(destpart, srcpart))
		}
	}

	// Creates a 'bytes memory' variable from the memory address 'addr', with the
	// length 'len'. The function will allocate new memory for the bytes array, and
	// the 'len bytes starting at 'addr' will be copied into that new memory.
	function fromBytes(bytes memory bts) internal pure returns (uint256 addr, uint256 len) {
		len = bts.length;
		assembly {
			/*BYTES_HEADER_SIZE*/
			addr := add(bts, 32)
		}
	}
}
