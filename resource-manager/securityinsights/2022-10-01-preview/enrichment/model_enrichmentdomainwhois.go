package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhois struct {
	Created     *string                       `json:"created,omitempty"`
	Domain      *string                       `json:"domain,omitempty"`
	Expires     *string                       `json:"expires,omitempty"`
	ParsedWhois *EnrichmentDomainWhoisDetails `json:"parsedWhois,omitempty"`
	Server      *string                       `json:"server,omitempty"`
	Updated     *string                       `json:"updated,omitempty"`
}
