package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCustomConfigurationDeploymentChannel string

const (
	MacOSCustomConfigurationDeploymentChannel_DeviceChannel MacOSCustomConfigurationDeploymentChannel = "deviceChannel"
	MacOSCustomConfigurationDeploymentChannel_UserChannel   MacOSCustomConfigurationDeploymentChannel = "userChannel"
)

func PossibleValuesForMacOSCustomConfigurationDeploymentChannel() []string {
	return []string{
		string(MacOSCustomConfigurationDeploymentChannel_DeviceChannel),
		string(MacOSCustomConfigurationDeploymentChannel_UserChannel),
	}
}

func (s *MacOSCustomConfigurationDeploymentChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCustomConfigurationDeploymentChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCustomConfigurationDeploymentChannel(input string) (*MacOSCustomConfigurationDeploymentChannel, error) {
	vals := map[string]MacOSCustomConfigurationDeploymentChannel{
		"devicechannel": MacOSCustomConfigurationDeploymentChannel_DeviceChannel,
		"userchannel":   MacOSCustomConfigurationDeploymentChannel_UserChannel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCustomConfigurationDeploymentChannel(input)
	return &out, nil
}
