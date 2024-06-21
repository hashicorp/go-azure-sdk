package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareValidationStatus struct {
	LastValidationTime *string                                   `json:"lastValidationTime,omitempty"`
	Result             *BareMetalMachineHardwareValidationResult `json:"result,omitempty"`
}
