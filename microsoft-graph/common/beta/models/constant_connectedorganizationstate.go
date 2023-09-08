package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedOrganizationState string

const (
	ConnectedOrganizationStateconfigured ConnectedOrganizationState = "Configured"
	ConnectedOrganizationStateproposed   ConnectedOrganizationState = "Proposed"
)

func PossibleValuesForConnectedOrganizationState() []string {
	return []string{
		string(ConnectedOrganizationStateconfigured),
		string(ConnectedOrganizationStateproposed),
	}
}

func parseConnectedOrganizationState(input string) (*ConnectedOrganizationState, error) {
	vals := map[string]ConnectedOrganizationState{
		"configured": ConnectedOrganizationStateconfigured,
		"proposed":   ConnectedOrganizationStateproposed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectedOrganizationState(input)
	return &out, nil
}
