package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDnsTxtRecord struct {
	Id               *string `json:"id,omitempty"`
	IsOptional       *bool   `json:"isOptional,omitempty"`
	Label            *string `json:"label,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	RecordType       *string `json:"recordType,omitempty"`
	SupportedService *string `json:"supportedService,omitempty"`
	Text             *string `json:"text,omitempty"`
	Ttl              *int64  `json:"ttl,omitempty"`
}
