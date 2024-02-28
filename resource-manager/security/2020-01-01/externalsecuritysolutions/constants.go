package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadConnectivityState string

const (
	AadConnectivityStateConnected   AadConnectivityState = "Connected"
	AadConnectivityStateDiscovered  AadConnectivityState = "Discovered"
	AadConnectivityStateNotLicensed AadConnectivityState = "NotLicensed"
)

func PossibleValuesForAadConnectivityState() []string {
	return []string{
		string(AadConnectivityStateConnected),
		string(AadConnectivityStateDiscovered),
		string(AadConnectivityStateNotLicensed),
	}
}

func (s *AadConnectivityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAadConnectivityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAadConnectivityState(input string) (*AadConnectivityState, error) {
	vals := map[string]AadConnectivityState{
		"connected":   AadConnectivityStateConnected,
		"discovered":  AadConnectivityStateDiscovered,
		"notlicensed": AadConnectivityStateNotLicensed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AadConnectivityState(input)
	return &out, nil
}

type ExternalSecuritySolutionKind string

const (
	ExternalSecuritySolutionKindAAD ExternalSecuritySolutionKind = "AAD"
	ExternalSecuritySolutionKindATA ExternalSecuritySolutionKind = "ATA"
	ExternalSecuritySolutionKindCEF ExternalSecuritySolutionKind = "CEF"
)

func PossibleValuesForExternalSecuritySolutionKind() []string {
	return []string{
		string(ExternalSecuritySolutionKindAAD),
		string(ExternalSecuritySolutionKindATA),
		string(ExternalSecuritySolutionKindCEF),
	}
}

func (s *ExternalSecuritySolutionKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalSecuritySolutionKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalSecuritySolutionKind(input string) (*ExternalSecuritySolutionKind, error) {
	vals := map[string]ExternalSecuritySolutionKind{
		"aad": ExternalSecuritySolutionKindAAD,
		"ata": ExternalSecuritySolutionKindATA,
		"cef": ExternalSecuritySolutionKindCEF,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalSecuritySolutionKind(input)
	return &out, nil
}
