package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreTerm struct {
	Children             *[]TermStoreTerm                 `json:"children,omitempty"`
	CreatedDateTime      *string                          `json:"createdDateTime,omitempty"`
	Descriptions         *[]TermStoreLocalizedDescription `json:"descriptions,omitempty"`
	Id                   *string                          `json:"id,omitempty"`
	Labels               *[]TermStoreLocalizedLabel       `json:"labels,omitempty"`
	LastModifiedDateTime *string                          `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	Properties           *[]KeyValue                      `json:"properties,omitempty"`
	Relations            *[]TermStoreRelation             `json:"relations,omitempty"`
	Set                  *TermStoreSet                    `json:"set,omitempty"`
}
