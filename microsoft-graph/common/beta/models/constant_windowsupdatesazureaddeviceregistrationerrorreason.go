package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesAzureADDeviceRegistrationErrorReason string

const (
	WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADDeviceId WindowsUpdatesAzureADDeviceRegistrationErrorReason = "InvalidAzureADDeviceId"
	WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADJoin     WindowsUpdatesAzureADDeviceRegistrationErrorReason = "InvalidAzureADJoin"
	WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidGlobalDeviceId  WindowsUpdatesAzureADDeviceRegistrationErrorReason = "InvalidGlobalDeviceId"
	WindowsUpdatesAzureADDeviceRegistrationErrorReasonmissingTrustType       WindowsUpdatesAzureADDeviceRegistrationErrorReason = "MissingTrustType"
)

func PossibleValuesForWindowsUpdatesAzureADDeviceRegistrationErrorReason() []string {
	return []string{
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADDeviceId),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADJoin),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidGlobalDeviceId),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReasonmissingTrustType),
	}
}

func parseWindowsUpdatesAzureADDeviceRegistrationErrorReason(input string) (*WindowsUpdatesAzureADDeviceRegistrationErrorReason, error) {
	vals := map[string]WindowsUpdatesAzureADDeviceRegistrationErrorReason{
		"invalidazureaddeviceid": WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADDeviceId,
		"invalidazureadjoin":     WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidAzureADJoin,
		"invalidglobaldeviceid":  WindowsUpdatesAzureADDeviceRegistrationErrorReasoninvalidGlobalDeviceId,
		"missingtrusttype":       WindowsUpdatesAzureADDeviceRegistrationErrorReasonmissingTrustType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesAzureADDeviceRegistrationErrorReason(input)
	return &out, nil
}
