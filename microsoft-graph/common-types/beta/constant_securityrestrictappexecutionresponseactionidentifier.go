package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRestrictAppExecutionResponseActionIdentifier string

const (
	SecurityRestrictAppExecutionResponseActionIdentifier_DeviceId SecurityRestrictAppExecutionResponseActionIdentifier = "deviceId"
)

func PossibleValuesForSecurityRestrictAppExecutionResponseActionIdentifier() []string {
	return []string{
		string(SecurityRestrictAppExecutionResponseActionIdentifier_DeviceId),
	}
}

func (s *SecurityRestrictAppExecutionResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRestrictAppExecutionResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRestrictAppExecutionResponseActionIdentifier(input string) (*SecurityRestrictAppExecutionResponseActionIdentifier, error) {
	vals := map[string]SecurityRestrictAppExecutionResponseActionIdentifier{
		"deviceid": SecurityRestrictAppExecutionResponseActionIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRestrictAppExecutionResponseActionIdentifier(input)
	return &out, nil
}
