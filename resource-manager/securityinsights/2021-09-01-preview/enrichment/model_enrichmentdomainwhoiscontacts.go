package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhoisContacts struct {
	Admin      *EnrichmentDomainWhoisContact `json:"admin,omitempty"`
	Billing    *EnrichmentDomainWhoisContact `json:"billing,omitempty"`
	Registrant *EnrichmentDomainWhoisContact `json:"registrant,omitempty"`
	Tech       *EnrichmentDomainWhoisContact `json:"tech,omitempty"`
}
