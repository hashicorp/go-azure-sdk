package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceHealthStatus string

const (
	TeamworkDeviceHealthStatuscritical  TeamworkDeviceHealthStatus = "Critical"
	TeamworkDeviceHealthStatushealthy   TeamworkDeviceHealthStatus = "Healthy"
	TeamworkDeviceHealthStatusnonUrgent TeamworkDeviceHealthStatus = "NonUrgent"
	TeamworkDeviceHealthStatusoffline   TeamworkDeviceHealthStatus = "Offline"
	TeamworkDeviceHealthStatusunknown   TeamworkDeviceHealthStatus = "Unknown"
)

func PossibleValuesForTeamworkDeviceHealthStatus() []string {
	return []string{
		string(TeamworkDeviceHealthStatuscritical),
		string(TeamworkDeviceHealthStatushealthy),
		string(TeamworkDeviceHealthStatusnonUrgent),
		string(TeamworkDeviceHealthStatusoffline),
		string(TeamworkDeviceHealthStatusunknown),
	}
}

func parseTeamworkDeviceHealthStatus(input string) (*TeamworkDeviceHealthStatus, error) {
	vals := map[string]TeamworkDeviceHealthStatus{
		"critical":  TeamworkDeviceHealthStatuscritical,
		"healthy":   TeamworkDeviceHealthStatushealthy,
		"nonurgent": TeamworkDeviceHealthStatusnonUrgent,
		"offline":   TeamworkDeviceHealthStatusoffline,
		"unknown":   TeamworkDeviceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceHealthStatus(input)
	return &out, nil
}
