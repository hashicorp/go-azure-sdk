package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInitiateInvestigationResponseActionIdentifier string

const (
	SecurityInitiateInvestigationResponseActionIdentifier_DeviceId SecurityInitiateInvestigationResponseActionIdentifier = "deviceId"
)

func PossibleValuesForSecurityInitiateInvestigationResponseActionIdentifier() []string {
	return []string{
		string(SecurityInitiateInvestigationResponseActionIdentifier_DeviceId),
	}
}

func (s *SecurityInitiateInvestigationResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityInitiateInvestigationResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityInitiateInvestigationResponseActionIdentifier(input string) (*SecurityInitiateInvestigationResponseActionIdentifier, error) {
	vals := map[string]SecurityInitiateInvestigationResponseActionIdentifier{
		"deviceid": SecurityInitiateInvestigationResponseActionIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityInitiateInvestigationResponseActionIdentifier(input)
	return &out, nil
}
