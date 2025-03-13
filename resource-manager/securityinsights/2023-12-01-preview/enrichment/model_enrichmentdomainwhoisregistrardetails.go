package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhoisRegistrarDetails struct {
	AbuseContactEmail *string `json:"abuseContactEmail,omitempty"`
	AbuseContactPhone *string `json:"abuseContactPhone,omitempty"`
	IanaId            *string `json:"ianaId,omitempty"`
	Name              *string `json:"name,omitempty"`
	Url               *string `json:"url,omitempty"`
	WhoisServer       *string `json:"whoisServer,omitempty"`
}
