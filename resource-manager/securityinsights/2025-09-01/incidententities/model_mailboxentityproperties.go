package incidententities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxEntityProperties struct {
	AdditionalData            *interface{} `json:"additionalData,omitempty"`
	DisplayName               *string      `json:"displayName,omitempty"`
	ExternalDirectoryObjectId *string      `json:"externalDirectoryObjectId,omitempty"`
	FriendlyName              *string      `json:"friendlyName,omitempty"`
	MailboxPrimaryAddress     *string      `json:"mailboxPrimaryAddress,omitempty"`
	Upn                       *string      `json:"upn,omitempty"`
}
