package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeSensitiveTypeSource string

const (
	SensitiveTypeSensitiveTypeSource_OutOfBox SensitiveTypeSensitiveTypeSource = "outOfBox"
	SensitiveTypeSensitiveTypeSource_Tenant   SensitiveTypeSensitiveTypeSource = "tenant"
)

func PossibleValuesForSensitiveTypeSensitiveTypeSource() []string {
	return []string{
		string(SensitiveTypeSensitiveTypeSource_OutOfBox),
		string(SensitiveTypeSensitiveTypeSource_Tenant),
	}
}

func (s *SensitiveTypeSensitiveTypeSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitiveTypeSensitiveTypeSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitiveTypeSensitiveTypeSource(input string) (*SensitiveTypeSensitiveTypeSource, error) {
	vals := map[string]SensitiveTypeSensitiveTypeSource{
		"outofbox": SensitiveTypeSensitiveTypeSource_OutOfBox,
		"tenant":   SensitiveTypeSensitiveTypeSource_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeSensitiveTypeSource(input)
	return &out, nil
}
