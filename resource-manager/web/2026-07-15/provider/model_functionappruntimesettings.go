package provider

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionAppRuntimeSettings struct {
	AppInsightsSettings                 *AppInsightsWebAppStackSettings  `json:"appInsightsSettings,omitempty"`
	AppSettingsDictionary               *map[string]string               `json:"appSettingsDictionary,omitempty"`
	EndOfLifeDate                       *string                          `json:"endOfLifeDate,omitempty"`
	GitHubActionSettings                *GitHubActionWebAppStackSettings `json:"gitHubActionSettings,omitempty"`
	IsAutoUpdate                        *bool                            `json:"isAutoUpdate,omitempty"`
	IsDefault                           *bool                            `json:"isDefault,omitempty"`
	IsDeprecated                        *bool                            `json:"isDeprecated,omitempty"`
	IsEarlyAccess                       *bool                            `json:"isEarlyAccess,omitempty"`
	IsHidden                            *bool                            `json:"isHidden,omitempty"`
	IsPreview                           *bool                            `json:"isPreview,omitempty"`
	RemoteDebuggingSupported            *bool                            `json:"remoteDebuggingSupported,omitempty"`
	RuntimeVersion                      *string                          `json:"runtimeVersion,omitempty"`
	SiteConfigPropertiesDictionary      *SiteConfigPropertiesDictionary  `json:"siteConfigPropertiesDictionary,omitempty"`
	SupportedFunctionsExtensionVersions *[]string                        `json:"supportedFunctionsExtensionVersions,omitempty"`
}

func (o *FunctionAppRuntimeSettings) GetEndOfLifeDateAsTime() (*time.Time, error) {
	if o.EndOfLifeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndOfLifeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *FunctionAppRuntimeSettings) SetEndOfLifeDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndOfLifeDate = &formatted
}
