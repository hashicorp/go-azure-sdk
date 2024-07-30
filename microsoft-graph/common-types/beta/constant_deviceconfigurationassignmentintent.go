package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignmentIntent string

const (
	DeviceConfigurationAssignmentIntent_Apply  DeviceConfigurationAssignmentIntent = "apply"
	DeviceConfigurationAssignmentIntent_Remove DeviceConfigurationAssignmentIntent = "remove"
)

func PossibleValuesForDeviceConfigurationAssignmentIntent() []string {
	return []string{
		string(DeviceConfigurationAssignmentIntent_Apply),
		string(DeviceConfigurationAssignmentIntent_Remove),
	}
}

func (s *DeviceConfigurationAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationAssignmentIntent(input string) (*DeviceConfigurationAssignmentIntent, error) {
	vals := map[string]DeviceConfigurationAssignmentIntent{
		"apply":  DeviceConfigurationAssignmentIntent_Apply,
		"remove": DeviceConfigurationAssignmentIntent_Remove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentIntent(input)
	return &out, nil
}
