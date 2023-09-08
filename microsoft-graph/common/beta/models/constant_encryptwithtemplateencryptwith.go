package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithTemplateEncryptWith string

const (
	EncryptWithTemplateEncryptWithtemplate          EncryptWithTemplateEncryptWith = "Template"
	EncryptWithTemplateEncryptWithuserDefinedRights EncryptWithTemplateEncryptWith = "UserDefinedRights"
)

func PossibleValuesForEncryptWithTemplateEncryptWith() []string {
	return []string{
		string(EncryptWithTemplateEncryptWithtemplate),
		string(EncryptWithTemplateEncryptWithuserDefinedRights),
	}
}

func parseEncryptWithTemplateEncryptWith(input string) (*EncryptWithTemplateEncryptWith, error) {
	vals := map[string]EncryptWithTemplateEncryptWith{
		"template":          EncryptWithTemplateEncryptWithtemplate,
		"userdefinedrights": EncryptWithTemplateEncryptWithuserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptWithTemplateEncryptWith(input)
	return &out, nil
}
