package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHost struct {
	Components        *[]SecurityHostComponent    `json:"components,omitempty"`
	Cookies           *[]SecurityHostCookie       `json:"cookies,omitempty"`
	FirstSeenDateTime *string                     `json:"firstSeenDateTime,omitempty"`
	Id                *string                     `json:"id,omitempty"`
	LastSeenDateTime  *string                     `json:"lastSeenDateTime,omitempty"`
	ODataType         *string                     `json:"@odata.type,omitempty"`
	PassiveDns        *[]SecurityPassiveDnsRecord `json:"passiveDns,omitempty"`
	PassiveDnsReverse *[]SecurityPassiveDnsRecord `json:"passiveDnsReverse,omitempty"`
	Reputation        *SecurityHostReputation     `json:"reputation,omitempty"`
	Trackers          *[]SecurityHostTracker      `json:"trackers,omitempty"`
}
