package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorDataType struct {
	LastDataReceivedQuery string `json:"lastDataReceivedQuery"`
	Name                  string `json:"name"`
}
