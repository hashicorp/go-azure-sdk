package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftEdgeAppChannel string

const (
	MacOSMicrosoftEdgeAppChannelbeta   MacOSMicrosoftEdgeAppChannel = "Beta"
	MacOSMicrosoftEdgeAppChanneldev    MacOSMicrosoftEdgeAppChannel = "Dev"
	MacOSMicrosoftEdgeAppChannelstable MacOSMicrosoftEdgeAppChannel = "Stable"
)

func PossibleValuesForMacOSMicrosoftEdgeAppChannel() []string {
	return []string{
		string(MacOSMicrosoftEdgeAppChannelbeta),
		string(MacOSMicrosoftEdgeAppChanneldev),
		string(MacOSMicrosoftEdgeAppChannelstable),
	}
}

func parseMacOSMicrosoftEdgeAppChannel(input string) (*MacOSMicrosoftEdgeAppChannel, error) {
	vals := map[string]MacOSMicrosoftEdgeAppChannel{
		"beta":   MacOSMicrosoftEdgeAppChannelbeta,
		"dev":    MacOSMicrosoftEdgeAppChanneldev,
		"stable": MacOSMicrosoftEdgeAppChannelstable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftEdgeAppChannel(input)
	return &out, nil
}
