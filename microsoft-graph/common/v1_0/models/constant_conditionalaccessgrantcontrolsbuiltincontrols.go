package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGrantControlsBuiltInControls string

const (
	ConditionalAccessGrantControlsBuiltInControlsapprovedApplication  ConditionalAccessGrantControlsBuiltInControls = "ApprovedApplication"
	ConditionalAccessGrantControlsBuiltInControlsblock                ConditionalAccessGrantControlsBuiltInControls = "Block"
	ConditionalAccessGrantControlsBuiltInControlscompliantApplication ConditionalAccessGrantControlsBuiltInControls = "CompliantApplication"
	ConditionalAccessGrantControlsBuiltInControlscompliantDevice      ConditionalAccessGrantControlsBuiltInControls = "CompliantDevice"
	ConditionalAccessGrantControlsBuiltInControlsdomainJoinedDevice   ConditionalAccessGrantControlsBuiltInControls = "DomainJoinedDevice"
	ConditionalAccessGrantControlsBuiltInControlsmfa                  ConditionalAccessGrantControlsBuiltInControls = "Mfa"
	ConditionalAccessGrantControlsBuiltInControlspasswordChange       ConditionalAccessGrantControlsBuiltInControls = "PasswordChange"
)

func PossibleValuesForConditionalAccessGrantControlsBuiltInControls() []string {
	return []string{
		string(ConditionalAccessGrantControlsBuiltInControlsapprovedApplication),
		string(ConditionalAccessGrantControlsBuiltInControlsblock),
		string(ConditionalAccessGrantControlsBuiltInControlscompliantApplication),
		string(ConditionalAccessGrantControlsBuiltInControlscompliantDevice),
		string(ConditionalAccessGrantControlsBuiltInControlsdomainJoinedDevice),
		string(ConditionalAccessGrantControlsBuiltInControlsmfa),
		string(ConditionalAccessGrantControlsBuiltInControlspasswordChange),
	}
}

func parseConditionalAccessGrantControlsBuiltInControls(input string) (*ConditionalAccessGrantControlsBuiltInControls, error) {
	vals := map[string]ConditionalAccessGrantControlsBuiltInControls{
		"approvedapplication":  ConditionalAccessGrantControlsBuiltInControlsapprovedApplication,
		"block":                ConditionalAccessGrantControlsBuiltInControlsblock,
		"compliantapplication": ConditionalAccessGrantControlsBuiltInControlscompliantApplication,
		"compliantdevice":      ConditionalAccessGrantControlsBuiltInControlscompliantDevice,
		"domainjoineddevice":   ConditionalAccessGrantControlsBuiltInControlsdomainJoinedDevice,
		"mfa":                  ConditionalAccessGrantControlsBuiltInControlsmfa,
		"passwordchange":       ConditionalAccessGrantControlsBuiltInControlspasswordChange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGrantControlsBuiltInControls(input)
	return &out, nil
}
