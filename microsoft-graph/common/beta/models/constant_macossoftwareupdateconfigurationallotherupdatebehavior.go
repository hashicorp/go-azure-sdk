package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordefault       MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "Default"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordownloadOnly  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "DownloadOnly"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallASAP   MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "InstallASAP"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallLater  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "InstallLater"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotConfigured MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "NotConfigured"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotifyOnly    MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "NotifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordefault),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordownloadOnly),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallASAP),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallLater),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotConfigured),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotifyOnly),
	}
}

func parseMacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordefault,
		"downloadonly":  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviordownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviorinstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationAllOtherUpdateBehaviornotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior(input)
	return &out, nil
}
