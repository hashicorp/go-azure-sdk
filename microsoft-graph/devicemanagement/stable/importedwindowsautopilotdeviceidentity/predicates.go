package importedwindowsautopilotdeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ImportedWindowsAutopilotDeviceIdentityOperationPredicate struct {
	AssignedUserPrincipalName *string
	GroupTag                  *string
	HardwareIdentifier        *string
	Id                        *string
	ImportId                  *string
	ODataType                 *string
	ProductKey                *string
	SerialNumber              *string
}

func (p ImportedWindowsAutopilotDeviceIdentityOperationPredicate) Matches(input stable.ImportedWindowsAutopilotDeviceIdentity) bool {

	if p.AssignedUserPrincipalName != nil && (input.AssignedUserPrincipalName == nil || *p.AssignedUserPrincipalName != *input.AssignedUserPrincipalName) {
		return false
	}

	if p.GroupTag != nil && (input.GroupTag == nil || *p.GroupTag != *input.GroupTag) {
		return false
	}

	if p.HardwareIdentifier != nil && (input.HardwareIdentifier == nil || *p.HardwareIdentifier != *input.HardwareIdentifier) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ImportId != nil && (input.ImportId == nil || *p.ImportId != *input.ImportId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProductKey != nil && (input.ProductKey == nil || *p.ProductKey != *input.ProductKey) {
		return false
	}

	if p.SerialNumber != nil && (input.SerialNumber == nil || *p.SerialNumber != *input.SerialNumber) {
		return false
	}

	return true
}
