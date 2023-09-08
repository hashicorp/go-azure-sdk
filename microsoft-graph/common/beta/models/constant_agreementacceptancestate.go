package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementAcceptanceState string

const (
	AgreementAcceptanceStateaccepted AgreementAcceptanceState = "Accepted"
	AgreementAcceptanceStatedeclined AgreementAcceptanceState = "Declined"
)

func PossibleValuesForAgreementAcceptanceState() []string {
	return []string{
		string(AgreementAcceptanceStateaccepted),
		string(AgreementAcceptanceStatedeclined),
	}
}

func parseAgreementAcceptanceState(input string) (*AgreementAcceptanceState, error) {
	vals := map[string]AgreementAcceptanceState{
		"accepted": AgreementAcceptanceStateaccepted,
		"declined": AgreementAcceptanceStatedeclined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgreementAcceptanceState(input)
	return &out, nil
}
