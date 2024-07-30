package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInOriginalTransferMethod string

const (
	SignInOriginalTransferMethod_AuthenticationTransfer SignInOriginalTransferMethod = "authenticationTransfer"
	SignInOriginalTransferMethod_DeviceCodeFlow         SignInOriginalTransferMethod = "deviceCodeFlow"
	SignInOriginalTransferMethod_None                   SignInOriginalTransferMethod = "none"
)

func PossibleValuesForSignInOriginalTransferMethod() []string {
	return []string{
		string(SignInOriginalTransferMethod_AuthenticationTransfer),
		string(SignInOriginalTransferMethod_DeviceCodeFlow),
		string(SignInOriginalTransferMethod_None),
	}
}

func (s *SignInOriginalTransferMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInOriginalTransferMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInOriginalTransferMethod(input string) (*SignInOriginalTransferMethod, error) {
	vals := map[string]SignInOriginalTransferMethod{
		"authenticationtransfer": SignInOriginalTransferMethod_AuthenticationTransfer,
		"devicecodeflow":         SignInOriginalTransferMethod_DeviceCodeFlow,
		"none":                   SignInOriginalTransferMethod_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInOriginalTransferMethod(input)
	return &out, nil
}
