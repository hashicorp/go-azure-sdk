package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunAntivirusScanResponseActionIdentifier string

const (
	SecurityRunAntivirusScanResponseActionIdentifier_DeviceId SecurityRunAntivirusScanResponseActionIdentifier = "deviceId"
)

func PossibleValuesForSecurityRunAntivirusScanResponseActionIdentifier() []string {
	return []string{
		string(SecurityRunAntivirusScanResponseActionIdentifier_DeviceId),
	}
}

func (s *SecurityRunAntivirusScanResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRunAntivirusScanResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRunAntivirusScanResponseActionIdentifier(input string) (*SecurityRunAntivirusScanResponseActionIdentifier, error) {
	vals := map[string]SecurityRunAntivirusScanResponseActionIdentifier{
		"deviceid": SecurityRunAntivirusScanResponseActionIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunAntivirusScanResponseActionIdentifier(input)
	return &out, nil
}
