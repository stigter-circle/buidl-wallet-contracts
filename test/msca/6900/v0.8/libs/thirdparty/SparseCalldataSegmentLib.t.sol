/*
 * Copyright 2024 Circle Internet Group, Inc. All rights reserved.

 * SPDX-License-Identifier: GPL-3.0-or-later

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
pragma solidity 0.8.24;

import {AccountTestUtils} from "../../utils/AccountTestUtils.sol";
import {SparseCalldataSegmentLib} from
    "../../../../../../src/msca/6900/v0.8/libs/thirdparty/SparseCalldataSegmentLib.sol";

// TODO: remove this till SparseCalldataSegmentLib is moved to modular account lib
// https://github.com/erc6900/modular-account-libs/issues/2
contract SparseCalldataSegmentLibTest is AccountTestUtils {
    using SparseCalldataSegmentLib for bytes;

    function testFuzz_sparseCalldataSegmentLib_encodeDecode_simple(bytes[] memory segments) public {
        bytes memory encoded = _encodeSimple(segments);
        bytes[] memory decoded = this.decodeSimple(encoded, segments.length);

        assertEq(decoded.length, segments.length, "decoded.length != segments.length");

        for (uint256 i = 0; i < segments.length; i++) {
            assertEq(decoded[i], segments[i]);
        }
    }

    function testFuzz_sparseCalldataSegmentLib_encodeDecode_withIndex(bytes[] memory segments, uint256 indexSeed)
        public
    {
        // Generate random indices
        uint8[] memory indices = new uint8[](segments.length);
        for (uint256 i = 0; i < segments.length; i++) {
            uint8 nextIndex = uint8(uint256(keccak256(abi.encodePacked(indexSeed, i))));
            indices[i] = nextIndex;
        }

        // Encode
        bytes memory encoded = _encodeWithIndex(segments, indices);

        // Decode
        (bytes[] memory decodedBodies, uint8[] memory decodedIndices) = this.decodeWithIndex(encoded, segments.length);

        assertEq(decodedBodies.length, segments.length, "decodedBodies.length != segments.length");
        assertEq(decodedIndices.length, segments.length, "decodedIndices.length != segments.length");

        for (uint256 i = 0; i < segments.length; i++) {
            assertEq(decodedBodies[i], segments[i]);
            assertEq(decodedIndices[i], indices[i]);
        }
    }

    function _encodeSimple(bytes[] memory segments) internal pure returns (bytes memory) {
        bytes memory result = "";

        for (uint256 i = 0; i < segments.length; i++) {
            result = abi.encodePacked(result, uint32(segments[i].length), segments[i]);
        }

        return result;
    }

    function _encodeWithIndex(bytes[] memory segments, uint8[] memory indices) internal pure returns (bytes memory) {
        assertEq(segments.length, indices.length);

        bytes memory result = "";

        for (uint256 i = 0; i < segments.length; i++) {
            result = abi.encodePacked(result, uint32(segments[i].length + 1), indices[i], segments[i]);
        }

        return result;
    }

    function decodeSimple(bytes calldata encoded, uint256 capacityHint) external pure returns (bytes[] memory) {
        bytes[] memory result = new bytes[](capacityHint);

        bytes calldata remainder = encoded;

        uint256 index = 0;
        while (remainder.length > 0) {
            bytes calldata segment;
            (segment, remainder) = remainder.getNextSegment();
            result[index] = segment;
            index++;
        }

        return result;
    }

    function decodeWithIndex(bytes calldata encoded, uint256 capacityHint)
        external
        pure
        returns (bytes[] memory, uint8[] memory)
    {
        bytes[] memory bodies = new bytes[](capacityHint);
        uint8[] memory indices = new uint8[](capacityHint);

        bytes calldata remainder = encoded;

        uint256 index = 0;
        while (remainder.length > 0) {
            bytes calldata segment;
            (segment, remainder) = remainder.getNextSegment();
            bodies[index] = segment.getBody();
            indices[index] = segment.getIndex();
            index++;
        }

        return (bodies, indices);
    }
}