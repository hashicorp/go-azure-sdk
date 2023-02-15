package productapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuth2AuthenticationSettingsContract struct {
	AuthorizationServerId *string `json:"authorizationServerId,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
}
