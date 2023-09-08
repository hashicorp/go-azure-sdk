package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsent struct {
	ODataType                        *string                                       `json:"@odata.type,omitempty"`
	ShareAPNSData                    *AdminConsentShareAPNSData                    `json:"shareAPNSData,omitempty"`
	ShareUserExperienceAnalyticsData *AdminConsentShareUserExperienceAnalyticsData `json:"shareUserExperienceAnalyticsData,omitempty"`
}
