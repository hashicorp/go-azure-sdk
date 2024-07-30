package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnector struct {
	ConnectorServerName   *string                                                 `json:"connectorServerName,omitempty"`
	ExchangeAlias         *string                                                 `json:"exchangeAlias,omitempty"`
	ExchangeConnectorType *DeviceManagementExchangeConnectorExchangeConnectorType `json:"exchangeConnectorType,omitempty"`
	ExchangeOrganization  *string                                                 `json:"exchangeOrganization,omitempty"`
	Id                    *string                                                 `json:"id,omitempty"`
	LastSyncDateTime      *string                                                 `json:"lastSyncDateTime,omitempty"`
	ODataType             *string                                                 `json:"@odata.type,omitempty"`
	PrimarySmtpAddress    *string                                                 `json:"primarySmtpAddress,omitempty"`
	ServerName            *string                                                 `json:"serverName,omitempty"`
	Status                *DeviceManagementExchangeConnectorStatus                `json:"status,omitempty"`
	Version               *string                                                 `json:"version,omitempty"`
}
