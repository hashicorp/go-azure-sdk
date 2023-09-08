package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceDeviceType string

const (
	TeamworkDeviceDeviceTypecollaborationBar TeamworkDeviceDeviceType = "CollaborationBar"
	TeamworkDeviceDeviceTypeipPhone          TeamworkDeviceDeviceType = "IpPhone"
	TeamworkDeviceDeviceTypelowCostPhone     TeamworkDeviceDeviceType = "LowCostPhone"
	TeamworkDeviceDeviceTypesip              TeamworkDeviceDeviceType = "Sip"
	TeamworkDeviceDeviceTypesurfaceHub       TeamworkDeviceDeviceType = "SurfaceHub"
	TeamworkDeviceDeviceTypeteamsDisplay     TeamworkDeviceDeviceType = "TeamsDisplay"
	TeamworkDeviceDeviceTypeteamsPanel       TeamworkDeviceDeviceType = "TeamsPanel"
	TeamworkDeviceDeviceTypeteamsRoom        TeamworkDeviceDeviceType = "TeamsRoom"
	TeamworkDeviceDeviceTypetouchConsole     TeamworkDeviceDeviceType = "TouchConsole"
	TeamworkDeviceDeviceTypeunknown          TeamworkDeviceDeviceType = "Unknown"
)

func PossibleValuesForTeamworkDeviceDeviceType() []string {
	return []string{
		string(TeamworkDeviceDeviceTypecollaborationBar),
		string(TeamworkDeviceDeviceTypeipPhone),
		string(TeamworkDeviceDeviceTypelowCostPhone),
		string(TeamworkDeviceDeviceTypesip),
		string(TeamworkDeviceDeviceTypesurfaceHub),
		string(TeamworkDeviceDeviceTypeteamsDisplay),
		string(TeamworkDeviceDeviceTypeteamsPanel),
		string(TeamworkDeviceDeviceTypeteamsRoom),
		string(TeamworkDeviceDeviceTypetouchConsole),
		string(TeamworkDeviceDeviceTypeunknown),
	}
}

func parseTeamworkDeviceDeviceType(input string) (*TeamworkDeviceDeviceType, error) {
	vals := map[string]TeamworkDeviceDeviceType{
		"collaborationbar": TeamworkDeviceDeviceTypecollaborationBar,
		"ipphone":          TeamworkDeviceDeviceTypeipPhone,
		"lowcostphone":     TeamworkDeviceDeviceTypelowCostPhone,
		"sip":              TeamworkDeviceDeviceTypesip,
		"surfacehub":       TeamworkDeviceDeviceTypesurfaceHub,
		"teamsdisplay":     TeamworkDeviceDeviceTypeteamsDisplay,
		"teamspanel":       TeamworkDeviceDeviceTypeteamsPanel,
		"teamsroom":        TeamworkDeviceDeviceTypeteamsRoom,
		"touchconsole":     TeamworkDeviceDeviceTypetouchConsole,
		"unknown":          TeamworkDeviceDeviceTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceDeviceType(input)
	return &out, nil
}
