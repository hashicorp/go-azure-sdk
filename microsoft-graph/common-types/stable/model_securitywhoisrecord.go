package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisRecord struct {
	Abuse                *SecurityWhoisContact         `json:"abuse,omitempty"`
	Admin                *SecurityWhoisContact         `json:"admin,omitempty"`
	Billing              *SecurityWhoisContact         `json:"billing,omitempty"`
	DomainStatus         *string                       `json:"domainStatus,omitempty"`
	ExpirationDateTime   *string                       `json:"expirationDateTime,omitempty"`
	FirstSeenDateTime    *string                       `json:"firstSeenDateTime,omitempty"`
	History              *[]SecurityWhoisHistoryRecord `json:"history,omitempty"`
	Host                 *SecurityHost                 `json:"host,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	LastSeenDateTime     *string                       `json:"lastSeenDateTime,omitempty"`
	LastUpdateDateTime   *string                       `json:"lastUpdateDateTime,omitempty"`
	Nameservers          *[]SecurityWhoisNameserver    `json:"nameservers,omitempty"`
	Noc                  *SecurityWhoisContact         `json:"noc,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	RawWhoisText         *string                       `json:"rawWhoisText,omitempty"`
	Registrant           *SecurityWhoisContact         `json:"registrant,omitempty"`
	Registrar            *SecurityWhoisContact         `json:"registrar,omitempty"`
	RegistrationDateTime *string                       `json:"registrationDateTime,omitempty"`
	Technical            *SecurityWhoisContact         `json:"technical,omitempty"`
	WhoisServer          *string                       `json:"whoisServer,omitempty"`
	Zone                 *SecurityWhoisContact         `json:"zone,omitempty"`
}
