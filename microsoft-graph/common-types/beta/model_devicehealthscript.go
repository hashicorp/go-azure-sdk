package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScript struct {
	Assignments                 *[]DeviceHealthScriptAssignment           `json:"assignments,omitempty"`
	CreatedDateTime             *string                                   `json:"createdDateTime,omitempty"`
	Description                 *string                                   `json:"description,omitempty"`
	DetectionScriptContent      *string                                   `json:"detectionScriptContent,omitempty"`
	DetectionScriptParameters   *[]DeviceHealthScriptParameter            `json:"detectionScriptParameters,omitempty"`
	DeviceHealthScriptType      *DeviceHealthScriptDeviceHealthScriptType `json:"deviceHealthScriptType,omitempty"`
	DeviceRunStates             *[]DeviceHealthScriptDeviceState          `json:"deviceRunStates,omitempty"`
	DisplayName                 *string                                   `json:"displayName,omitempty"`
	EnforceSignatureCheck       *bool                                     `json:"enforceSignatureCheck,omitempty"`
	HighestAvailableVersion     *string                                   `json:"highestAvailableVersion,omitempty"`
	Id                          *string                                   `json:"id,omitempty"`
	IsGlobalScript              *bool                                     `json:"isGlobalScript,omitempty"`
	LastModifiedDateTime        *string                                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                   `json:"@odata.type,omitempty"`
	Publisher                   *string                                   `json:"publisher,omitempty"`
	RemediationScriptContent    *string                                   `json:"remediationScriptContent,omitempty"`
	RemediationScriptParameters *[]DeviceHealthScriptParameter            `json:"remediationScriptParameters,omitempty"`
	RoleScopeTagIds             *[]string                                 `json:"roleScopeTagIds,omitempty"`
	RunAs32Bit                  *bool                                     `json:"runAs32Bit,omitempty"`
	RunAsAccount                *DeviceHealthScriptRunAsAccount           `json:"runAsAccount,omitempty"`
	RunSummary                  *DeviceHealthScriptRunSummary             `json:"runSummary,omitempty"`
	Version                     *string                                   `json:"version,omitempty"`
}
