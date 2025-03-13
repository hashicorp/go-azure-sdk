package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftPurviewInformationProtectionDataConnectorProperties struct {
	DataTypes MicrosoftPurviewInformationProtectionConnectorDataTypes `json:"dataTypes"`
	TenantId  string                                                  `json:"tenantId"`
}
