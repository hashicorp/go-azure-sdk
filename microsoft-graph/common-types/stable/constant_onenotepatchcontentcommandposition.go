package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommandPosition string

const (
	OnenotePatchContentCommandPosition_After  OnenotePatchContentCommandPosition = "After"
	OnenotePatchContentCommandPosition_Before OnenotePatchContentCommandPosition = "Before"
)

func PossibleValuesForOnenotePatchContentCommandPosition() []string {
	return []string{
		string(OnenotePatchContentCommandPosition_After),
		string(OnenotePatchContentCommandPosition_Before),
	}
}

func (s *OnenotePatchContentCommandPosition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenotePatchContentCommandPosition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenotePatchContentCommandPosition(input string) (*OnenotePatchContentCommandPosition, error) {
	vals := map[string]OnenotePatchContentCommandPosition{
		"after":  OnenotePatchContentCommandPosition_After,
		"before": OnenotePatchContentCommandPosition_Before,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchContentCommandPosition(input)
	return &out, nil
}
