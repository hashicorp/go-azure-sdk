package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDnsSrvRecord struct {
	Id               *string `json:"id,omitempty"`
	IsOptional       *bool   `json:"isOptional,omitempty"`
	Label            *string `json:"label,omitempty"`
	NameTarget       *string `json:"nameTarget,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Port             *int64  `json:"port,omitempty"`
	Priority         *int64  `json:"priority,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	RecordType       *string `json:"recordType,omitempty"`
	Service          *string `json:"service,omitempty"`
	SupportedService *string `json:"supportedService,omitempty"`
	Ttl              *int64  `json:"ttl,omitempty"`
	Weight           *int64  `json:"weight,omitempty"`
}
