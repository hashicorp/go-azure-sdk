package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolUpgradeSettings struct {
	DrainTimeout   *int64  `json:"drainTimeout,omitempty"`
	MaxSurge       *string `json:"maxSurge,omitempty"`
	MaxUnavailable *string `json:"maxUnavailable,omitempty"`
}
