package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpAddress struct {
	AutonomousSystem  *SecurityAutonomousSystem     `json:"autonomousSystem,omitempty"`
	ChildHostPairs    *[]SecurityHostPair           `json:"childHostPairs,omitempty"`
	Components        *[]SecurityHostComponent      `json:"components,omitempty"`
	Cookies           *[]SecurityHostCookie         `json:"cookies,omitempty"`
	CountryOrRegion   *string                       `json:"countryOrRegion,omitempty"`
	FirstSeenDateTime *string                       `json:"firstSeenDateTime,omitempty"`
	HostPairs         *[]SecurityHostPair           `json:"hostPairs,omitempty"`
	HostingProvider   *string                       `json:"hostingProvider,omitempty"`
	Id                *string                       `json:"id,omitempty"`
	LastSeenDateTime  *string                       `json:"lastSeenDateTime,omitempty"`
	Netblock          *string                       `json:"netblock,omitempty"`
	ODataType         *string                       `json:"@odata.type,omitempty"`
	ParentHostPairs   *[]SecurityHostPair           `json:"parentHostPairs,omitempty"`
	PassiveDns        *[]SecurityPassiveDnsRecord   `json:"passiveDns,omitempty"`
	PassiveDnsReverse *[]SecurityPassiveDnsRecord   `json:"passiveDnsReverse,omitempty"`
	Ports             *[]SecurityHostPort           `json:"ports,omitempty"`
	Reputation        *SecurityHostReputation       `json:"reputation,omitempty"`
	SslCertificates   *[]SecurityHostSslCertificate `json:"sslCertificates,omitempty"`
	Subdomains        *[]SecuritySubdomain          `json:"subdomains,omitempty"`
	Trackers          *[]SecurityHostTracker        `json:"trackers,omitempty"`
	Whois             *SecurityWhoisRecord          `json:"whois,omitempty"`
}
