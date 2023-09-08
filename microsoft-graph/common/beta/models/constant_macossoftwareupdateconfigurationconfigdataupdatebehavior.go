package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordefault       MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "Default"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordownloadOnly  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "DownloadOnly"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallASAP   MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "InstallASAP"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallLater  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "InstallLater"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotConfigured MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "NotConfigured"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotifyOnly    MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "NotifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordefault),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordownloadOnly),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallASAP),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallLater),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotConfigured),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotifyOnly),
	}
}

func parseMacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordefault,
		"downloadonly":  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviordownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviorinstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationConfigDataUpdateBehaviornotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior(input)
	return &out, nil
}
