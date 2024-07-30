package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectSiteAccessType string

const (
	ProtectSiteAccessType_Block   ProtectSiteAccessType = "block"
	ProtectSiteAccessType_Full    ProtectSiteAccessType = "full"
	ProtectSiteAccessType_Limited ProtectSiteAccessType = "limited"
)

func PossibleValuesForProtectSiteAccessType() []string {
	return []string{
		string(ProtectSiteAccessType_Block),
		string(ProtectSiteAccessType_Full),
		string(ProtectSiteAccessType_Limited),
	}
}

func (s *ProtectSiteAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectSiteAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectSiteAccessType(input string) (*ProtectSiteAccessType, error) {
	vals := map[string]ProtectSiteAccessType{
		"block":   ProtectSiteAccessType_Block,
		"full":    ProtectSiteAccessType_Full,
		"limited": ProtectSiteAccessType_Limited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectSiteAccessType(input)
	return &out, nil
}
