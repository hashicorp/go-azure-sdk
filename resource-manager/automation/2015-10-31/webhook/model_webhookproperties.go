package webhook

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebhookProperties struct {
	CreationTime     *string                     `json:"creationTime,omitempty"`
	Description      *string                     `json:"description,omitempty"`
	ExpiryTime       *string                     `json:"expiryTime,omitempty"`
	IsEnabled        *bool                       `json:"isEnabled,omitempty"`
	LastInvokedTime  *string                     `json:"lastInvokedTime,omitempty"`
	LastModifiedBy   *string                     `json:"lastModifiedBy,omitempty"`
	LastModifiedTime *string                     `json:"lastModifiedTime,omitempty"`
	Parameters       *map[string]string          `json:"parameters,omitempty"`
	RunOn            *string                     `json:"runOn,omitempty"`
	Runbook          *RunbookAssociationProperty `json:"runbook,omitempty"`
	Uri              *string                     `json:"uri,omitempty"`
}
