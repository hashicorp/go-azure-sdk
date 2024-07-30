package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCustodianStatus string

const (
	SecurityEdiscoveryCustodianStatus_Active   SecurityEdiscoveryCustodianStatus = "active"
	SecurityEdiscoveryCustodianStatus_Released SecurityEdiscoveryCustodianStatus = "released"
)

func PossibleValuesForSecurityEdiscoveryCustodianStatus() []string {
	return []string{
		string(SecurityEdiscoveryCustodianStatus_Active),
		string(SecurityEdiscoveryCustodianStatus_Released),
	}
}

func (s *SecurityEdiscoveryCustodianStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryCustodianStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryCustodianStatus(input string) (*SecurityEdiscoveryCustodianStatus, error) {
	vals := map[string]SecurityEdiscoveryCustodianStatus{
		"active":   SecurityEdiscoveryCustodianStatus_Active,
		"released": SecurityEdiscoveryCustodianStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCustodianStatus(input)
	return &out, nil
}
