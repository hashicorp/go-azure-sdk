package commitmentplans

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostingModel string

const (
	HostingModelConnectedContainer    HostingModel = "ConnectedContainer"
	HostingModelDisconnectedContainer HostingModel = "DisconnectedContainer"
	HostingModelWeb                   HostingModel = "Web"
)

func PossibleValuesForHostingModel() []string {
	return []string{
		string(HostingModelConnectedContainer),
		string(HostingModelDisconnectedContainer),
		string(HostingModelWeb),
	}
}

func parseHostingModel(input string) (*HostingModel, error) {
	vals := map[string]HostingModel{
		"connectedcontainer":    HostingModelConnectedContainer,
		"disconnectedcontainer": HostingModelDisconnectedContainer,
		"web":                   HostingModelWeb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostingModel(input)
	return &out, nil
}
