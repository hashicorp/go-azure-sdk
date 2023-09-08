package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureRolloutPolicy struct {
	AppliesTo               *[]DirectoryObject           `json:"appliesTo,omitempty"`
	Description             *string                      `json:"description,omitempty"`
	DisplayName             *string                      `json:"displayName,omitempty"`
	Feature                 *FeatureRolloutPolicyFeature `json:"feature,omitempty"`
	Id                      *string                      `json:"id,omitempty"`
	IsAppliedToOrganization *bool                        `json:"isAppliedToOrganization,omitempty"`
	IsEnabled               *bool                        `json:"isEnabled,omitempty"`
	ODataType               *string                      `json:"@odata.type,omitempty"`
}
