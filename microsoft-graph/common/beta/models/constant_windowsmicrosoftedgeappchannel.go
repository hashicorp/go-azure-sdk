package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMicrosoftEdgeAppChannel string

const (
	WindowsMicrosoftEdgeAppChannelbeta   WindowsMicrosoftEdgeAppChannel = "Beta"
	WindowsMicrosoftEdgeAppChanneldev    WindowsMicrosoftEdgeAppChannel = "Dev"
	WindowsMicrosoftEdgeAppChannelstable WindowsMicrosoftEdgeAppChannel = "Stable"
)

func PossibleValuesForWindowsMicrosoftEdgeAppChannel() []string {
	return []string{
		string(WindowsMicrosoftEdgeAppChannelbeta),
		string(WindowsMicrosoftEdgeAppChanneldev),
		string(WindowsMicrosoftEdgeAppChannelstable),
	}
}

func parseWindowsMicrosoftEdgeAppChannel(input string) (*WindowsMicrosoftEdgeAppChannel, error) {
	vals := map[string]WindowsMicrosoftEdgeAppChannel{
		"beta":   WindowsMicrosoftEdgeAppChannelbeta,
		"dev":    WindowsMicrosoftEdgeAppChanneldev,
		"stable": WindowsMicrosoftEdgeAppChannelstable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMicrosoftEdgeAppChannel(input)
	return &out, nil
}
