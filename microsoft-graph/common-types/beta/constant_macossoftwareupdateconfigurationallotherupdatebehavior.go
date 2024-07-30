package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior string

const (
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_Default       MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "default"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_DownloadOnly  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "downloadOnly"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallASAP   MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "installASAP"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallLater  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "installLater"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotConfigured MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "notConfigured"
	MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotifyOnly    MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior = "notifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_Default),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_DownloadOnly),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallASAP),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallLater),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotConfigured),
		string(MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotifyOnly),
	}
}

func (s *MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior(input string) (*MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior{
		"default":       MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_Default,
		"downloadonly":  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_DownloadOnly,
		"installasap":   MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallASAP,
		"installlater":  MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_InstallLater,
		"notconfigured": MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotConfigured,
		"notifyonly":    MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior_NotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior(input)
	return &out, nil
}
