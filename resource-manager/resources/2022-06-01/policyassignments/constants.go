package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnforcementMode string

const (
	EnforcementModeDefault      EnforcementMode = "Default"
	EnforcementModeDoNotEnforce EnforcementMode = "DoNotEnforce"
)

func PossibleValuesForEnforcementMode() []string {
	return []string{
		string(EnforcementModeDefault),
		string(EnforcementModeDoNotEnforce),
	}
}

type OverrideKind string

const (
	OverrideKindPolicyEffect OverrideKind = "policyEffect"
)

func PossibleValuesForOverrideKind() []string {
	return []string{
		string(OverrideKindPolicyEffect),
	}
}

type SelectorKind string

const (
	SelectorKindPolicyDefinitionReferenceId SelectorKind = "policyDefinitionReferenceId"
	SelectorKindResourceLocation            SelectorKind = "resourceLocation"
	SelectorKindResourceType                SelectorKind = "resourceType"
	SelectorKindResourceWithoutLocation     SelectorKind = "resourceWithoutLocation"
)

func PossibleValuesForSelectorKind() []string {
	return []string{
		string(SelectorKindPolicyDefinitionReferenceId),
		string(SelectorKindResourceLocation),
		string(SelectorKindResourceType),
		string(SelectorKindResourceWithoutLocation),
	}
}
