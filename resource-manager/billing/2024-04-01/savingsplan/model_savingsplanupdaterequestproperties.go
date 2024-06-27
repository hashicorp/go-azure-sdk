package savingsplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanUpdateRequestProperties struct {
	AppliedScopeProperties *AppliedScopeProperties `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType       *AppliedScopeType       `json:"appliedScopeType,omitempty"`
	DisplayName            *string                 `json:"displayName,omitempty"`
	Renew                  *bool                   `json:"renew,omitempty"`
	RenewProperties        *RenewProperties        `json:"renewProperties,omitempty"`
}
