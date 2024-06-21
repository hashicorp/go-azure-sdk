package consumerinvitation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerInvitationProperties struct {
	DataSetCount       *int64            `json:"dataSetCount,omitempty"`
	Description        *string           `json:"description,omitempty"`
	InvitationId       string            `json:"invitationId"`
	InvitationStatus   *InvitationStatus `json:"invitationStatus,omitempty"`
	Location           *string           `json:"location,omitempty"`
	ProviderEmail      *string           `json:"providerEmail,omitempty"`
	ProviderName       *string           `json:"providerName,omitempty"`
	ProviderTenantName *string           `json:"providerTenantName,omitempty"`
	RespondedAt        *string           `json:"respondedAt,omitempty"`
	SentAt             *string           `json:"sentAt,omitempty"`
	ShareName          *string           `json:"shareName,omitempty"`
	TermsOfUse         *string           `json:"termsOfUse,omitempty"`
	UserEmail          *string           `json:"userEmail,omitempty"`
	UserName           *string           `json:"userName,omitempty"`
}
