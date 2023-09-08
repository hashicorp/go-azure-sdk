package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemporaryAccessPassAuthenticationMethodConfiguration struct {
	DefaultLength            *int64                                                     `json:"defaultLength,omitempty"`
	DefaultLifetimeInMinutes *int64                                                     `json:"defaultLifetimeInMinutes,omitempty"`
	ExcludeTargets           *[]ExcludeTarget                                           `json:"excludeTargets,omitempty"`
	Id                       *string                                                    `json:"id,omitempty"`
	IncludeTargets           *[]AuthenticationMethodTarget                              `json:"includeTargets,omitempty"`
	IsUsableOnce             *bool                                                      `json:"isUsableOnce,omitempty"`
	MaximumLifetimeInMinutes *int64                                                     `json:"maximumLifetimeInMinutes,omitempty"`
	MinimumLifetimeInMinutes *int64                                                     `json:"minimumLifetimeInMinutes,omitempty"`
	ODataType                *string                                                    `json:"@odata.type,omitempty"`
	State                    *TemporaryAccessPassAuthenticationMethodConfigurationState `json:"state,omitempty"`
}
