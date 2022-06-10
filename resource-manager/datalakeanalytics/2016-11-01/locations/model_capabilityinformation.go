package locations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityInformation struct {
	AccountCount    *int64             `json:"accountCount,omitempty"`
	MaxAccountCount *int64             `json:"maxAccountCount,omitempty"`
	MigrationState  *bool              `json:"migrationState,omitempty"`
	State           *SubscriptionState `json:"state,omitempty"`
	SubscriptionId  *string            `json:"subscriptionId,omitempty"`
}
