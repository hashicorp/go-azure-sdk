package daprsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprSubscriptionBulkSubscribeOptions struct {
	Enabled            *bool  `json:"enabled,omitempty"`
	MaxAwaitDurationMs *int64 `json:"maxAwaitDurationMs,omitempty"`
	MaxMessagesCount   *int64 `json:"maxMessagesCount,omitempty"`
}
