package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GCPDataConnectorProperties struct {
	Auth                    GCPAuthProperties    `json:"auth"`
	ConnectorDefinitionName string               `json:"connectorDefinitionName"`
	DcrConfig               *DCRConfiguration    `json:"dcrConfig,omitempty"`
	Request                 GCPRequestProperties `json:"request"`
}
