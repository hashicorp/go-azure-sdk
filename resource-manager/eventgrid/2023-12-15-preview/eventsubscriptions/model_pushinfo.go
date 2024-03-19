package eventsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PushInfo struct {
	DeadLetterDestinationWithResourceIdentity *DeadLetterWithResourceIdentity `json:"deadLetterDestinationWithResourceIdentity,omitempty"`
	DeliveryWithResourceIdentity              *DeliveryWithResourceIdentity   `json:"deliveryWithResourceIdentity,omitempty"`
	EventTimeToLive                           *string                         `json:"eventTimeToLive,omitempty"`
	MaxDeliveryCount                          *int64                          `json:"maxDeliveryCount,omitempty"`
}
