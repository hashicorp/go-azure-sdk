package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientUserAgentPlatform string

const (
	CallRecordsClientUserAgentPlatform_Android    CallRecordsClientUserAgentPlatform = "android"
	CallRecordsClientUserAgentPlatform_HoloLens   CallRecordsClientUserAgentPlatform = "holoLens"
	CallRecordsClientUserAgentPlatform_IOS        CallRecordsClientUserAgentPlatform = "iOS"
	CallRecordsClientUserAgentPlatform_IpPhone    CallRecordsClientUserAgentPlatform = "ipPhone"
	CallRecordsClientUserAgentPlatform_MacOS      CallRecordsClientUserAgentPlatform = "macOS"
	CallRecordsClientUserAgentPlatform_RoomSystem CallRecordsClientUserAgentPlatform = "roomSystem"
	CallRecordsClientUserAgentPlatform_SurfaceHub CallRecordsClientUserAgentPlatform = "surfaceHub"
	CallRecordsClientUserAgentPlatform_Unknown    CallRecordsClientUserAgentPlatform = "unknown"
	CallRecordsClientUserAgentPlatform_Web        CallRecordsClientUserAgentPlatform = "web"
	CallRecordsClientUserAgentPlatform_Windows    CallRecordsClientUserAgentPlatform = "windows"
)

func PossibleValuesForCallRecordsClientUserAgentPlatform() []string {
	return []string{
		string(CallRecordsClientUserAgentPlatform_Android),
		string(CallRecordsClientUserAgentPlatform_HoloLens),
		string(CallRecordsClientUserAgentPlatform_IOS),
		string(CallRecordsClientUserAgentPlatform_IpPhone),
		string(CallRecordsClientUserAgentPlatform_MacOS),
		string(CallRecordsClientUserAgentPlatform_RoomSystem),
		string(CallRecordsClientUserAgentPlatform_SurfaceHub),
		string(CallRecordsClientUserAgentPlatform_Unknown),
		string(CallRecordsClientUserAgentPlatform_Web),
		string(CallRecordsClientUserAgentPlatform_Windows),
	}
}

func (s *CallRecordsClientUserAgentPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsClientUserAgentPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsClientUserAgentPlatform(input string) (*CallRecordsClientUserAgentPlatform, error) {
	vals := map[string]CallRecordsClientUserAgentPlatform{
		"android":    CallRecordsClientUserAgentPlatform_Android,
		"hololens":   CallRecordsClientUserAgentPlatform_HoloLens,
		"ios":        CallRecordsClientUserAgentPlatform_IOS,
		"ipphone":    CallRecordsClientUserAgentPlatform_IpPhone,
		"macos":      CallRecordsClientUserAgentPlatform_MacOS,
		"roomsystem": CallRecordsClientUserAgentPlatform_RoomSystem,
		"surfacehub": CallRecordsClientUserAgentPlatform_SurfaceHub,
		"unknown":    CallRecordsClientUserAgentPlatform_Unknown,
		"web":        CallRecordsClientUserAgentPlatform_Web,
		"windows":    CallRecordsClientUserAgentPlatform_Windows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsClientUserAgentPlatform(input)
	return &out, nil
}
