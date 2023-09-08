package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessEnumeratedExternalTenantsMembershipKind string

const (
	ConditionalAccessEnumeratedExternalTenantsMembershipKindall        ConditionalAccessEnumeratedExternalTenantsMembershipKind = "All"
	ConditionalAccessEnumeratedExternalTenantsMembershipKindenumerated ConditionalAccessEnumeratedExternalTenantsMembershipKind = "Enumerated"
)

func PossibleValuesForConditionalAccessEnumeratedExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessEnumeratedExternalTenantsMembershipKindall),
		string(ConditionalAccessEnumeratedExternalTenantsMembershipKindenumerated),
	}
}

func parseConditionalAccessEnumeratedExternalTenantsMembershipKind(input string) (*ConditionalAccessEnumeratedExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessEnumeratedExternalTenantsMembershipKind{
		"all":        ConditionalAccessEnumeratedExternalTenantsMembershipKindall,
		"enumerated": ConditionalAccessEnumeratedExternalTenantsMembershipKindenumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessEnumeratedExternalTenantsMembershipKind(input)
	return &out, nil
}
