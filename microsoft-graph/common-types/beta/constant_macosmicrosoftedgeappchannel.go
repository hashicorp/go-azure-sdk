package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMicrosoftEdgeAppChannel string

const (
	MacOSMicrosoftEdgeAppChannel_Beta   MacOSMicrosoftEdgeAppChannel = "beta"
	MacOSMicrosoftEdgeAppChannel_Dev    MacOSMicrosoftEdgeAppChannel = "dev"
	MacOSMicrosoftEdgeAppChannel_Stable MacOSMicrosoftEdgeAppChannel = "stable"
)

func PossibleValuesForMacOSMicrosoftEdgeAppChannel() []string {
	return []string{
		string(MacOSMicrosoftEdgeAppChannel_Beta),
		string(MacOSMicrosoftEdgeAppChannel_Dev),
		string(MacOSMicrosoftEdgeAppChannel_Stable),
	}
}

func (s *MacOSMicrosoftEdgeAppChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSMicrosoftEdgeAppChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSMicrosoftEdgeAppChannel(input string) (*MacOSMicrosoftEdgeAppChannel, error) {
	vals := map[string]MacOSMicrosoftEdgeAppChannel{
		"beta":   MacOSMicrosoftEdgeAppChannel_Beta,
		"dev":    MacOSMicrosoftEdgeAppChannel_Dev,
		"stable": MacOSMicrosoftEdgeAppChannel_Stable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSMicrosoftEdgeAppChannel(input)
	return &out, nil
}
