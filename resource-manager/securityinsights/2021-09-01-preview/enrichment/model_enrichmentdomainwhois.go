package enrichment

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *EnrichmentDomainWhois) GetCreatedAsTime() (*time.Time, error) {
	if o.Created == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Created, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrichmentDomainWhois) SetCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Created = &formatted
}

func (o *EnrichmentDomainWhois) GetExpiresAsTime() (*time.Time, error) {
	if o.Expires == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expires, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrichmentDomainWhois) SetExpiresAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expires = &formatted
}

func (o *EnrichmentDomainWhois) GetUpdatedAsTime() (*time.Time, error) {
	if o.Updated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Updated, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrichmentDomainWhois) SetUpdatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Updated = &formatted
}
