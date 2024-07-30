package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppRegistrationFlaggedReasons string

const (
	IosManagedAppRegistrationFlaggedReasons_None         IosManagedAppRegistrationFlaggedReasons = "none"
	IosManagedAppRegistrationFlaggedReasons_RootedDevice IosManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForIosManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(IosManagedAppRegistrationFlaggedReasons_None),
		string(IosManagedAppRegistrationFlaggedReasons_RootedDevice),
	}
}

func (s *IosManagedAppRegistrationFlaggedReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppRegistrationFlaggedReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppRegistrationFlaggedReasons(input string) (*IosManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]IosManagedAppRegistrationFlaggedReasons{
		"none":         IosManagedAppRegistrationFlaggedReasons_None,
		"rooteddevice": IosManagedAppRegistrationFlaggedReasons_RootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
