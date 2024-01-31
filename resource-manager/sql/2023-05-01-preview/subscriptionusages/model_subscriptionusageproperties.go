package subscriptionusages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionUsageProperties struct {
	CurrentValue *float64 `json:"currentValue,omitempty"`
	DisplayName  *string  `json:"displayName,omitempty"`
	Limit        *float64 `json:"limit,omitempty"`
	Unit         *string  `json:"unit,omitempty"`
}
