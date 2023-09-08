package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermColumn struct {
	AllowMultipleValues    *bool          `json:"allowMultipleValues,omitempty"`
	ODataType              *string        `json:"@odata.type,omitempty"`
	ParentTerm             *TermStoreTerm `json:"parentTerm,omitempty"`
	ShowFullyQualifiedName *bool          `json:"showFullyQualifiedName,omitempty"`
	TermSet                *TermStoreSet  `json:"termSet,omitempty"`
}
