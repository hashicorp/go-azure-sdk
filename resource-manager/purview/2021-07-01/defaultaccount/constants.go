package defaultaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeType string

const (
	ScopeTypeSubscription ScopeType = "Subscription"
	ScopeTypeTenant       ScopeType = "Tenant"
)

func PossibleValuesForScopeType() []string {
	return []string{
		string(ScopeTypeSubscription),
		string(ScopeTypeTenant),
	}
}
