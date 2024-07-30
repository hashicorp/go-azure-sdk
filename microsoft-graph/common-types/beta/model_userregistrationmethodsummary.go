package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationMethodSummary struct {
	ODataType                    *string                                 `json:"@odata.type,omitempty"`
	TotalUserCount               *int64                                  `json:"totalUserCount,omitempty"`
	UserRegistrationMethodCounts *[]UserRegistrationMethodCount          `json:"userRegistrationMethodCounts,omitempty"`
	UserRoles                    *UserRegistrationMethodSummaryUserRoles `json:"userRoles,omitempty"`
	UserTypes                    *UserRegistrationMethodSummaryUserTypes `json:"userTypes,omitempty"`
}
