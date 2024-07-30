package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessTemplateScenarios string

const (
	ConditionalAccessTemplateScenarios_EmergingThreats  ConditionalAccessTemplateScenarios = "emergingThreats"
	ConditionalAccessTemplateScenarios_New              ConditionalAccessTemplateScenarios = "new"
	ConditionalAccessTemplateScenarios_ProtectAdmins    ConditionalAccessTemplateScenarios = "protectAdmins"
	ConditionalAccessTemplateScenarios_RemoteWork       ConditionalAccessTemplateScenarios = "remoteWork"
	ConditionalAccessTemplateScenarios_SecureFoundation ConditionalAccessTemplateScenarios = "secureFoundation"
	ConditionalAccessTemplateScenarios_ZeroTrust        ConditionalAccessTemplateScenarios = "zeroTrust"
)

func PossibleValuesForConditionalAccessTemplateScenarios() []string {
	return []string{
		string(ConditionalAccessTemplateScenarios_EmergingThreats),
		string(ConditionalAccessTemplateScenarios_New),
		string(ConditionalAccessTemplateScenarios_ProtectAdmins),
		string(ConditionalAccessTemplateScenarios_RemoteWork),
		string(ConditionalAccessTemplateScenarios_SecureFoundation),
		string(ConditionalAccessTemplateScenarios_ZeroTrust),
	}
}

func (s *ConditionalAccessTemplateScenarios) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessTemplateScenarios(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessTemplateScenarios(input string) (*ConditionalAccessTemplateScenarios, error) {
	vals := map[string]ConditionalAccessTemplateScenarios{
		"emergingthreats":  ConditionalAccessTemplateScenarios_EmergingThreats,
		"new":              ConditionalAccessTemplateScenarios_New,
		"protectadmins":    ConditionalAccessTemplateScenarios_ProtectAdmins,
		"remotework":       ConditionalAccessTemplateScenarios_RemoteWork,
		"securefoundation": ConditionalAccessTemplateScenarios_SecureFoundation,
		"zerotrust":        ConditionalAccessTemplateScenarios_ZeroTrust,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessTemplateScenarios(input)
	return &out, nil
}
