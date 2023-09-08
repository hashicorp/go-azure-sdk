package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessExternalTenantsMembershipKind string

const (
	ConditionalAccessExternalTenantsMembershipKindall        ConditionalAccessExternalTenantsMembershipKind = "All"
	ConditionalAccessExternalTenantsMembershipKindenumerated ConditionalAccessExternalTenantsMembershipKind = "Enumerated"
)

func PossibleValuesForConditionalAccessExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessExternalTenantsMembershipKindall),
		string(ConditionalAccessExternalTenantsMembershipKindenumerated),
	}
}

func parseConditionalAccessExternalTenantsMembershipKind(input string) (*ConditionalAccessExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessExternalTenantsMembershipKind{
		"all":        ConditionalAccessExternalTenantsMembershipKindall,
		"enumerated": ConditionalAccessExternalTenantsMembershipKindenumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessExternalTenantsMembershipKind(input)
	return &out, nil
}
