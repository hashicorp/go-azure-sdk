package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMDeviceStateOperationPredicate struct {
	CreatedDateTime                          *string
	DeviceName                               *string
	Id                                       *string
	LastSyncDateTime                         *string
	ModifiedDateTime                         *string
	ODataType                                *string
	StateDetails                             *string
	UniversalIntegratedCircuitCardIdentifier *string
	UserName                                 *string
}

func (p EmbeddedSIMDeviceStateOperationPredicate) Matches(input EmbeddedSIMDeviceState) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StateDetails != nil && (input.StateDetails == nil || *p.StateDetails != *input.StateDetails) {
		return false
	}

	if p.UniversalIntegratedCircuitCardIdentifier != nil && (input.UniversalIntegratedCircuitCardIdentifier == nil || *p.UniversalIntegratedCircuitCardIdentifier != *input.UniversalIntegratedCircuitCardIdentifier) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
