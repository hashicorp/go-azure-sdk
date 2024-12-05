package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MTPDataConnectorDataTypes struct {
	Alerts    *DataConnectorDataTypeCommon `json:"alerts,omitempty"`
	Incidents DataConnectorDataTypeCommon  `json:"incidents"`
}
