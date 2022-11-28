package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhoisContacts struct {
	Admin      *EnrichmentDomainWhoisContact `json:"admin"`
	Billing    *EnrichmentDomainWhoisContact `json:"billing"`
	Registrant *EnrichmentDomainWhoisContact `json:"registrant"`
	Tech       *EnrichmentDomainWhoisContact `json:"tech"`
}
