package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceSettingsOperationPredicate struct {
	AllowSessionsToUnenrolledDevices *bool
	BlockChat                        *bool
	Id                               *string
	ODataType                        *string
}

func (p RemoteAssistanceSettingsOperationPredicate) Matches(input RemoteAssistanceSettings) bool {

	if p.AllowSessionsToUnenrolledDevices != nil && (input.AllowSessionsToUnenrolledDevices == nil || *p.AllowSessionsToUnenrolledDevices != *input.AllowSessionsToUnenrolledDevices) {
		return false
	}

	if p.BlockChat != nil && (input.BlockChat == nil || *p.BlockChat != *input.BlockChat) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
