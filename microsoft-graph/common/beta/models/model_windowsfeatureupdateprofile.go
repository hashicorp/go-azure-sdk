package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFeatureUpdateProfile struct {
	Assignments                  *[]WindowsFeatureUpdateProfileAssignment `json:"assignments,omitempty"`
	CreatedDateTime              *string                                  `json:"createdDateTime,omitempty"`
	DeployableContentDisplayName *string                                  `json:"deployableContentDisplayName,omitempty"`
	Description                  *string                                  `json:"description,omitempty"`
	DisplayName                  *string                                  `json:"displayName,omitempty"`
	EndOfSupportDate             *string                                  `json:"endOfSupportDate,omitempty"`
	FeatureUpdateVersion         *string                                  `json:"featureUpdateVersion,omitempty"`
	Id                           *string                                  `json:"id,omitempty"`
	LastModifiedDateTime         *string                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType                    *string                                  `json:"@odata.type,omitempty"`
	RoleScopeTagIds              *[]string                                `json:"roleScopeTagIds,omitempty"`
	RolloutSettings              *WindowsUpdateRolloutSettings            `json:"rolloutSettings,omitempty"`
}
