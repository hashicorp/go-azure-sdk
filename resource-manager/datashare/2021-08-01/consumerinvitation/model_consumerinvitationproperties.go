package consumerinvitation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerInvitationProperties struct {
	DataSetCount       *int64            `json:"dataSetCount,omitempty"`
	Description        *string           `json:"description,omitempty"`
	ExpirationDate     *string           `json:"expirationDate,omitempty"`
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

func (o *ConsumerInvitationProperties) GetExpirationDateAsTime() (*time.Time, error) {
	if o.ExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ConsumerInvitationProperties) SetExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationDate = &formatted
}

func (o *ConsumerInvitationProperties) GetRespondedAtAsTime() (*time.Time, error) {
	if o.RespondedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RespondedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *ConsumerInvitationProperties) SetRespondedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RespondedAt = &formatted
}

func (o *ConsumerInvitationProperties) GetSentAtAsTime() (*time.Time, error) {
	if o.SentAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SentAt, "2006-01-02T15:04:05Z07:00")
}

func (o *ConsumerInvitationProperties) SetSentAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SentAt = &formatted
}
