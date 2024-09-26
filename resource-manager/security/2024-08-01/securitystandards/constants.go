package securitystandards

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardSupportedCloud string

const (
	StandardSupportedCloudAWS   StandardSupportedCloud = "AWS"
	StandardSupportedCloudAzure StandardSupportedCloud = "Azure"
	StandardSupportedCloudGCP   StandardSupportedCloud = "GCP"
)

func PossibleValuesForStandardSupportedCloud() []string {
	return []string{
		string(StandardSupportedCloudAWS),
		string(StandardSupportedCloudAzure),
		string(StandardSupportedCloudGCP),
	}
}

func (s *StandardSupportedCloud) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStandardSupportedCloud(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStandardSupportedCloud(input string) (*StandardSupportedCloud, error) {
	vals := map[string]StandardSupportedCloud{
		"aws":   StandardSupportedCloudAWS,
		"azure": StandardSupportedCloudAzure,
		"gcp":   StandardSupportedCloudGCP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StandardSupportedCloud(input)
	return &out, nil
}

type StandardType string

const (
	StandardTypeCompliance StandardType = "Compliance"
	StandardTypeCustom     StandardType = "Custom"
	StandardTypeDefault    StandardType = "Default"
)

func PossibleValuesForStandardType() []string {
	return []string{
		string(StandardTypeCompliance),
		string(StandardTypeCustom),
		string(StandardTypeDefault),
	}
}

func (s *StandardType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStandardType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStandardType(input string) (*StandardType, error) {
	vals := map[string]StandardType{
		"compliance": StandardTypeCompliance,
		"custom":     StandardTypeCustom,
		"default":    StandardTypeDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StandardType(input)
	return &out, nil
}
