package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationFeedOrientation string

const (
	PrintJobConfigurationFeedOrientation_LongEdgeFirst  PrintJobConfigurationFeedOrientation = "longEdgeFirst"
	PrintJobConfigurationFeedOrientation_ShortEdgeFirst PrintJobConfigurationFeedOrientation = "shortEdgeFirst"
)

func PossibleValuesForPrintJobConfigurationFeedOrientation() []string {
	return []string{
		string(PrintJobConfigurationFeedOrientation_LongEdgeFirst),
		string(PrintJobConfigurationFeedOrientation_ShortEdgeFirst),
	}
}

func (s *PrintJobConfigurationFeedOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationFeedOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationFeedOrientation(input string) (*PrintJobConfigurationFeedOrientation, error) {
	vals := map[string]PrintJobConfigurationFeedOrientation{
		"longedgefirst":  PrintJobConfigurationFeedOrientation_LongEdgeFirst,
		"shortedgefirst": PrintJobConfigurationFeedOrientation_ShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFeedOrientation(input)
	return &out, nil
}
