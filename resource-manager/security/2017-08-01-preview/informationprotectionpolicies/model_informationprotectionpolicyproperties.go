package informationprotectionpolicies

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPolicyProperties struct {
	InformationTypes *map[string]InformationType  `json:"informationTypes,omitempty"`
	Labels           *map[string]SensitivityLabel `json:"labels,omitempty"`
	LastModifiedUtc  *string                      `json:"lastModifiedUtc,omitempty"`
	Version          *string                      `json:"version,omitempty"`
}

func (o *InformationProtectionPolicyProperties) GetLastModifiedUtcAsTime() (*time.Time, error) {
	if o.LastModifiedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedUtc, "2006-01-02T15:04:05Z07:00")
}
