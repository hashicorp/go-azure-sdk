package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptions struct {
	BlockDataRecoveryAgent                         *bool                                               `json:"blockDataRecoveryAgent,omitempty"`
	EnableBitLockerAfterRecoveryInformationToStore *bool                                               `json:"enableBitLockerAfterRecoveryInformationToStore,omitempty"`
	EnableRecoveryInformationSaveToStore           *bool                                               `json:"enableRecoveryInformationSaveToStore,omitempty"`
	HideRecoveryOptions                            *bool                                               `json:"hideRecoveryOptions,omitempty"`
	ODataType                                      *string                                             `json:"@odata.type,omitempty"`
	RecoveryInformationToStore                     *BitLockerRecoveryOptionsRecoveryInformationToStore `json:"recoveryInformationToStore,omitempty"`
	RecoveryKeyUsage                               *BitLockerRecoveryOptionsRecoveryKeyUsage           `json:"recoveryKeyUsage,omitempty"`
	RecoveryPasswordUsage                          *BitLockerRecoveryOptionsRecoveryPasswordUsage      `json:"recoveryPasswordUsage,omitempty"`
}
