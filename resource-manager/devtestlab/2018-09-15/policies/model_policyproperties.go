package policies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyProperties struct {
	CreatedDate       *string              `json:"createdDate,omitempty"`
	Description       *string              `json:"description,omitempty"`
	EvaluatorType     *PolicyEvaluatorType `json:"evaluatorType,omitempty"`
	FactData          *string              `json:"factData,omitempty"`
	FactName          *PolicyFactName      `json:"factName,omitempty"`
	ProvisioningState *string              `json:"provisioningState,omitempty"`
	Status            *PolicyStatus        `json:"status,omitempty"`
	Threshold         *string              `json:"threshold,omitempty"`
	UniqueIdentifier  *string              `json:"uniqueIdentifier,omitempty"`
}
