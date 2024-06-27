package invoice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceDocument struct {
	DocumentNumbers *[]string            `json:"documentNumbers,omitempty"`
	ExternalUrl     *string              `json:"externalUrl,omitempty"`
	Kind            *InvoiceDocumentType `json:"kind,omitempty"`
	Name            *string              `json:"name,omitempty"`
	Source          *DocumentSource      `json:"source,omitempty"`
	Url             *string              `json:"url,omitempty"`
}
