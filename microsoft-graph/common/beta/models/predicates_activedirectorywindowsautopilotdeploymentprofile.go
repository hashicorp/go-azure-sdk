package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryWindowsAutopilotDeploymentProfileOperationPredicate struct {
	CreatedDateTime                        *string
	Description                            *string
	DeviceNameTemplate                     *string
	DisplayName                            *string
	EnableWhiteGlove                       *bool
	ExtractHardwareHash                    *bool
	HybridAzureADJoinSkipConnectivityCheck *bool
	Id                                     *string
	Language                               *string
	LastModifiedDateTime                   *string
	ManagementServiceAppId                 *string
	ODataType                              *string
}

func (p ActiveDirectoryWindowsAutopilotDeploymentProfileOperationPredicate) Matches(input ActiveDirectoryWindowsAutopilotDeploymentProfile) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceNameTemplate != nil && (input.DeviceNameTemplate == nil || *p.DeviceNameTemplate != *input.DeviceNameTemplate) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnableWhiteGlove != nil && (input.EnableWhiteGlove == nil || *p.EnableWhiteGlove != *input.EnableWhiteGlove) {
		return false
	}

	if p.ExtractHardwareHash != nil && (input.ExtractHardwareHash == nil || *p.ExtractHardwareHash != *input.ExtractHardwareHash) {
		return false
	}

	if p.HybridAzureADJoinSkipConnectivityCheck != nil && (input.HybridAzureADJoinSkipConnectivityCheck == nil || *p.HybridAzureADJoinSkipConnectivityCheck != *input.HybridAzureADJoinSkipConnectivityCheck) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Language != nil && (input.Language == nil || *p.Language != *input.Language) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ManagementServiceAppId != nil && (input.ManagementServiceAppId == nil || *p.ManagementServiceAppId != *input.ManagementServiceAppId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
