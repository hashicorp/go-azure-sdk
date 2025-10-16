package broker

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerSubscriberQueueCustomPolicySettings struct {
	Dynamic             *BrokerSubscriberQueueDynamic `json:"dynamic,omitempty"`
	SubscriberClientIds *[]string                     `json:"subscriberClientIds,omitempty"`
}
