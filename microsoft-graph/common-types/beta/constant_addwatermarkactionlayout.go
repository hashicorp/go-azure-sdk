package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermarkActionLayout string

const (
	AddWatermarkActionLayout_Diagonal   AddWatermarkActionLayout = "diagonal"
	AddWatermarkActionLayout_Horizontal AddWatermarkActionLayout = "horizontal"
)

func PossibleValuesForAddWatermarkActionLayout() []string {
	return []string{
		string(AddWatermarkActionLayout_Diagonal),
		string(AddWatermarkActionLayout_Horizontal),
	}
}

func (s *AddWatermarkActionLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddWatermarkActionLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddWatermarkActionLayout(input string) (*AddWatermarkActionLayout, error) {
	vals := map[string]AddWatermarkActionLayout{
		"diagonal":   AddWatermarkActionLayout_Diagonal,
		"horizontal": AddWatermarkActionLayout_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddWatermarkActionLayout(input)
	return &out, nil
}
