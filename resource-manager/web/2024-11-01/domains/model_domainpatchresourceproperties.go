package domains

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainPatchResourceProperties struct {
	AuthCode                    *string                       `json:"authCode,omitempty"`
	AutoRenew                   *bool                         `json:"autoRenew,omitempty"`
	Consent                     DomainPurchaseConsent         `json:"consent"`
	ContactAdmin                Contact                       `json:"contactAdmin"`
	ContactBilling              Contact                       `json:"contactBilling"`
	ContactRegistrant           Contact                       `json:"contactRegistrant"`
	ContactTech                 Contact                       `json:"contactTech"`
	CreatedTime                 *string                       `json:"createdTime,omitempty"`
	DnsType                     *DnsType                      `json:"dnsType,omitempty"`
	DnsZoneId                   *string                       `json:"dnsZoneId,omitempty"`
	DomainNotRenewableReasons   *[]ResourceNotRenewableReason `json:"domainNotRenewableReasons,omitempty"`
	ExpirationTime              *string                       `json:"expirationTime,omitempty"`
	LastRenewedTime             *string                       `json:"lastRenewedTime,omitempty"`
	ManagedHostNames            *[]HostName                   `json:"managedHostNames,omitempty"`
	NameServers                 *[]string                     `json:"nameServers,omitempty"`
	Privacy                     *bool                         `json:"privacy,omitempty"`
	ProvisioningState           *ProvisioningState            `json:"provisioningState,omitempty"`
	ReadyForDnsRecordManagement *bool                         `json:"readyForDnsRecordManagement,omitempty"`
	RegistrationStatus          *DomainStatus                 `json:"registrationStatus,omitempty"`
	TargetDnsType               *DnsType                      `json:"targetDnsType,omitempty"`
}

func (o *DomainPatchResourceProperties) GetCreatedTimeAsTime() (*time.Time, error) {
	if o.CreatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DomainPatchResourceProperties) SetCreatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTime = &formatted
}

func (o *DomainPatchResourceProperties) GetExpirationTimeAsTime() (*time.Time, error) {
	if o.ExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DomainPatchResourceProperties) SetExpirationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTime = &formatted
}

func (o *DomainPatchResourceProperties) GetLastRenewedTimeAsTime() (*time.Time, error) {
	if o.LastRenewedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRenewedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DomainPatchResourceProperties) SetLastRenewedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRenewedTime = &formatted
}
