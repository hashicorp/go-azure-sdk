package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateStateSummaryState string

const (
	MacOSSoftwareUpdateStateSummaryState_Available                   MacOSSoftwareUpdateStateSummaryState = "available"
	MacOSSoftwareUpdateStateSummaryState_CommandFailed               MacOSSoftwareUpdateStateSummaryState = "commandFailed"
	MacOSSoftwareUpdateStateSummaryState_DownloadFailed              MacOSSoftwareUpdateStateSummaryState = "downloadFailed"
	MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientNetwork MacOSSoftwareUpdateStateSummaryState = "downloadInsufficientNetwork"
	MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientPower   MacOSSoftwareUpdateStateSummaryState = "downloadInsufficientPower"
	MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientSpace   MacOSSoftwareUpdateStateSummaryState = "downloadInsufficientSpace"
	MacOSSoftwareUpdateStateSummaryState_Downloaded                  MacOSSoftwareUpdateStateSummaryState = "downloaded"
	MacOSSoftwareUpdateStateSummaryState_Downloading                 MacOSSoftwareUpdateStateSummaryState = "downloading"
	MacOSSoftwareUpdateStateSummaryState_Idle                        MacOSSoftwareUpdateStateSummaryState = "idle"
	MacOSSoftwareUpdateStateSummaryState_InstallFailed               MacOSSoftwareUpdateStateSummaryState = "installFailed"
	MacOSSoftwareUpdateStateSummaryState_InstallInsufficientPower    MacOSSoftwareUpdateStateSummaryState = "installInsufficientPower"
	MacOSSoftwareUpdateStateSummaryState_InstallInsufficientSpace    MacOSSoftwareUpdateStateSummaryState = "installInsufficientSpace"
	MacOSSoftwareUpdateStateSummaryState_Installing                  MacOSSoftwareUpdateStateSummaryState = "installing"
	MacOSSoftwareUpdateStateSummaryState_Scheduled                   MacOSSoftwareUpdateStateSummaryState = "scheduled"
	MacOSSoftwareUpdateStateSummaryState_Success                     MacOSSoftwareUpdateStateSummaryState = "success"
)

func PossibleValuesForMacOSSoftwareUpdateStateSummaryState() []string {
	return []string{
		string(MacOSSoftwareUpdateStateSummaryState_Available),
		string(MacOSSoftwareUpdateStateSummaryState_CommandFailed),
		string(MacOSSoftwareUpdateStateSummaryState_DownloadFailed),
		string(MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientNetwork),
		string(MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientPower),
		string(MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientSpace),
		string(MacOSSoftwareUpdateStateSummaryState_Downloaded),
		string(MacOSSoftwareUpdateStateSummaryState_Downloading),
		string(MacOSSoftwareUpdateStateSummaryState_Idle),
		string(MacOSSoftwareUpdateStateSummaryState_InstallFailed),
		string(MacOSSoftwareUpdateStateSummaryState_InstallInsufficientPower),
		string(MacOSSoftwareUpdateStateSummaryState_InstallInsufficientSpace),
		string(MacOSSoftwareUpdateStateSummaryState_Installing),
		string(MacOSSoftwareUpdateStateSummaryState_Scheduled),
		string(MacOSSoftwareUpdateStateSummaryState_Success),
	}
}

func (s *MacOSSoftwareUpdateStateSummaryState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateStateSummaryState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateStateSummaryState(input string) (*MacOSSoftwareUpdateStateSummaryState, error) {
	vals := map[string]MacOSSoftwareUpdateStateSummaryState{
		"available":                   MacOSSoftwareUpdateStateSummaryState_Available,
		"commandfailed":               MacOSSoftwareUpdateStateSummaryState_CommandFailed,
		"downloadfailed":              MacOSSoftwareUpdateStateSummaryState_DownloadFailed,
		"downloadinsufficientnetwork": MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientNetwork,
		"downloadinsufficientpower":   MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientPower,
		"downloadinsufficientspace":   MacOSSoftwareUpdateStateSummaryState_DownloadInsufficientSpace,
		"downloaded":                  MacOSSoftwareUpdateStateSummaryState_Downloaded,
		"downloading":                 MacOSSoftwareUpdateStateSummaryState_Downloading,
		"idle":                        MacOSSoftwareUpdateStateSummaryState_Idle,
		"installfailed":               MacOSSoftwareUpdateStateSummaryState_InstallFailed,
		"installinsufficientpower":    MacOSSoftwareUpdateStateSummaryState_InstallInsufficientPower,
		"installinsufficientspace":    MacOSSoftwareUpdateStateSummaryState_InstallInsufficientSpace,
		"installing":                  MacOSSoftwareUpdateStateSummaryState_Installing,
		"scheduled":                   MacOSSoftwareUpdateStateSummaryState_Scheduled,
		"success":                     MacOSSoftwareUpdateStateSummaryState_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateStateSummaryState(input)
	return &out, nil
}
