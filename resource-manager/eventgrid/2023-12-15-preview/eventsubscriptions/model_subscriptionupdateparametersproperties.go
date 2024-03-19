package eventsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionUpdateParametersProperties struct {
	DeliveryConfiguration *DeliveryConfiguration `json:"deliveryConfiguration,omitempty"`
	EventDeliverySchema   *DeliverySchema        `json:"eventDeliverySchema,omitempty"`
	FiltersConfiguration  *FiltersConfiguration  `json:"filtersConfiguration,omitempty"`
}
