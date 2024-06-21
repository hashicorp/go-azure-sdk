package syncagents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncAgentProperties struct {
	ExpiryTime     *string         `json:"expiryTime,omitempty"`
	IsUpToDate     *bool           `json:"isUpToDate,omitempty"`
	LastAliveTime  *string         `json:"lastAliveTime,omitempty"`
	Name           *string         `json:"name,omitempty"`
	State          *SyncAgentState `json:"state,omitempty"`
	SyncDatabaseId *string         `json:"syncDatabaseId,omitempty"`
	Version        *string         `json:"version,omitempty"`
}
