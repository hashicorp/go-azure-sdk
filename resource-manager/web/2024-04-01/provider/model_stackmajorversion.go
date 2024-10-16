package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StackMajorVersion struct {
	AppSettingsDictionary          *map[string]interface{} `json:"appSettingsDictionary,omitempty"`
	ApplicationInsights            *bool                   `json:"applicationInsights,omitempty"`
	DisplayVersion                 *string                 `json:"displayVersion,omitempty"`
	IsDefault                      *bool                   `json:"isDefault,omitempty"`
	IsDeprecated                   *bool                   `json:"isDeprecated,omitempty"`
	IsHidden                       *bool                   `json:"isHidden,omitempty"`
	IsPreview                      *bool                   `json:"isPreview,omitempty"`
	MinorVersions                  *[]StackMinorVersion    `json:"minorVersions,omitempty"`
	RuntimeVersion                 *string                 `json:"runtimeVersion,omitempty"`
	SiteConfigPropertiesDictionary *map[string]interface{} `json:"siteConfigPropertiesDictionary,omitempty"`
}
