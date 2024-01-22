package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Document struct {
	DocumentNumbers *[]string     `json:"documentNumbers,omitempty"`
	Kind            *DocumentType `json:"kind,omitempty"`
	Url             *string       `json:"url,omitempty"`
}
