package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptContentEncryptWith string

const (
	EncryptContentEncryptWithtemplate          EncryptContentEncryptWith = "Template"
	EncryptContentEncryptWithuserDefinedRights EncryptContentEncryptWith = "UserDefinedRights"
)

func PossibleValuesForEncryptContentEncryptWith() []string {
	return []string{
		string(EncryptContentEncryptWithtemplate),
		string(EncryptContentEncryptWithuserDefinedRights),
	}
}

func parseEncryptContentEncryptWith(input string) (*EncryptContentEncryptWith, error) {
	vals := map[string]EncryptContentEncryptWith{
		"template":          EncryptContentEncryptWithtemplate,
		"userdefinedrights": EncryptContentEncryptWithuserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptContentEncryptWith(input)
	return &out, nil
}
