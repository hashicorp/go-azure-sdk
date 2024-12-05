package billingstatistics

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingStatisticKind string

const (
	BillingStatisticKindSapSolutionUsage BillingStatisticKind = "SapSolutionUsage"
)

func PossibleValuesForBillingStatisticKind() []string {
	return []string{
		string(BillingStatisticKindSapSolutionUsage),
	}
}

func (s *BillingStatisticKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingStatisticKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingStatisticKind(input string) (*BillingStatisticKind, error) {
	vals := map[string]BillingStatisticKind{
		"sapsolutionusage": BillingStatisticKindSapSolutionUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingStatisticKind(input)
	return &out, nil
}
