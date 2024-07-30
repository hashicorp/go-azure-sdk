package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_Default       MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "default"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_DownloadOnly  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "downloadOnly"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallASAP   MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "installASAP"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallLater  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "installLater"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotConfigured MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "notConfigured"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotifyOnly    MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "notifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_Default),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_DownloadOnly),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallASAP),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallLater),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotConfigured),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotifyOnly),
	}
}

func (s *MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_Default,
		"downloadonly":  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_DownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_InstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior_NotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior(input)
	return &out, nil
}
