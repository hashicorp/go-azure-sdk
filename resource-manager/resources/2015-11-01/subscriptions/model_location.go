package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Location struct {
	DisplayName    *string `json:"displayName,omitempty"`
	Id             *string `json:"id,omitempty"`
	Latitude       *string `json:"latitude,omitempty"`
	Longitude      *string `json:"longitude,omitempty"`
	Name           *string `json:"name,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}
