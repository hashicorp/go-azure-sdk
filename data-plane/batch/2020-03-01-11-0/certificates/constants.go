package certificates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateFormat string

const (
	CertificateFormatCer CertificateFormat = "cer"
	CertificateFormatPfx CertificateFormat = "pfx"
)

func PossibleValuesForCertificateFormat() []string {
	return []string{
		string(CertificateFormatCer),
		string(CertificateFormatPfx),
	}
}

func (s *CertificateFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateFormat(input string) (*CertificateFormat, error) {
	vals := map[string]CertificateFormat{
		"cer": CertificateFormatCer,
		"pfx": CertificateFormatPfx,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateFormat(input)
	return &out, nil
}

type CertificateState string

const (
	CertificateStateActive       CertificateState = "active"
	CertificateStateDeletefailed CertificateState = "deletefailed"
	CertificateStateDeleting     CertificateState = "deleting"
)

func PossibleValuesForCertificateState() []string {
	return []string{
		string(CertificateStateActive),
		string(CertificateStateDeletefailed),
		string(CertificateStateDeleting),
	}
}

func (s *CertificateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateState(input string) (*CertificateState, error) {
	vals := map[string]CertificateState{
		"active":       CertificateStateActive,
		"deletefailed": CertificateStateDeletefailed,
		"deleting":     CertificateStateDeleting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateState(input)
	return &out, nil
}
