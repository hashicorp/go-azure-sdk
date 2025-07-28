package incidentalerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostEntityProperties struct {
	AdditionalData *interface{} `json:"additionalData,omitempty"`
	AzureID        *string      `json:"azureID,omitempty"`
	DnsDomain      *string      `json:"dnsDomain,omitempty"`
	FriendlyName   *string      `json:"friendlyName,omitempty"`
	HostName       *string      `json:"hostName,omitempty"`
	IsDomainJoined *bool        `json:"isDomainJoined,omitempty"`
	NetBiosName    *string      `json:"netBiosName,omitempty"`
	NtDomain       *string      `json:"ntDomain,omitempty"`
	OmsAgentID     *string      `json:"omsAgentID,omitempty"`
	OsFamily       *OSFamily    `json:"osFamily,omitempty"`
	OsVersion      *string      `json:"osVersion,omitempty"`
}
