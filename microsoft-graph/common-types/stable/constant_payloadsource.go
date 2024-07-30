package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadSource string

const (
	PayloadSource_Global  PayloadSource = "global"
	PayloadSource_Tenant  PayloadSource = "tenant"
	PayloadSource_Unknown PayloadSource = "unknown"
)

func PossibleValuesForPayloadSource() []string {
	return []string{
		string(PayloadSource_Global),
		string(PayloadSource_Tenant),
		string(PayloadSource_Unknown),
	}
}

func (s *PayloadSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadSource(input string) (*PayloadSource, error) {
	vals := map[string]PayloadSource{
		"global":  PayloadSource_Global,
		"tenant":  PayloadSource_Tenant,
		"unknown": PayloadSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadSource(input)
	return &out, nil
}
