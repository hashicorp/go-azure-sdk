package namespacediscoveredassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MqttDestinationConfiguration struct {
	Qos    *MqttDestinationQos `json:"qos,omitempty"`
	Retain *TopicRetainType    `json:"retain,omitempty"`
	Topic  string              `json:"topic"`
	Ttl    *int64              `json:"ttl,omitempty"`
}
