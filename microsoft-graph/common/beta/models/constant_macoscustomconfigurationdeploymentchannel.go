package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCustomConfigurationDeploymentChannel string

const (
	MacOSCustomConfigurationDeploymentChanneldeviceChannel MacOSCustomConfigurationDeploymentChannel = "DeviceChannel"
	MacOSCustomConfigurationDeploymentChanneluserChannel   MacOSCustomConfigurationDeploymentChannel = "UserChannel"
)

func PossibleValuesForMacOSCustomConfigurationDeploymentChannel() []string {
	return []string{
		string(MacOSCustomConfigurationDeploymentChanneldeviceChannel),
		string(MacOSCustomConfigurationDeploymentChanneluserChannel),
	}
}

func parseMacOSCustomConfigurationDeploymentChannel(input string) (*MacOSCustomConfigurationDeploymentChannel, error) {
	vals := map[string]MacOSCustomConfigurationDeploymentChannel{
		"devicechannel": MacOSCustomConfigurationDeploymentChanneldeviceChannel,
		"userchannel":   MacOSCustomConfigurationDeploymentChanneluserChannel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCustomConfigurationDeploymentChannel(input)
	return &out, nil
}
