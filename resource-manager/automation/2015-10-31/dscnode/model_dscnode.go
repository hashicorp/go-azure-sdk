package dscnode

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNode struct {
	AccountId         *string                                       `json:"accountId,omitempty"`
	Etag              *string                                       `json:"etag,omitempty"`
	ExtensionHandler  *[]DscNodeExtensionHandlerAssociationProperty `json:"extensionHandler,omitempty"`
	IP                *string                                       `json:"ip,omitempty"`
	Id                *string                                       `json:"id,omitempty"`
	LastSeen          *string                                       `json:"lastSeen,omitempty"`
	Name              *string                                       `json:"name,omitempty"`
	NodeConfiguration *DscNodeConfigurationAssociationProperty      `json:"nodeConfiguration,omitempty"`
	NodeId            *string                                       `json:"nodeId,omitempty"`
	RegistrationTime  *string                                       `json:"registrationTime,omitempty"`
	Status            *string                                       `json:"status,omitempty"`
	Type              *string                                       `json:"type,omitempty"`
}
