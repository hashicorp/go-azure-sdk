package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedPermissionClassificationClassification string

const (
	DelegatedPermissionClassificationClassification_High   DelegatedPermissionClassificationClassification = "high"
	DelegatedPermissionClassificationClassification_Low    DelegatedPermissionClassificationClassification = "low"
	DelegatedPermissionClassificationClassification_Medium DelegatedPermissionClassificationClassification = "medium"
)

func PossibleValuesForDelegatedPermissionClassificationClassification() []string {
	return []string{
		string(DelegatedPermissionClassificationClassification_High),
		string(DelegatedPermissionClassificationClassification_Low),
		string(DelegatedPermissionClassificationClassification_Medium),
	}
}

func (s *DelegatedPermissionClassificationClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedPermissionClassificationClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedPermissionClassificationClassification(input string) (*DelegatedPermissionClassificationClassification, error) {
	vals := map[string]DelegatedPermissionClassificationClassification{
		"high":   DelegatedPermissionClassificationClassification_High,
		"low":    DelegatedPermissionClassificationClassification_Low,
		"medium": DelegatedPermissionClassificationClassification_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedPermissionClassificationClassification(input)
	return &out, nil
}
