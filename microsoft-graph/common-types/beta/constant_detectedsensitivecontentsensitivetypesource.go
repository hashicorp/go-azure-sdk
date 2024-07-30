package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentSensitiveTypeSource string

const (
	DetectedSensitiveContentSensitiveTypeSource_OutOfBox DetectedSensitiveContentSensitiveTypeSource = "outOfBox"
	DetectedSensitiveContentSensitiveTypeSource_Tenant   DetectedSensitiveContentSensitiveTypeSource = "tenant"
)

func PossibleValuesForDetectedSensitiveContentSensitiveTypeSource() []string {
	return []string{
		string(DetectedSensitiveContentSensitiveTypeSource_OutOfBox),
		string(DetectedSensitiveContentSensitiveTypeSource_Tenant),
	}
}

func (s *DetectedSensitiveContentSensitiveTypeSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedSensitiveContentSensitiveTypeSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedSensitiveContentSensitiveTypeSource(input string) (*DetectedSensitiveContentSensitiveTypeSource, error) {
	vals := map[string]DetectedSensitiveContentSensitiveTypeSource{
		"outofbox": DetectedSensitiveContentSensitiveTypeSource_OutOfBox,
		"tenant":   DetectedSensitiveContentSensitiveTypeSource_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentSensitiveTypeSource(input)
	return &out, nil
}
