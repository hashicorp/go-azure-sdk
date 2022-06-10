package customapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WsdlDefinition struct {
	Content      *string           `json:"content,omitempty"`
	ImportMethod *WsdlImportMethod `json:"importMethod,omitempty"`
	Service      *WsdlService      `json:"service,omitempty"`
	Url          *string           `json:"url,omitempty"`
}
