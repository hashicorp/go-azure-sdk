package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileAssignmentIntent string

const (
	DeviceManagementResourceAccessProfileAssignmentIntent_Apply  DeviceManagementResourceAccessProfileAssignmentIntent = "apply"
	DeviceManagementResourceAccessProfileAssignmentIntent_Remove DeviceManagementResourceAccessProfileAssignmentIntent = "remove"
)

func PossibleValuesForDeviceManagementResourceAccessProfileAssignmentIntent() []string {
	return []string{
		string(DeviceManagementResourceAccessProfileAssignmentIntent_Apply),
		string(DeviceManagementResourceAccessProfileAssignmentIntent_Remove),
	}
}

func (s *DeviceManagementResourceAccessProfileAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementResourceAccessProfileAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementResourceAccessProfileAssignmentIntent(input string) (*DeviceManagementResourceAccessProfileAssignmentIntent, error) {
	vals := map[string]DeviceManagementResourceAccessProfileAssignmentIntent{
		"apply":  DeviceManagementResourceAccessProfileAssignmentIntent_Apply,
		"remove": DeviceManagementResourceAccessProfileAssignmentIntent_Remove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementResourceAccessProfileAssignmentIntent(input)
	return &out, nil
}
