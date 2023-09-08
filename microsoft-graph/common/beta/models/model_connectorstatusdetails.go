package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetails struct {
	ConnectorInstanceId *string                              `json:"connectorInstanceId,omitempty"`
	ConnectorName       *ConnectorStatusDetailsConnectorName `json:"connectorName,omitempty"`
	EventDateTime       *string                              `json:"eventDateTime,omitempty"`
	ODataType           *string                              `json:"@odata.type,omitempty"`
	Status              *ConnectorStatusDetailsStatus        `json:"status,omitempty"`
}
