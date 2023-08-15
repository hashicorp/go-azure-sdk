package reservationsummaries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Datagrain string

const (
	DatagrainDaily   Datagrain = "daily"
	DatagrainMonthly Datagrain = "monthly"
)

func PossibleValuesForDatagrain() []string {
	return []string{
		string(DatagrainDaily),
		string(DatagrainMonthly),
	}
}

func (s *Datagrain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatagrain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatagrain(input string) (*Datagrain, error) {
	vals := map[string]Datagrain{
		"daily":   DatagrainDaily,
		"monthly": DatagrainMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Datagrain(input)
	return &out, nil
}
