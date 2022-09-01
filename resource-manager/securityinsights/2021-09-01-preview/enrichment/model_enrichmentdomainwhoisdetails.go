package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhoisDetails struct {
	Contacts    *EnrichmentDomainWhoisContacts         `json:"contacts,omitempty"`
	NameServers *[]string                              `json:"nameServers,omitempty"`
	Registrar   *EnrichmentDomainWhoisRegistrarDetails `json:"registrar,omitempty"`
	Statuses    *[]string                              `json:"statuses,omitempty"`
}
