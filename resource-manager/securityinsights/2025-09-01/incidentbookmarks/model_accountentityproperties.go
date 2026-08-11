package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountEntityProperties struct {
	AadTenantId    *string                 `json:"aadTenantId,omitempty"`
	AadUserId      *string                 `json:"aadUserId,omitempty"`
	AccountName    *string                 `json:"accountName,omitempty"`
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	DisplayName    *string                 `json:"displayName,omitempty"`
	DnsDomain      *string                 `json:"dnsDomain,omitempty"`
	FriendlyName   *string                 `json:"friendlyName,omitempty"`
	HostEntityId   *string                 `json:"hostEntityId,omitempty"`
	IsDomainJoined *bool                   `json:"isDomainJoined,omitempty"`
	NtDomain       *string                 `json:"ntDomain,omitempty"`
	ObjectGuid     *string                 `json:"objectGuid,omitempty"`
	Puid           *string                 `json:"puid,omitempty"`
	Sid            *string                 `json:"sid,omitempty"`
	UpnSuffix      *string                 `json:"upnSuffix,omitempty"`
}
