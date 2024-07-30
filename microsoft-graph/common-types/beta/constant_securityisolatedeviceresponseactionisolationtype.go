package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseActionIsolationType string

const (
	SecurityIsolateDeviceResponseActionIsolationType_Full      SecurityIsolateDeviceResponseActionIsolationType = "full"
	SecurityIsolateDeviceResponseActionIsolationType_Selective SecurityIsolateDeviceResponseActionIsolationType = "selective"
)

func PossibleValuesForSecurityIsolateDeviceResponseActionIsolationType() []string {
	return []string{
		string(SecurityIsolateDeviceResponseActionIsolationType_Full),
		string(SecurityIsolateDeviceResponseActionIsolationType_Selective),
	}
}

func (s *SecurityIsolateDeviceResponseActionIsolationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIsolateDeviceResponseActionIsolationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIsolateDeviceResponseActionIsolationType(input string) (*SecurityIsolateDeviceResponseActionIsolationType, error) {
	vals := map[string]SecurityIsolateDeviceResponseActionIsolationType{
		"full":      SecurityIsolateDeviceResponseActionIsolationType_Full,
		"selective": SecurityIsolateDeviceResponseActionIsolationType_Selective,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIsolateDeviceResponseActionIsolationType(input)
	return &out, nil
}
