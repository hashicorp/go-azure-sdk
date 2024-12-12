package activesessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInfoProperties struct {
	CustomInfo      *CustomInfoProperties      `json:"customInfo,omitempty"`
	MarketplaceInfo *MarketplaceInfoProperties `json:"marketplaceInfo,omitempty"`
	Type            Type                       `json:"type"`
}
