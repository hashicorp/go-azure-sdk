package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoredEmailAddress struct {
	Address             *string                                `json:"address,omitempty"`
	ItemId              *string                                `json:"itemId,omitempty"`
	ODataType           *string                                `json:"@odata.type,omitempty"`
	RelevanceScore      *float64                               `json:"relevanceScore,omitempty"`
	SelectionLikelihood *ScoredEmailAddressSelectionLikelihood `json:"selectionLikelihood,omitempty"`
}
