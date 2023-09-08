package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostSecurityState struct {
	Fqdn                      *string `json:"fqdn,omitempty"`
	IsAzureAdJoined           *bool   `json:"isAzureAdJoined,omitempty"`
	IsAzureAdRegistered       *bool   `json:"isAzureAdRegistered,omitempty"`
	IsHybridAzureDomainJoined *bool   `json:"isHybridAzureDomainJoined,omitempty"`
	NetBiosName               *string `json:"netBiosName,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	Os                        *string `json:"os,omitempty"`
	PrivateIpAddress          *string `json:"privateIpAddress,omitempty"`
	PublicIpAddress           *string `json:"publicIpAddress,omitempty"`
	RiskScore                 *string `json:"riskScore,omitempty"`
}
