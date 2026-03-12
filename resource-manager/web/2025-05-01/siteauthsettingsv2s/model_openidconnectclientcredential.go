package siteauthsettingsv2s

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectClientCredential struct {
	ClientSecretSettingName *string `json:"clientSecretSettingName,omitempty"`
	Method                  *Method `json:"method,omitempty"`
}
