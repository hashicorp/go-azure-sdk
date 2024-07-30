package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_Default       MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "default"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_DownloadOnly  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "downloadOnly"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallASAP   MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "installASAP"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallLater  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "installLater"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotConfigured MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "notConfigured"
	MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotifyOnly    MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior = "notifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_Default),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_DownloadOnly),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallASAP),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallLater),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotConfigured),
		string(MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotifyOnly),
	}
}

func (s *MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_Default,
		"downloadonly":  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_DownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_InstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior_NotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior(input)
	return &out, nil
}
