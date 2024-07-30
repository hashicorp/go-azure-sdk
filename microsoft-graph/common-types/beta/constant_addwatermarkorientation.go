package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermarkOrientation string

const (
	AddWatermarkOrientation_Diagonal   AddWatermarkOrientation = "diagonal"
	AddWatermarkOrientation_Horizontal AddWatermarkOrientation = "horizontal"
)

func PossibleValuesForAddWatermarkOrientation() []string {
	return []string{
		string(AddWatermarkOrientation_Diagonal),
		string(AddWatermarkOrientation_Horizontal),
	}
}

func (s *AddWatermarkOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddWatermarkOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddWatermarkOrientation(input string) (*AddWatermarkOrientation, error) {
	vals := map[string]AddWatermarkOrientation{
		"diagonal":   AddWatermarkOrientation_Diagonal,
		"horizontal": AddWatermarkOrientation_Horizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddWatermarkOrientation(input)
	return &out, nil
}
