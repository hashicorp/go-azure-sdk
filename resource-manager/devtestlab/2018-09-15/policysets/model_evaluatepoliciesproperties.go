package policysets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluatePoliciesProperties struct {
	FactData     *string `json:"factData,omitempty"`
	FactName     *string `json:"factName,omitempty"`
	UserObjectId *string `json:"userObjectId,omitempty"`
	ValueOffset  *string `json:"valueOffset,omitempty"`
}
