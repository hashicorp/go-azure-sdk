package dataconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubConnectionProperties struct {
	Compression               *Compression        `json:"compression,omitempty"`
	ConsumerGroup             string              `json:"consumerGroup"`
	DataFormat                *EventHubDataFormat `json:"dataFormat,omitempty"`
	EventHubResourceId        string              `json:"eventHubResourceId"`
	EventSystemProperties     *[]string           `json:"eventSystemProperties,omitempty"`
	ManagedIdentityResourceId *string             `json:"managedIdentityResourceId,omitempty"`
	MappingRuleName           *string             `json:"mappingRuleName,omitempty"`
	ProvisioningState         *ProvisioningState  `json:"provisioningState,omitempty"`
	TableName                 *string             `json:"tableName,omitempty"`
}
