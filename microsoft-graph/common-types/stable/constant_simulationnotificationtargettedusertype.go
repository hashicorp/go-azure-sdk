package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationNotificationTargettedUserType string

const (
	SimulationNotificationTargettedUserType_AllUsers    SimulationNotificationTargettedUserType = "allUsers"
	SimulationNotificationTargettedUserType_Clicked     SimulationNotificationTargettedUserType = "clicked"
	SimulationNotificationTargettedUserType_Compromised SimulationNotificationTargettedUserType = "compromised"
	SimulationNotificationTargettedUserType_Unknown     SimulationNotificationTargettedUserType = "unknown"
)

func PossibleValuesForSimulationNotificationTargettedUserType() []string {
	return []string{
		string(SimulationNotificationTargettedUserType_AllUsers),
		string(SimulationNotificationTargettedUserType_Clicked),
		string(SimulationNotificationTargettedUserType_Compromised),
		string(SimulationNotificationTargettedUserType_Unknown),
	}
}

func (s *SimulationNotificationTargettedUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationNotificationTargettedUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationNotificationTargettedUserType(input string) (*SimulationNotificationTargettedUserType, error) {
	vals := map[string]SimulationNotificationTargettedUserType{
		"allusers":    SimulationNotificationTargettedUserType_AllUsers,
		"clicked":     SimulationNotificationTargettedUserType_Clicked,
		"compromised": SimulationNotificationTargettedUserType_Compromised,
		"unknown":     SimulationNotificationTargettedUserType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationNotificationTargettedUserType(input)
	return &out, nil
}
