package commitmenttiers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostingModel string

const (
	HostingModelConnectedContainer    HostingModel = "ConnectedContainer"
	HostingModelDisconnectedContainer HostingModel = "DisconnectedContainer"
	HostingModelProvisionedWeb        HostingModel = "ProvisionedWeb"
	HostingModelWeb                   HostingModel = "Web"
)

func PossibleValuesForHostingModel() []string {
	return []string{
		string(HostingModelConnectedContainer),
		string(HostingModelDisconnectedContainer),
		string(HostingModelProvisionedWeb),
		string(HostingModelWeb),
	}
}

func (s *HostingModel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostingModel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostingModel(input string) (*HostingModel, error) {
	vals := map[string]HostingModel{
		"connectedcontainer":    HostingModelConnectedContainer,
		"disconnectedcontainer": HostingModelDisconnectedContainer,
		"provisionedweb":        HostingModelProvisionedWeb,
		"web":                   HostingModelWeb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostingModel(input)
	return &out, nil
}
