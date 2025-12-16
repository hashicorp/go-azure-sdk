package diagnosticsettings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProactiveDiagnosticsConsent string

const (
	ProactiveDiagnosticsConsentDisabled ProactiveDiagnosticsConsent = "Disabled"
	ProactiveDiagnosticsConsentEnabled  ProactiveDiagnosticsConsent = "Enabled"
)

func PossibleValuesForProactiveDiagnosticsConsent() []string {
	return []string{
		string(ProactiveDiagnosticsConsentDisabled),
		string(ProactiveDiagnosticsConsentEnabled),
	}
}

func (s *ProactiveDiagnosticsConsent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProactiveDiagnosticsConsent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProactiveDiagnosticsConsent(input string) (*ProactiveDiagnosticsConsent, error) {
	vals := map[string]ProactiveDiagnosticsConsent{
		"disabled": ProactiveDiagnosticsConsentDisabled,
		"enabled":  ProactiveDiagnosticsConsentEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProactiveDiagnosticsConsent(input)
	return &out, nil
}
