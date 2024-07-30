package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScript struct {
	Assignments          *[]DeviceManagementScriptAssignment                  `json:"assignments,omitempty"`
	CreatedDateTime      *string                                              `json:"createdDateTime,omitempty"`
	CustomAttributeName  *string                                              `json:"customAttributeName,omitempty"`
	CustomAttributeType  *DeviceCustomAttributeShellScriptCustomAttributeType `json:"customAttributeType,omitempty"`
	Description          *string                                              `json:"description,omitempty"`
	DeviceRunStates      *[]DeviceManagementScriptDeviceState                 `json:"deviceRunStates,omitempty"`
	DisplayName          *string                                              `json:"displayName,omitempty"`
	FileName             *string                                              `json:"fileName,omitempty"`
	GroupAssignments     *[]DeviceManagementScriptGroupAssignment             `json:"groupAssignments,omitempty"`
	Id                   *string                                              `json:"id,omitempty"`
	LastModifiedDateTime *string                                              `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                              `json:"@odata.type,omitempty"`
	RoleScopeTagIds      *[]string                                            `json:"roleScopeTagIds,omitempty"`
	RunAsAccount         *DeviceCustomAttributeShellScriptRunAsAccount        `json:"runAsAccount,omitempty"`
	RunSummary           *DeviceManagementScriptRunSummary                    `json:"runSummary,omitempty"`
	ScriptContent        *string                                              `json:"scriptContent,omitempty"`
	UserRunStates        *[]DeviceManagementScriptUserState                   `json:"userRunStates,omitempty"`
}
