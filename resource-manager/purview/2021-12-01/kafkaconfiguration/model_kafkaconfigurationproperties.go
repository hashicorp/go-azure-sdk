package kafkaconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KafkaConfigurationProperties struct {
	ConsumerGroup       *string              `json:"consumerGroup,omitempty"`
	Credentials         *Credentials         `json:"credentials,omitempty"`
	EventHubPartitionId *string              `json:"eventHubPartitionId,omitempty"`
	EventHubResourceId  *string              `json:"eventHubResourceId,omitempty"`
	EventHubType        *EventHubType        `json:"eventHubType,omitempty"`
	EventStreamingState *EventStreamingState `json:"eventStreamingState,omitempty"`
	EventStreamingType  *EventStreamingType  `json:"eventStreamingType,omitempty"`
}
