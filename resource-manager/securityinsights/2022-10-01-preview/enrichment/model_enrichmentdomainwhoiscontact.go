package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentDomainWhoisContact struct {
	City    *string   `json:"city,omitempty"`
	Country *string   `json:"country,omitempty"`
	Email   *string   `json:"email,omitempty"`
	Fax     *string   `json:"fax,omitempty"`
	Name    *string   `json:"name,omitempty"`
	Org     *string   `json:"org,omitempty"`
	Phone   *string   `json:"phone,omitempty"`
	Postal  *string   `json:"postal,omitempty"`
	State   *string   `json:"state,omitempty"`
	Street  *[]string `json:"street,omitempty"`
}
