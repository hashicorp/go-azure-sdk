package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummary struct {
	ODataType                     *string                                  `json:"@odata.type,omitempty"`
	TotalUserCount                *int64                                   `json:"totalUserCount,omitempty"`
	UserRegistrationFeatureCounts *[]UserRegistrationFeatureCount          `json:"userRegistrationFeatureCounts,omitempty"`
	UserRoles                     *UserRegistrationFeatureSummaryUserRoles `json:"userRoles,omitempty"`
	UserTypes                     *UserRegistrationFeatureSummaryUserTypes `json:"userTypes,omitempty"`
}
