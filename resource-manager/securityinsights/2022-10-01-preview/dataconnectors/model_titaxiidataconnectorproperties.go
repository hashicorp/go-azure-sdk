package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiTaxiiDataConnectorProperties struct {
	CollectionId        *string                       `json:"collectionId,omitempty"`
	DataTypes           TiTaxiiDataConnectorDataTypes `json:"dataTypes"`
	FriendlyName        *string                       `json:"friendlyName,omitempty"`
	Password            *string                       `json:"password,omitempty"`
	PollingFrequency    PollingFrequency              `json:"pollingFrequency"`
	TaxiiLookbackPeriod *string                       `json:"taxiiLookbackPeriod,omitempty"`
	TaxiiServer         *string                       `json:"taxiiServer,omitempty"`
	TenantId            string                        `json:"tenantId"`
	UserName            *string                       `json:"userName,omitempty"`
	WorkspaceId         *string                       `json:"workspaceId,omitempty"`
}
