package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseActionIdentifier string

const (
	SecurityIsolateDeviceResponseActionIdentifier_DeviceId SecurityIsolateDeviceResponseActionIdentifier = "deviceId"
)

func PossibleValuesForSecurityIsolateDeviceResponseActionIdentifier() []string {
	return []string{
		string(SecurityIsolateDeviceResponseActionIdentifier_DeviceId),
	}
}

func (s *SecurityIsolateDeviceResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIsolateDeviceResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIsolateDeviceResponseActionIdentifier(input string) (*SecurityIsolateDeviceResponseActionIdentifier, error) {
	vals := map[string]SecurityIsolateDeviceResponseActionIdentifier{
		"deviceid": SecurityIsolateDeviceResponseActionIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIsolateDeviceResponseActionIdentifier(input)
	return &out, nil
}
