package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationNetworkInterface string

const (
	MacOSWiredNetworkConfigurationNetworkInterface_AnyEthernet          MacOSWiredNetworkConfigurationNetworkInterface = "anyEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_FirstActiveEthernet  MacOSWiredNetworkConfigurationNetworkInterface = "firstActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_FirstEthernet        MacOSWiredNetworkConfigurationNetworkInterface = "firstEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_SecondActiveEthernet MacOSWiredNetworkConfigurationNetworkInterface = "secondActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_SecondEthernet       MacOSWiredNetworkConfigurationNetworkInterface = "secondEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_ThirdActiveEthernet  MacOSWiredNetworkConfigurationNetworkInterface = "thirdActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterface_ThirdEthernet        MacOSWiredNetworkConfigurationNetworkInterface = "thirdEthernet"
)

func PossibleValuesForMacOSWiredNetworkConfigurationNetworkInterface() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationNetworkInterface_AnyEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_FirstActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_FirstEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_SecondActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_SecondEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_ThirdActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterface_ThirdEthernet),
	}
}

func (s *MacOSWiredNetworkConfigurationNetworkInterface) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiredNetworkConfigurationNetworkInterface(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiredNetworkConfigurationNetworkInterface(input string) (*MacOSWiredNetworkConfigurationNetworkInterface, error) {
	vals := map[string]MacOSWiredNetworkConfigurationNetworkInterface{
		"anyethernet":          MacOSWiredNetworkConfigurationNetworkInterface_AnyEthernet,
		"firstactiveethernet":  MacOSWiredNetworkConfigurationNetworkInterface_FirstActiveEthernet,
		"firstethernet":        MacOSWiredNetworkConfigurationNetworkInterface_FirstEthernet,
		"secondactiveethernet": MacOSWiredNetworkConfigurationNetworkInterface_SecondActiveEthernet,
		"secondethernet":       MacOSWiredNetworkConfigurationNetworkInterface_SecondEthernet,
		"thirdactiveethernet":  MacOSWiredNetworkConfigurationNetworkInterface_ThirdActiveEthernet,
		"thirdethernet":        MacOSWiredNetworkConfigurationNetworkInterface_ThirdEthernet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationNetworkInterface(input)
	return &out, nil
}
