package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAllowedCombinationsResult struct {
	AdditionalInformation       *string                                                `json:"additionalInformation,omitempty"`
	ConditionalAccessReferences *[]string                                              `json:"conditionalAccessReferences,omitempty"`
	CurrentCombinations         *[]UpdateAllowedCombinationsResultCurrentCombinations  `json:"currentCombinations,omitempty"`
	ODataType                   *string                                                `json:"@odata.type,omitempty"`
	PreviousCombinations        *[]UpdateAllowedCombinationsResultPreviousCombinations `json:"previousCombinations,omitempty"`
}
