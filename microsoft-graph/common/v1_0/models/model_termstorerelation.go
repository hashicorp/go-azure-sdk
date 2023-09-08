package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreRelation struct {
	FromTerm     *TermStoreTerm                 `json:"fromTerm,omitempty"`
	Id           *string                        `json:"id,omitempty"`
	ODataType    *string                        `json:"@odata.type,omitempty"`
	Relationship *TermStoreRelationRelationship `json:"relationship,omitempty"`
	Set          *TermStoreSet                  `json:"set,omitempty"`
	ToTerm       *TermStoreTerm                 `json:"toTerm,omitempty"`
}
