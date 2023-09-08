package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterDetails struct {
	Bio                   *ItemBody `json:"bio,omitempty"`
	Company               *string   `json:"company,omitempty"`
	JobTitle              *string   `json:"jobTitle,omitempty"`
	LinkedInProfileWebUrl *string   `json:"linkedInProfileWebUrl,omitempty"`
	ODataType             *string   `json:"@odata.type,omitempty"`
	PersonalSiteWebUrl    *string   `json:"personalSiteWebUrl,omitempty"`
	TwitterProfileWebUrl  *string   `json:"twitterProfileWebUrl,omitempty"`
}
