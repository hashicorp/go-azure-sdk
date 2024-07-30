package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithTemplateEncryptWith string

const (
	EncryptWithTemplateEncryptWith_Template          EncryptWithTemplateEncryptWith = "template"
	EncryptWithTemplateEncryptWith_UserDefinedRights EncryptWithTemplateEncryptWith = "userDefinedRights"
)

func PossibleValuesForEncryptWithTemplateEncryptWith() []string {
	return []string{
		string(EncryptWithTemplateEncryptWith_Template),
		string(EncryptWithTemplateEncryptWith_UserDefinedRights),
	}
}

func (s *EncryptWithTemplateEncryptWith) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptWithTemplateEncryptWith(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptWithTemplateEncryptWith(input string) (*EncryptWithTemplateEncryptWith, error) {
	vals := map[string]EncryptWithTemplateEncryptWith{
		"template":          EncryptWithTemplateEncryptWith_Template,
		"userdefinedrights": EncryptWithTemplateEncryptWith_UserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptWithTemplateEncryptWith(input)
	return &out, nil
}
