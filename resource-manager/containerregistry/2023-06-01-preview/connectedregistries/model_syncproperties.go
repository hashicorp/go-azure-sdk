package connectedregistries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncProperties struct {
	GatewayEndpoint *string `json:"gatewayEndpoint,omitempty"`
	LastSyncTime    *string `json:"lastSyncTime,omitempty"`
	MessageTtl      string  `json:"messageTtl"`
	Schedule        *string `json:"schedule,omitempty"`
	SyncWindow      *string `json:"syncWindow,omitempty"`
	TokenId         string  `json:"tokenId"`
}
