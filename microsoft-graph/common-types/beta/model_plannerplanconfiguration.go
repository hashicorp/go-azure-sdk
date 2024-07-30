package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanConfiguration struct {
	Buckets              *[]PlannerPlanConfigurationBucketDefinition `json:"buckets,omitempty"`
	CreatedBy            *IdentitySet                                `json:"createdBy,omitempty"`
	CreatedDateTime      *string                                     `json:"createdDateTime,omitempty"`
	DefaultLanguage      *string                                     `json:"defaultLanguage,omitempty"`
	Id                   *string                                     `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                                     `json:"lastModifiedDateTime,omitempty"`
	Localizations        *[]PlannerPlanConfigurationLocalization     `json:"localizations,omitempty"`
	ODataType            *string                                     `json:"@odata.type,omitempty"`
}
