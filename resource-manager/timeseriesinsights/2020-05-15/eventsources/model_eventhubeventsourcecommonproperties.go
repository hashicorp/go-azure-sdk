package eventsources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubEventSourceCommonProperties struct {
	ConsumerGroupName     string                    `json:"consumerGroupName"`
	CreationTime          *string                   `json:"creationTime,omitempty"`
	EventHubName          string                    `json:"eventHubName"`
	EventSourceResourceId string                    `json:"eventSourceResourceId"`
	IngressStartAt        *IngressStartAtProperties `json:"ingressStartAt,omitempty"`
	KeyName               string                    `json:"keyName"`
	LocalTimestamp        *LocalTimestamp           `json:"localTimestamp,omitempty"`
	ProvisioningState     *ProvisioningState        `json:"provisioningState,omitempty"`
	ServiceBusNamespace   string                    `json:"serviceBusNamespace"`
	TimestampPropertyName *string                   `json:"timestampPropertyName,omitempty"`
}
