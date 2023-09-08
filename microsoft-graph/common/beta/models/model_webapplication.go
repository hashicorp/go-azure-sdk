package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplication struct {
	HomePageUrl             *string                `json:"homePageUrl,omitempty"`
	ImplicitGrantSettings   *ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`
	LogoutUrl               *string                `json:"logoutUrl,omitempty"`
	ODataType               *string                `json:"@odata.type,omitempty"`
	Oauth2AllowImplicitFlow *bool                  `json:"oauth2AllowImplicitFlow,omitempty"`
	RedirectUriSettings     *[]RedirectUriSettings `json:"redirectUriSettings,omitempty"`
	RedirectUris            *[]string              `json:"redirectUris,omitempty"`
}
