package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInCrossTenantAccessType string

const (
	SignInCrossTenantAccessType_B2bCollaboration SignInCrossTenantAccessType = "b2bCollaboration"
	SignInCrossTenantAccessType_B2bDirectConnect SignInCrossTenantAccessType = "b2bDirectConnect"
	SignInCrossTenantAccessType_MicrosoftSupport SignInCrossTenantAccessType = "microsoftSupport"
	SignInCrossTenantAccessType_None             SignInCrossTenantAccessType = "none"
	SignInCrossTenantAccessType_Passthrough      SignInCrossTenantAccessType = "passthrough"
	SignInCrossTenantAccessType_ServiceProvider  SignInCrossTenantAccessType = "serviceProvider"
)

func PossibleValuesForSignInCrossTenantAccessType() []string {
	return []string{
		string(SignInCrossTenantAccessType_B2bCollaboration),
		string(SignInCrossTenantAccessType_B2bDirectConnect),
		string(SignInCrossTenantAccessType_MicrosoftSupport),
		string(SignInCrossTenantAccessType_None),
		string(SignInCrossTenantAccessType_Passthrough),
		string(SignInCrossTenantAccessType_ServiceProvider),
	}
}

func (s *SignInCrossTenantAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInCrossTenantAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInCrossTenantAccessType(input string) (*SignInCrossTenantAccessType, error) {
	vals := map[string]SignInCrossTenantAccessType{
		"b2bcollaboration": SignInCrossTenantAccessType_B2bCollaboration,
		"b2bdirectconnect": SignInCrossTenantAccessType_B2bDirectConnect,
		"microsoftsupport": SignInCrossTenantAccessType_MicrosoftSupport,
		"none":             SignInCrossTenantAccessType_None,
		"passthrough":      SignInCrossTenantAccessType_Passthrough,
		"serviceprovider":  SignInCrossTenantAccessType_ServiceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInCrossTenantAccessType(input)
	return &out, nil
}
