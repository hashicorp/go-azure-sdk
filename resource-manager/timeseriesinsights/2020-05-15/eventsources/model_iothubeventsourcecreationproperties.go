package eventsources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTHubEventSourceCreationProperties struct {
	ConsumerGroupName     string                    `json:"consumerGroupName"`
	CreationTime          *string                   `json:"creationTime,omitempty"`
	EventSourceResourceId string                    `json:"eventSourceResourceId"`
	IngressStartAt        *IngressStartAtProperties `json:"ingressStartAt,omitempty"`
	IotHubName            string                    `json:"iotHubName"`
	KeyName               string                    `json:"keyName"`
	LocalTimestamp        *LocalTimestamp           `json:"localTimestamp,omitempty"`
	ProvisioningState     *ProvisioningState        `json:"provisioningState,omitempty"`
	SharedAccessKey       string                    `json:"sharedAccessKey"`
	TimestampPropertyName *string                   `json:"timestampPropertyName,omitempty"`
}
