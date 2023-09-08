package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LandingPage struct {
	CreatedBy            *EmailIdentity       `json:"createdBy,omitempty"`
	CreatedDateTime      *string              `json:"createdDateTime,omitempty"`
	Description          *string              `json:"description,omitempty"`
	Details              *[]LandingPageDetail `json:"details,omitempty"`
	DisplayName          *string              `json:"displayName,omitempty"`
	Id                   *string              `json:"id,omitempty"`
	LastModifiedBy       *EmailIdentity       `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string              `json:"lastModifiedDateTime,omitempty"`
	Locale               *string              `json:"locale,omitempty"`
	ODataType            *string              `json:"@odata.type,omitempty"`
	Source               *LandingPageSource   `json:"source,omitempty"`
	Status               *LandingPageStatus   `json:"status,omitempty"`
	SupportedLocales     *[]string            `json:"supportedLocales,omitempty"`
}
