package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpAddress struct {
	AutonomousSystem  *SecurityAutonomousSystem   `json:"autonomousSystem,omitempty"`
	Components        *[]SecurityHostComponent    `json:"components,omitempty"`
	Cookies           *[]SecurityHostCookie       `json:"cookies,omitempty"`
	CountryOrRegion   *string                     `json:"countryOrRegion,omitempty"`
	FirstSeenDateTime *string                     `json:"firstSeenDateTime,omitempty"`
	HostingProvider   *string                     `json:"hostingProvider,omitempty"`
	Id                *string                     `json:"id,omitempty"`
	LastSeenDateTime  *string                     `json:"lastSeenDateTime,omitempty"`
	Netblock          *string                     `json:"netblock,omitempty"`
	ODataType         *string                     `json:"@odata.type,omitempty"`
	PassiveDns        *[]SecurityPassiveDnsRecord `json:"passiveDns,omitempty"`
	PassiveDnsReverse *[]SecurityPassiveDnsRecord `json:"passiveDnsReverse,omitempty"`
	Reputation        *SecurityHostReputation     `json:"reputation,omitempty"`
	Trackers          *[]SecurityHostTracker      `json:"trackers,omitempty"`
}
