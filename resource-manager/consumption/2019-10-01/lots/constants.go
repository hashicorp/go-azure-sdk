package lots

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotSource string

const (
	LotSourcePromotionalCredit LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit   LotSource = "PurchasedCredit"
)

func PossibleValuesForLotSource() []string {
	return []string{
		string(LotSourcePromotionalCredit),
		string(LotSourcePurchasedCredit),
	}
}

func (s *LotSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLotSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLotSource(input string) (*LotSource, error) {
	vals := map[string]LotSource{
		"promotionalcredit": LotSourcePromotionalCredit,
		"purchasedcredit":   LotSourcePurchasedCredit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LotSource(input)
	return &out, nil
}
