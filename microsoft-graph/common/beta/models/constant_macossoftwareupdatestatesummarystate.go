package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateStateSummaryState string

const (
	MacOSSoftwareUpdateStateSummaryStateavailable                   MacOSSoftwareUpdateStateSummaryState = "Available"
	MacOSSoftwareUpdateStateSummaryStatecommandFailed               MacOSSoftwareUpdateStateSummaryState = "CommandFailed"
	MacOSSoftwareUpdateStateSummaryStatedownloadFailed              MacOSSoftwareUpdateStateSummaryState = "DownloadFailed"
	MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientNetwork MacOSSoftwareUpdateStateSummaryState = "DownloadInsufficientNetwork"
	MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientPower   MacOSSoftwareUpdateStateSummaryState = "DownloadInsufficientPower"
	MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientSpace   MacOSSoftwareUpdateStateSummaryState = "DownloadInsufficientSpace"
	MacOSSoftwareUpdateStateSummaryStatedownloaded                  MacOSSoftwareUpdateStateSummaryState = "Downloaded"
	MacOSSoftwareUpdateStateSummaryStatedownloading                 MacOSSoftwareUpdateStateSummaryState = "Downloading"
	MacOSSoftwareUpdateStateSummaryStateidle                        MacOSSoftwareUpdateStateSummaryState = "Idle"
	MacOSSoftwareUpdateStateSummaryStateinstallFailed               MacOSSoftwareUpdateStateSummaryState = "InstallFailed"
	MacOSSoftwareUpdateStateSummaryStateinstallInsufficientPower    MacOSSoftwareUpdateStateSummaryState = "InstallInsufficientPower"
	MacOSSoftwareUpdateStateSummaryStateinstallInsufficientSpace    MacOSSoftwareUpdateStateSummaryState = "InstallInsufficientSpace"
	MacOSSoftwareUpdateStateSummaryStateinstalling                  MacOSSoftwareUpdateStateSummaryState = "Installing"
	MacOSSoftwareUpdateStateSummaryStatescheduled                   MacOSSoftwareUpdateStateSummaryState = "Scheduled"
	MacOSSoftwareUpdateStateSummaryStatesuccess                     MacOSSoftwareUpdateStateSummaryState = "Success"
)

func PossibleValuesForMacOSSoftwareUpdateStateSummaryState() []string {
	return []string{
		string(MacOSSoftwareUpdateStateSummaryStateavailable),
		string(MacOSSoftwareUpdateStateSummaryStatecommandFailed),
		string(MacOSSoftwareUpdateStateSummaryStatedownloadFailed),
		string(MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientNetwork),
		string(MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientPower),
		string(MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientSpace),
		string(MacOSSoftwareUpdateStateSummaryStatedownloaded),
		string(MacOSSoftwareUpdateStateSummaryStatedownloading),
		string(MacOSSoftwareUpdateStateSummaryStateidle),
		string(MacOSSoftwareUpdateStateSummaryStateinstallFailed),
		string(MacOSSoftwareUpdateStateSummaryStateinstallInsufficientPower),
		string(MacOSSoftwareUpdateStateSummaryStateinstallInsufficientSpace),
		string(MacOSSoftwareUpdateStateSummaryStateinstalling),
		string(MacOSSoftwareUpdateStateSummaryStatescheduled),
		string(MacOSSoftwareUpdateStateSummaryStatesuccess),
	}
}

func parseMacOSSoftwareUpdateStateSummaryState(input string) (*MacOSSoftwareUpdateStateSummaryState, error) {
	vals := map[string]MacOSSoftwareUpdateStateSummaryState{
		"available":                   MacOSSoftwareUpdateStateSummaryStateavailable,
		"commandfailed":               MacOSSoftwareUpdateStateSummaryStatecommandFailed,
		"downloadfailed":              MacOSSoftwareUpdateStateSummaryStatedownloadFailed,
		"downloadinsufficientnetwork": MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientNetwork,
		"downloadinsufficientpower":   MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientPower,
		"downloadinsufficientspace":   MacOSSoftwareUpdateStateSummaryStatedownloadInsufficientSpace,
		"downloaded":                  MacOSSoftwareUpdateStateSummaryStatedownloaded,
		"downloading":                 MacOSSoftwareUpdateStateSummaryStatedownloading,
		"idle":                        MacOSSoftwareUpdateStateSummaryStateidle,
		"installfailed":               MacOSSoftwareUpdateStateSummaryStateinstallFailed,
		"installinsufficientpower":    MacOSSoftwareUpdateStateSummaryStateinstallInsufficientPower,
		"installinsufficientspace":    MacOSSoftwareUpdateStateSummaryStateinstallInsufficientSpace,
		"installing":                  MacOSSoftwareUpdateStateSummaryStateinstalling,
		"scheduled":                   MacOSSoftwareUpdateStateSummaryStatescheduled,
		"success":                     MacOSSoftwareUpdateStateSummaryStatesuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateStateSummaryState(input)
	return &out, nil
}
