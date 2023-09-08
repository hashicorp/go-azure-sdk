package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlterationResponseQueryAlterationType string

const (
	AlterationResponseQueryAlterationTypemodification AlterationResponseQueryAlterationType = "Modification"
	AlterationResponseQueryAlterationTypesuggestion   AlterationResponseQueryAlterationType = "Suggestion"
)

func PossibleValuesForAlterationResponseQueryAlterationType() []string {
	return []string{
		string(AlterationResponseQueryAlterationTypemodification),
		string(AlterationResponseQueryAlterationTypesuggestion),
	}
}

func parseAlterationResponseQueryAlterationType(input string) (*AlterationResponseQueryAlterationType, error) {
	vals := map[string]AlterationResponseQueryAlterationType{
		"modification": AlterationResponseQueryAlterationTypemodification,
		"suggestion":   AlterationResponseQueryAlterationTypesuggestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlterationResponseQueryAlterationType(input)
	return &out, nil
}
