package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreSet struct {
	Children        *[]TermStoreTerm          `json:"children,omitempty"`
	CreatedDateTime *string                   `json:"createdDateTime,omitempty"`
	Description     *string                   `json:"description,omitempty"`
	Id              *string                   `json:"id,omitempty"`
	LocalizedNames  *[]TermStoreLocalizedName `json:"localizedNames,omitempty"`
	ODataType       *string                   `json:"@odata.type,omitempty"`
	ParentGroup     *TermStoreGroup           `json:"parentGroup,omitempty"`
	Properties      *[]KeyValue               `json:"properties,omitempty"`
	Relations       *[]TermStoreRelation      `json:"relations,omitempty"`
	Terms           *[]TermStoreTerm          `json:"terms,omitempty"`
}
