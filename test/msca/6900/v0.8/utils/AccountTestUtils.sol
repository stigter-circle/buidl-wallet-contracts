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

import {TestUtils} from "../../../../util/TestUtils.sol";
import {ModuleEntity} from "../../../../../src/msca/6900/v0.8/common/Types.sol";
import {
    RESERVED_VALIDATION_DATA_INDEX,
    GLOBAL_VALIDATION_FLAG,
    PER_SELECTOR_VALIDATION_FLAG
} from "../../../../../src/msca/6900/v0.8/common/Constants.sol";
import {ModuleEntityLib} from "../../../../../src/msca/6900/v0.8/libs/thirdparty/ModuleEntityLib.sol";

contract AccountTestUtils is TestUtils {
    using ModuleEntityLib for ModuleEntity;

    struct PreValidationHookData {
        uint8 index;
        bytes hookData;
    }

    // @notice Helper function forked from 6900 to encode a signature, according to the per-hook and per-validation data
    // format.
    // @param validationData signature
    function encodeSignature(
        PreValidationHookData[] memory preValidationHookData,
        ModuleEntity validationFunction,
        bytes memory validationData,
        bool isValidationGlobal
    ) internal pure returns (bytes memory signature) {
        if (isValidationGlobal) {
            signature = abi.encodePacked(validationFunction, GLOBAL_VALIDATION_FLAG);
        } else {
            signature = abi.encodePacked(validationFunction, PER_SELECTOR_VALIDATION_FLAG);
        }
        for (uint256 i = 0; i < preValidationHookData.length; ++i) {
            signature = abi.encodePacked(
                signature, packSparseDataWithIndex(preValidationHookData[i].index, preValidationHookData[i].hookData)
            );
        }
        signature = abi.encodePacked(signature, packSparseDataWithIndex(RESERVED_VALIDATION_DATA_INDEX, validationData));
        return signature;
    }

    // @notice Helper function forked from 6900 to pack validation data with an index, according to the sparse calldata
    // segment spec.
    function packSparseDataWithIndex(uint8 index, bytes memory callData) internal pure returns (bytes memory) {
        return abi.encodePacked(uint32(callData.length + 1), index, callData);
    }
}