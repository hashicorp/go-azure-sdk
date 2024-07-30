package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectGroupPrivacy string

const (
	ProtectGroupPrivacy_Private     ProtectGroupPrivacy = "private"
	ProtectGroupPrivacy_Public      ProtectGroupPrivacy = "public"
	ProtectGroupPrivacy_Unspecified ProtectGroupPrivacy = "unspecified"
)

func PossibleValuesForProtectGroupPrivacy() []string {
	return []string{
		string(ProtectGroupPrivacy_Private),
		string(ProtectGroupPrivacy_Public),
		string(ProtectGroupPrivacy_Unspecified),
	}
}

func (s *ProtectGroupPrivacy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectGroupPrivacy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectGroupPrivacy(input string) (*ProtectGroupPrivacy, error) {
	vals := map[string]ProtectGroupPrivacy{
		"private":     ProtectGroupPrivacy_Private,
		"public":      ProtectGroupPrivacy_Public,
		"unspecified": ProtectGroupPrivacy_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectGroupPrivacy(input)
	return &out, nil
}
