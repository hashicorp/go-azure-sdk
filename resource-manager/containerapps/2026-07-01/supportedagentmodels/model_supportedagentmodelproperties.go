package supportedagentmodels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedAgentModelProperties struct {
	Default             bool    `json:"default"`
	Model               string  `json:"model"`
	ModelDisplayName    *string `json:"modelDisplayName,omitempty"`
	Multiplier          *string `json:"multiplier,omitempty"`
	Provider            string  `json:"provider"`
	ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
}
