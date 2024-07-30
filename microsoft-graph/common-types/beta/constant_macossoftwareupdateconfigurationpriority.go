package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationPriority string

const (
	MacOSSoftwareUpdateConfigurationPriority_High MacOSSoftwareUpdateConfigurationPriority = "high"
	MacOSSoftwareUpdateConfigurationPriority_Low  MacOSSoftwareUpdateConfigurationPriority = "low"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationPriority() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationPriority_High),
		string(MacOSSoftwareUpdateConfigurationPriority_Low),
	}
}

func (s *MacOSSoftwareUpdateConfigurationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationPriority(input string) (*MacOSSoftwareUpdateConfigurationPriority, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationPriority{
		"high": MacOSSoftwareUpdateConfigurationPriority_High,
		"low":  MacOSSoftwareUpdateConfigurationPriority_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationPriority(input)
	return &out, nil
}
