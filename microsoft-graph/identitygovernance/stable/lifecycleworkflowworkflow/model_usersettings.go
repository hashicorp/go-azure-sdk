package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettings struct {
	ContributionToContentDiscoveryAsOrganizationDisabled *bool             `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	ContributionToContentDiscoveryDisabled               *bool             `json:"contributionToContentDiscoveryDisabled,omitempty"`
	Id                                                   *string           `json:"id,omitempty"`
	ODataType                                            *string           `json:"@odata.type,omitempty"`
	ShiftPreferences                                     *ShiftPreferences `json:"shiftPreferences,omitempty"`
}
