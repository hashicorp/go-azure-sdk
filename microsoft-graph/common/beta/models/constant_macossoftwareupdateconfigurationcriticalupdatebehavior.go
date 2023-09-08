package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordefault       MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "Default"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordownloadOnly  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "DownloadOnly"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallASAP   MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "InstallASAP"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallLater  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "InstallLater"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotConfigured MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "NotConfigured"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotifyOnly    MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "NotifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationCriticalUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordefault),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordownloadOnly),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallASAP),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallLater),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotConfigured),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotifyOnly),
	}
}

func parseMacOSSoftwareUpdateConfigurationCriticalUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordefault,
		"downloadonly":  MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviordownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviorinstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationCriticalUpdateBehaviornotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior(input)
	return &out, nil
}
