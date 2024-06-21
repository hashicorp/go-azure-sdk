package hypervsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAgentProperties struct {
	Id               *string `json:"id,omitempty"`
	KeyVaultId       *string `json:"keyVaultId,omitempty"`
	KeyVaultUri      *string `json:"keyVaultUri,omitempty"`
	LastHeartBeatUtc *string `json:"lastHeartBeatUtc,omitempty"`
	Version          *string `json:"version,omitempty"`
}
