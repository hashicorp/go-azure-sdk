package policy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySummary struct {
	Name       *string     `json:"name,omitempty"`
	PolicyType *PolicyType `json:"policyType,omitempty"`
	Scope      *string     `json:"scope,omitempty"`
	Value      *string     `json:"value,omitempty"`
}
