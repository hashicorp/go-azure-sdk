package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaign struct {
	EnforceRegistrationAfterAllowedSnoozes *bool                                                     `json:"enforceRegistrationAfterAllowedSnoozes,omitempty"`
	ExcludeTargets                         *[]ExcludeTarget                                          `json:"excludeTargets,omitempty"`
	IncludeTargets                         *[]AuthenticationMethodsRegistrationCampaignIncludeTarget `json:"includeTargets,omitempty"`
	ODataType                              *string                                                   `json:"@odata.type,omitempty"`
	SnoozeDurationInDays                   *int64                                                    `json:"snoozeDurationInDays,omitempty"`
	State                                  *AuthenticationMethodsRegistrationCampaignState           `json:"state,omitempty"`
}
