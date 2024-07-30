package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_Default       MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "default"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_DownloadOnly  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "downloadOnly"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallASAP   MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "installASAP"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallLater  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "installLater"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotConfigured MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "notConfigured"
	MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotifyOnly    MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior = "notifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationCriticalUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_Default),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_DownloadOnly),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallASAP),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallLater),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotConfigured),
		string(MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotifyOnly),
	}
}

func (s *MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationCriticalUpdateBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationCriticalUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_Default,
		"downloadonly":  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_DownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_InstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior_NotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior(input)
	return &out, nil
}
