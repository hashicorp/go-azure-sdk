package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMicrosoftEdgeAppChannel string

const (
	WindowsMicrosoftEdgeAppChannel_Beta   WindowsMicrosoftEdgeAppChannel = "beta"
	WindowsMicrosoftEdgeAppChannel_Dev    WindowsMicrosoftEdgeAppChannel = "dev"
	WindowsMicrosoftEdgeAppChannel_Stable WindowsMicrosoftEdgeAppChannel = "stable"
)

func PossibleValuesForWindowsMicrosoftEdgeAppChannel() []string {
	return []string{
		string(WindowsMicrosoftEdgeAppChannel_Beta),
		string(WindowsMicrosoftEdgeAppChannel_Dev),
		string(WindowsMicrosoftEdgeAppChannel_Stable),
	}
}

func (s *WindowsMicrosoftEdgeAppChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsMicrosoftEdgeAppChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsMicrosoftEdgeAppChannel(input string) (*WindowsMicrosoftEdgeAppChannel, error) {
	vals := map[string]WindowsMicrosoftEdgeAppChannel{
		"beta":   WindowsMicrosoftEdgeAppChannel_Beta,
		"dev":    WindowsMicrosoftEdgeAppChannel_Dev,
		"stable": WindowsMicrosoftEdgeAppChannel_Stable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsMicrosoftEdgeAppChannel(input)
	return &out, nil
}
