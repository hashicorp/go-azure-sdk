package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageSeverity string

const (
	ServiceUpdateMessageSeverity_Critical ServiceUpdateMessageSeverity = "critical"
	ServiceUpdateMessageSeverity_High     ServiceUpdateMessageSeverity = "high"
	ServiceUpdateMessageSeverity_Normal   ServiceUpdateMessageSeverity = "normal"
)

func PossibleValuesForServiceUpdateMessageSeverity() []string {
	return []string{
		string(ServiceUpdateMessageSeverity_Critical),
		string(ServiceUpdateMessageSeverity_High),
		string(ServiceUpdateMessageSeverity_Normal),
	}
}

func (s *ServiceUpdateMessageSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceUpdateMessageSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceUpdateMessageSeverity(input string) (*ServiceUpdateMessageSeverity, error) {
	vals := map[string]ServiceUpdateMessageSeverity{
		"critical": ServiceUpdateMessageSeverity_Critical,
		"high":     ServiceUpdateMessageSeverity_High,
		"normal":   ServiceUpdateMessageSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateMessageSeverity(input)
	return &out, nil
}
