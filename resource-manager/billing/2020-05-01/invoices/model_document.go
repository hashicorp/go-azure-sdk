package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Document struct {
	Kind   *DocumentType   `json:"kind,omitempty"`
	Source *DocumentSource `json:"source,omitempty"`
	Url    *string         `json:"url,omitempty"`
}
