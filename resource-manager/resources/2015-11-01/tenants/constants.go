package tenants

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceNameStatus string

const (
	ResourceNameStatusAllowed  ResourceNameStatus = "Allowed"
	ResourceNameStatusReserved ResourceNameStatus = "Reserved"
)

func PossibleValuesForResourceNameStatus() []string {
	return []string{
		string(ResourceNameStatusAllowed),
		string(ResourceNameStatusReserved),
	}
}

func (s *ResourceNameStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceNameStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceNameStatus(input string) (*ResourceNameStatus, error) {
	vals := map[string]ResourceNameStatus{
		"allowed":  ResourceNameStatusAllowed,
		"reserved": ResourceNameStatusReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceNameStatus(input)
	return &out, nil
}
