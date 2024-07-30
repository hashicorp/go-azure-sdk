package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceDeviceType string

const (
	TeamworkDeviceDeviceType_CollaborationBar TeamworkDeviceDeviceType = "collaborationBar"
	TeamworkDeviceDeviceType_IpPhone          TeamworkDeviceDeviceType = "ipPhone"
	TeamworkDeviceDeviceType_LowCostPhone     TeamworkDeviceDeviceType = "lowCostPhone"
	TeamworkDeviceDeviceType_Sip              TeamworkDeviceDeviceType = "sip"
	TeamworkDeviceDeviceType_SurfaceHub       TeamworkDeviceDeviceType = "surfaceHub"
	TeamworkDeviceDeviceType_TeamsDisplay     TeamworkDeviceDeviceType = "teamsDisplay"
	TeamworkDeviceDeviceType_TeamsPanel       TeamworkDeviceDeviceType = "teamsPanel"
	TeamworkDeviceDeviceType_TeamsRoom        TeamworkDeviceDeviceType = "teamsRoom"
	TeamworkDeviceDeviceType_TouchConsole     TeamworkDeviceDeviceType = "touchConsole"
	TeamworkDeviceDeviceType_Unknown          TeamworkDeviceDeviceType = "unknown"
)

func PossibleValuesForTeamworkDeviceDeviceType() []string {
	return []string{
		string(TeamworkDeviceDeviceType_CollaborationBar),
		string(TeamworkDeviceDeviceType_IpPhone),
		string(TeamworkDeviceDeviceType_LowCostPhone),
		string(TeamworkDeviceDeviceType_Sip),
		string(TeamworkDeviceDeviceType_SurfaceHub),
		string(TeamworkDeviceDeviceType_TeamsDisplay),
		string(TeamworkDeviceDeviceType_TeamsPanel),
		string(TeamworkDeviceDeviceType_TeamsRoom),
		string(TeamworkDeviceDeviceType_TouchConsole),
		string(TeamworkDeviceDeviceType_Unknown),
	}
}

func (s *TeamworkDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceDeviceType(input string) (*TeamworkDeviceDeviceType, error) {
	vals := map[string]TeamworkDeviceDeviceType{
		"collaborationbar": TeamworkDeviceDeviceType_CollaborationBar,
		"ipphone":          TeamworkDeviceDeviceType_IpPhone,
		"lowcostphone":     TeamworkDeviceDeviceType_LowCostPhone,
		"sip":              TeamworkDeviceDeviceType_Sip,
		"surfacehub":       TeamworkDeviceDeviceType_SurfaceHub,
		"teamsdisplay":     TeamworkDeviceDeviceType_TeamsDisplay,
		"teamspanel":       TeamworkDeviceDeviceType_TeamsPanel,
		"teamsroom":        TeamworkDeviceDeviceType_TeamsRoom,
		"touchconsole":     TeamworkDeviceDeviceType_TouchConsole,
		"unknown":          TeamworkDeviceDeviceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceDeviceType(input)
	return &out, nil
}
