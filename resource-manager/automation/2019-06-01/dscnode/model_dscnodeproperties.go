package dscnode

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeProperties struct {
	AccountId         *string                                       `json:"accountId,omitempty"`
	Etag              *string                                       `json:"etag,omitempty"`
	ExtensionHandler  *[]DscNodeExtensionHandlerAssociationProperty `json:"extensionHandler,omitempty"`
	IP                *string                                       `json:"ip,omitempty"`
	LastSeen          *string                                       `json:"lastSeen,omitempty"`
	NodeConfiguration *DscNodeConfigurationAssociationProperty      `json:"nodeConfiguration,omitempty"`
	NodeId            *string                                       `json:"nodeId,omitempty"`
	RegistrationTime  *string                                       `json:"registrationTime,omitempty"`
	Status            *string                                       `json:"status,omitempty"`
	TotalCount        *int64                                        `json:"totalCount,omitempty"`
}
