package daprsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprSubscriptionProperties struct {
	BulkSubscribe   *DaprSubscriptionBulkSubscribeOptions `json:"bulkSubscribe,omitempty"`
	DeadLetterTopic *string                               `json:"deadLetterTopic,omitempty"`
	Metadata        *map[string]string                    `json:"metadata,omitempty"`
	PubsubName      *string                               `json:"pubsubName,omitempty"`
	Routes          *DaprSubscriptionRoutes               `json:"routes,omitempty"`
	Scopes          *[]string                             `json:"scopes,omitempty"`
	Topic           *string                               `json:"topic,omitempty"`
}
