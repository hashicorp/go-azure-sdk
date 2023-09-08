package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSOnboardingSettings struct {
	Id                        *string                                     `json:"id,omitempty"`
	LastDirectorySyncDateTime *string                                     `json:"lastDirectorySyncDateTime,omitempty"`
	LastModifiedDateTime      *string                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string                                     `json:"@odata.type,omitempty"`
	OnboardingStatus          *ChromeOSOnboardingSettingsOnboardingStatus `json:"onboardingStatus,omitempty"`
	OwnerUserPrincipalName    *string                                     `json:"ownerUserPrincipalName,omitempty"`
}
