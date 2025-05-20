package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppInsightsWebAppStackSettings struct {
	IsDefaultOff *bool `json:"isDefaultOff,omitempty"`
	IsSupported  *bool `json:"isSupported,omitempty"`
}
