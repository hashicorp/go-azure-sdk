package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAuthInfo struct {
	Audience                     *string   `json:"audience,omitempty"`
	Issuer                       *string   `json:"issuer,omitempty"`
	JwksUri                      *string   `json:"jwksUri,omitempty"`
	Scopes                       *[]string `json:"scopes,omitempty"`
	WellKnownOpenIdConfiguration *string   `json:"wellKnownOpenIdConfiguration,omitempty"`
}
