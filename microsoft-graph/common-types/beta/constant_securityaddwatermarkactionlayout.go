package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddWatermarkActionLayout string

const (
	SecurityAddWatermarkActionLayout_Diagonal   SecurityAddWatermarkActionLayout = "diagonal"
	SecurityAddWatermarkActionLayout_Horizontal SecurityAddWatermarkActionLayout = "horizontal"
)

func PossibleValuesForSecurityAddWatermarkActionLayout() []string {
	return []string{
		string(SecurityAddWatermarkActionLayout_Diagonal),
		string(SecurityAddWatermarkActionLayout_Horizontal),
	}
}

func (s *SecurityAddWatermarkActionLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAddWatermarkActionLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAddWatermarkActionLayout(input string) (*SecurityAddWatermarkActionLayout, error) {
	vals := map[string]SecurityAddWatermarkActionLayout{
		"diagonal":   SecurityAddWatermarkActionLayout_Diagonal,
		"horizontal": SecurityAddWatermarkActionLayout_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddWatermarkActionLayout(input)
	return &out, nil
}
