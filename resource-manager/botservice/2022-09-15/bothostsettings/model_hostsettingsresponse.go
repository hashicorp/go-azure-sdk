package bothostsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostSettingsResponse struct {
	BotOpenIdMetadata                  *string `json:"BotOpenIdMetadata,omitempty"`
	OAuthURL                           *string `json:"OAuthUrl,omitempty"`
	ToBotFromChannelOpenIdMetadataURL  *string `json:"ToBotFromChannelOpenIdMetadataUrl,omitempty"`
	ToBotFromChannelTokenIssuer        *string `json:"ToBotFromChannelTokenIssuer,omitempty"`
	ToBotFromEmulatorOpenIdMetadataURL *string `json:"ToBotFromEmulatorOpenIdMetadataUrl,omitempty"`
	ToChannelFromBotLoginURL           *string `json:"ToChannelFromBotLoginUrl,omitempty"`
	ToChannelFromBotOAuthScope         *string `json:"ToChannelFromBotOAuthScope,omitempty"`
	ValidateAuthority                  *bool   `json:"ValidateAuthority,omitempty"`
}
