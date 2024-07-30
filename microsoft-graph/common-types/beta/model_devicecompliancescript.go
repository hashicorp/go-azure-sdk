package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScript struct {
	Assignments            *[]DeviceHealthScriptAssignment      `json:"assignments,omitempty"`
	CreatedDateTime        *string                              `json:"createdDateTime,omitempty"`
	Description            *string                              `json:"description,omitempty"`
	DetectionScriptContent *string                              `json:"detectionScriptContent,omitempty"`
	DeviceRunStates        *[]DeviceComplianceScriptDeviceState `json:"deviceRunStates,omitempty"`
	DisplayName            *string                              `json:"displayName,omitempty"`
	EnforceSignatureCheck  *bool                                `json:"enforceSignatureCheck,omitempty"`
	Id                     *string                              `json:"id,omitempty"`
	LastModifiedDateTime   *string                              `json:"lastModifiedDateTime,omitempty"`
	ODataType              *string                              `json:"@odata.type,omitempty"`
	Publisher              *string                              `json:"publisher,omitempty"`
	RoleScopeTagIds        *[]string                            `json:"roleScopeTagIds,omitempty"`
	RunAs32Bit             *bool                                `json:"runAs32Bit,omitempty"`
	RunAsAccount           *DeviceComplianceScriptRunAsAccount  `json:"runAsAccount,omitempty"`
	RunSummary             *DeviceComplianceScriptRunSummary    `json:"runSummary,omitempty"`
	Version                *string                              `json:"version,omitempty"`
}
