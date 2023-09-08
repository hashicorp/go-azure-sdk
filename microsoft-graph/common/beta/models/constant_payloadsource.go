package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadSource string

const (
	PayloadSourceglobal  PayloadSource = "Global"
	PayloadSourcetenant  PayloadSource = "Tenant"
	PayloadSourceunknown PayloadSource = "Unknown"
)

func PossibleValuesForPayloadSource() []string {
	return []string{
		string(PayloadSourceglobal),
		string(PayloadSourcetenant),
		string(PayloadSourceunknown),
	}
}

func parsePayloadSource(input string) (*PayloadSource, error) {
	vals := map[string]PayloadSource{
		"global":  PayloadSourceglobal,
		"tenant":  PayloadSourcetenant,
		"unknown": PayloadSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadSource(input)
	return &out, nil
}
