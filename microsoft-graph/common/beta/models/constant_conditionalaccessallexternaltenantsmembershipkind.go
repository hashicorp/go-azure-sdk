package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAllExternalTenantsMembershipKind string

const (
	ConditionalAccessAllExternalTenantsMembershipKindall        ConditionalAccessAllExternalTenantsMembershipKind = "All"
	ConditionalAccessAllExternalTenantsMembershipKindenumerated ConditionalAccessAllExternalTenantsMembershipKind = "Enumerated"
)

func PossibleValuesForConditionalAccessAllExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessAllExternalTenantsMembershipKindall),
		string(ConditionalAccessAllExternalTenantsMembershipKindenumerated),
	}
}

func parseConditionalAccessAllExternalTenantsMembershipKind(input string) (*ConditionalAccessAllExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessAllExternalTenantsMembershipKind{
		"all":        ConditionalAccessAllExternalTenantsMembershipKindall,
		"enumerated": ConditionalAccessAllExternalTenantsMembershipKindenumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessAllExternalTenantsMembershipKind(input)
	return &out, nil
}
