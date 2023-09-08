package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTemplate struct {
	Categories                 *[]string `json:"categories,omitempty"`
	Description                *string   `json:"description,omitempty"`
	DisplayName                *string   `json:"displayName,omitempty"`
	HomePageUrl                *string   `json:"homePageUrl,omitempty"`
	Id                         *string   `json:"id,omitempty"`
	LogoUrl                    *string   `json:"logoUrl,omitempty"`
	ODataType                  *string   `json:"@odata.type,omitempty"`
	Publisher                  *string   `json:"publisher,omitempty"`
	SupportedProvisioningTypes *[]string `json:"supportedProvisioningTypes,omitempty"`
	SupportedSingleSignOnModes *[]string `json:"supportedSingleSignOnModes,omitempty"`
}
