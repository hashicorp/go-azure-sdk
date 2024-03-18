package organizations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSource string

const (
	BillingSourceAZURE    BillingSource = "AZURE"
	BillingSourceNEWRELIC BillingSource = "NEWRELIC"
)

func PossibleValuesForBillingSource() []string {
	return []string{
		string(BillingSourceAZURE),
		string(BillingSourceNEWRELIC),
	}
}

func (s *BillingSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSource(input string) (*BillingSource, error) {
	vals := map[string]BillingSource{
		"azure":    BillingSourceAZURE,
		"newrelic": BillingSourceNEWRELIC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSource(input)
	return &out, nil
}
