package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptWithUserDefinedRightsEncryptWith string

const (
	EncryptWithUserDefinedRightsEncryptWith_Template          EncryptWithUserDefinedRightsEncryptWith = "template"
	EncryptWithUserDefinedRightsEncryptWith_UserDefinedRights EncryptWithUserDefinedRightsEncryptWith = "userDefinedRights"
)

func PossibleValuesForEncryptWithUserDefinedRightsEncryptWith() []string {
	return []string{
		string(EncryptWithUserDefinedRightsEncryptWith_Template),
		string(EncryptWithUserDefinedRightsEncryptWith_UserDefinedRights),
	}
}

func (s *EncryptWithUserDefinedRightsEncryptWith) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptWithUserDefinedRightsEncryptWith(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptWithUserDefinedRightsEncryptWith(input string) (*EncryptWithUserDefinedRightsEncryptWith, error) {
	vals := map[string]EncryptWithUserDefinedRightsEncryptWith{
		"template":          EncryptWithUserDefinedRightsEncryptWith_Template,
		"userdefinedrights": EncryptWithUserDefinedRightsEncryptWith_UserDefinedRights,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptWithUserDefinedRightsEncryptWith(input)
	return &out, nil
}
