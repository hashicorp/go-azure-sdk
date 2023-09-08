package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordefault       MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "Default"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordownloadOnly  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "DownloadOnly"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallASAP   MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "InstallASAP"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallLater  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "InstallLater"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotConfigured MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "NotConfigured"
	MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotifyOnly    MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior = "NotifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordefault),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordownloadOnly),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallASAP),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallLater),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotConfigured),
		string(MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotifyOnly),
	}
}

func parseMacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordefault,
		"downloadonly":  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviordownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviorinstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationFirmwareUpdateBehaviornotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior(input)
	return &out, nil
}
