package domains

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
