package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInOriginalTransferMethod string

const (
	SignInOriginalTransferMethodauthenticationTransfer SignInOriginalTransferMethod = "AuthenticationTransfer"
	SignInOriginalTransferMethoddeviceCodeFlow         SignInOriginalTransferMethod = "DeviceCodeFlow"
	SignInOriginalTransferMethodnone                   SignInOriginalTransferMethod = "None"
)

func PossibleValuesForSignInOriginalTransferMethod() []string {
	return []string{
		string(SignInOriginalTransferMethodauthenticationTransfer),
		string(SignInOriginalTransferMethoddeviceCodeFlow),
		string(SignInOriginalTransferMethodnone),
	}
}

func parseSignInOriginalTransferMethod(input string) (*SignInOriginalTransferMethod, error) {
	vals := map[string]SignInOriginalTransferMethod{
		"authenticationtransfer": SignInOriginalTransferMethodauthenticationTransfer,
		"devicecodeflow":         SignInOriginalTransferMethoddeviceCodeFlow,
		"none":                   SignInOriginalTransferMethodnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInOriginalTransferMethod(input)
	return &out, nil
}
