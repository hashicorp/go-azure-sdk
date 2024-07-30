package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadPlatform string

const (
	PayloadPlatform_Email   PayloadPlatform = "email"
	PayloadPlatform_Sms     PayloadPlatform = "sms"
	PayloadPlatform_Teams   PayloadPlatform = "teams"
	PayloadPlatform_Unknown PayloadPlatform = "unknown"
)

func PossibleValuesForPayloadPlatform() []string {
	return []string{
		string(PayloadPlatform_Email),
		string(PayloadPlatform_Sms),
		string(PayloadPlatform_Teams),
		string(PayloadPlatform_Unknown),
	}
}

func (s *PayloadPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadPlatform(input string) (*PayloadPlatform, error) {
	vals := map[string]PayloadPlatform{
		"email":   PayloadPlatform_Email,
		"sms":     PayloadPlatform_Sms,
		"teams":   PayloadPlatform_Teams,
		"unknown": PayloadPlatform_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadPlatform(input)
	return &out, nil
}
