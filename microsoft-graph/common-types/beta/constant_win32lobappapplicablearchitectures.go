package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppApplicableArchitectures string

const (
	Win32LobAppApplicableArchitectures_Arm     Win32LobAppApplicableArchitectures = "arm"
	Win32LobAppApplicableArchitectures_Arm64   Win32LobAppApplicableArchitectures = "arm64"
	Win32LobAppApplicableArchitectures_Neutral Win32LobAppApplicableArchitectures = "neutral"
	Win32LobAppApplicableArchitectures_None    Win32LobAppApplicableArchitectures = "none"
	Win32LobAppApplicableArchitectures_X64     Win32LobAppApplicableArchitectures = "x64"
	Win32LobAppApplicableArchitectures_X86     Win32LobAppApplicableArchitectures = "x86"
)

func PossibleValuesForWin32LobAppApplicableArchitectures() []string {
	return []string{
		string(Win32LobAppApplicableArchitectures_Arm),
		string(Win32LobAppApplicableArchitectures_Arm64),
		string(Win32LobAppApplicableArchitectures_Neutral),
		string(Win32LobAppApplicableArchitectures_None),
		string(Win32LobAppApplicableArchitectures_X64),
		string(Win32LobAppApplicableArchitectures_X86),
	}
}

func (s *Win32LobAppApplicableArchitectures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppApplicableArchitectures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppApplicableArchitectures(input string) (*Win32LobAppApplicableArchitectures, error) {
	vals := map[string]Win32LobAppApplicableArchitectures{
		"arm":     Win32LobAppApplicableArchitectures_Arm,
		"arm64":   Win32LobAppApplicableArchitectures_Arm64,
		"neutral": Win32LobAppApplicableArchitectures_Neutral,
		"none":    Win32LobAppApplicableArchitectures_None,
		"x64":     Win32LobAppApplicableArchitectures_X64,
		"x86":     Win32LobAppApplicableArchitectures_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppApplicableArchitectures(input)
	return &out, nil
}
