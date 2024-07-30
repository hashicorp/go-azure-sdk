package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateProfile struct {
	Assignments                  *[]WindowsQualityUpdateProfileAssignment `json:"assignments,omitempty"`
	CreatedDateTime              *string                                  `json:"createdDateTime,omitempty"`
	DeployableContentDisplayName *string                                  `json:"deployableContentDisplayName,omitempty"`
	Description                  *string                                  `json:"description,omitempty"`
	DisplayName                  *string                                  `json:"displayName,omitempty"`
	ExpeditedUpdateSettings      *ExpeditedWindowsQualityUpdateSettings   `json:"expeditedUpdateSettings,omitempty"`
	Id                           *string                                  `json:"id,omitempty"`
	LastModifiedDateTime         *string                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType                    *string                                  `json:"@odata.type,omitempty"`
	ReleaseDateDisplayName       *string                                  `json:"releaseDateDisplayName,omitempty"`
	RoleScopeTagIds              *[]string                                `json:"roleScopeTagIds,omitempty"`
}
