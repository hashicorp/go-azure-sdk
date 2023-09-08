package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithUserDefinedRightsEncryptWith string

const (
	EncryptWithUserDefinedRightsEncryptWithtemplate          EncryptWithUserDefinedRightsEncryptWith = "Template"
	EncryptWithUserDefinedRightsEncryptWithuserDefinedRights EncryptWithUserDefinedRightsEncryptWith = "UserDefinedRights"
)

func PossibleValuesForEncryptWithUserDefinedRightsEncryptWith() []string {
	return []string{
		string(EncryptWithUserDefinedRightsEncryptWithtemplate),
		string(EncryptWithUserDefinedRightsEncryptWithuserDefinedRights),
	}
}

func parseEncryptWithUserDefinedRightsEncryptWith(input string) (*EncryptWithUserDefinedRightsEncryptWith, error) {
	vals := map[string]EncryptWithUserDefinedRightsEncryptWith{
		"template":          EncryptWithUserDefinedRightsEncryptWithtemplate,
		"userdefinedrights": EncryptWithUserDefinedRightsEncryptWithuserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptWithUserDefinedRightsEncryptWith(input)
	return &out, nil
}
