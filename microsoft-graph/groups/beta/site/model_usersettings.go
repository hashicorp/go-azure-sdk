package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettings struct {
	ContactMergeSuggestions                              *ContactMergeSuggestions     `json:"contactMergeSuggestions,omitempty"`
	ContributionToContentDiscoveryAsOrganizationDisabled *bool                        `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	ContributionToContentDiscoveryDisabled               *bool                        `json:"contributionToContentDiscoveryDisabled,omitempty"`
	Id                                                   *string                      `json:"id,omitempty"`
	ItemInsights                                         *UserInsightsSettings        `json:"itemInsights,omitempty"`
	ODataType                                            *string                      `json:"@odata.type,omitempty"`
	RegionalAndLanguageSettings                          *RegionalAndLanguageSettings `json:"regionalAndLanguageSettings,omitempty"`
	ShiftPreferences                                     *ShiftPreferences            `json:"shiftPreferences,omitempty"`
}
