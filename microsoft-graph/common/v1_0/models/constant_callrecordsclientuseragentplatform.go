package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientUserAgentPlatform string

const (
	CallRecordsClientUserAgentPlatformandroid    CallRecordsClientUserAgentPlatform = "Android"
	CallRecordsClientUserAgentPlatformholoLens   CallRecordsClientUserAgentPlatform = "HoloLens"
	CallRecordsClientUserAgentPlatformiOS        CallRecordsClientUserAgentPlatform = "IOS"
	CallRecordsClientUserAgentPlatformipPhone    CallRecordsClientUserAgentPlatform = "IpPhone"
	CallRecordsClientUserAgentPlatformmacOS      CallRecordsClientUserAgentPlatform = "MacOS"
	CallRecordsClientUserAgentPlatformroomSystem CallRecordsClientUserAgentPlatform = "RoomSystem"
	CallRecordsClientUserAgentPlatformsurfaceHub CallRecordsClientUserAgentPlatform = "SurfaceHub"
	CallRecordsClientUserAgentPlatformunknown    CallRecordsClientUserAgentPlatform = "Unknown"
	CallRecordsClientUserAgentPlatformweb        CallRecordsClientUserAgentPlatform = "Web"
	CallRecordsClientUserAgentPlatformwindows    CallRecordsClientUserAgentPlatform = "Windows"
)

func PossibleValuesForCallRecordsClientUserAgentPlatform() []string {
	return []string{
		string(CallRecordsClientUserAgentPlatformandroid),
		string(CallRecordsClientUserAgentPlatformholoLens),
		string(CallRecordsClientUserAgentPlatformiOS),
		string(CallRecordsClientUserAgentPlatformipPhone),
		string(CallRecordsClientUserAgentPlatformmacOS),
		string(CallRecordsClientUserAgentPlatformroomSystem),
		string(CallRecordsClientUserAgentPlatformsurfaceHub),
		string(CallRecordsClientUserAgentPlatformunknown),
		string(CallRecordsClientUserAgentPlatformweb),
		string(CallRecordsClientUserAgentPlatformwindows),
	}
}

func parseCallRecordsClientUserAgentPlatform(input string) (*CallRecordsClientUserAgentPlatform, error) {
	vals := map[string]CallRecordsClientUserAgentPlatform{
		"android":    CallRecordsClientUserAgentPlatformandroid,
		"hololens":   CallRecordsClientUserAgentPlatformholoLens,
		"ios":        CallRecordsClientUserAgentPlatformiOS,
		"ipphone":    CallRecordsClientUserAgentPlatformipPhone,
		"macos":      CallRecordsClientUserAgentPlatformmacOS,
		"roomsystem": CallRecordsClientUserAgentPlatformroomSystem,
		"surfacehub": CallRecordsClientUserAgentPlatformsurfaceHub,
		"unknown":    CallRecordsClientUserAgentPlatformunknown,
		"web":        CallRecordsClientUserAgentPlatformweb,
		"windows":    CallRecordsClientUserAgentPlatformwindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsClientUserAgentPlatform(input)
	return &out, nil
}
