package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsOperationPredicate struct {
	BlockDataRecoveryAgent                         *bool
	EnableBitLockerAfterRecoveryInformationToStore *bool
	EnableRecoveryInformationSaveToStore           *bool
	HideRecoveryOptions                            *bool
	ODataType                                      *string
}

func (p BitLockerRecoveryOptionsOperationPredicate) Matches(input BitLockerRecoveryOptions) bool {

	if p.BlockDataRecoveryAgent != nil && (input.BlockDataRecoveryAgent == nil || *p.BlockDataRecoveryAgent != *input.BlockDataRecoveryAgent) {
		return false
	}

	if p.EnableBitLockerAfterRecoveryInformationToStore != nil && (input.EnableBitLockerAfterRecoveryInformationToStore == nil || *p.EnableBitLockerAfterRecoveryInformationToStore != *input.EnableBitLockerAfterRecoveryInformationToStore) {
		return false
	}

	if p.EnableRecoveryInformationSaveToStore != nil && (input.EnableRecoveryInformationSaveToStore == nil || *p.EnableRecoveryInformationSaveToStore != *input.EnableRecoveryInformationSaveToStore) {
		return false
	}

	if p.HideRecoveryOptions != nil && (input.HideRecoveryOptions == nil || *p.HideRecoveryOptions != *input.HideRecoveryOptions) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
