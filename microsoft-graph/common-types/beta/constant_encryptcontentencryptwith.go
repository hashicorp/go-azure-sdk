package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptContentEncryptWith string

const (
	EncryptContentEncryptWith_Template          EncryptContentEncryptWith = "template"
	EncryptContentEncryptWith_UserDefinedRights EncryptContentEncryptWith = "userDefinedRights"
)

func PossibleValuesForEncryptContentEncryptWith() []string {
	return []string{
		string(EncryptContentEncryptWith_Template),
		string(EncryptContentEncryptWith_UserDefinedRights),
	}
}

func (s *EncryptContentEncryptWith) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptContentEncryptWith(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptContentEncryptWith(input string) (*EncryptContentEncryptWith, error) {
	vals := map[string]EncryptContentEncryptWith{
		"template":          EncryptContentEncryptWith_Template,
		"userdefinedrights": EncryptContentEncryptWith_UserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptContentEncryptWith(input)
	return &out, nil
}
