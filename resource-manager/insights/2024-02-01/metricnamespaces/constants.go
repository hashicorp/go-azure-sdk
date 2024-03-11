package metricnamespaces

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceClassification string

const (
	NamespaceClassificationCustom   NamespaceClassification = "Custom"
	NamespaceClassificationPlatform NamespaceClassification = "Platform"
	NamespaceClassificationQos      NamespaceClassification = "Qos"
)

func PossibleValuesForNamespaceClassification() []string {
	return []string{
		string(NamespaceClassificationCustom),
		string(NamespaceClassificationPlatform),
		string(NamespaceClassificationQos),
	}
}

func (s *NamespaceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNamespaceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNamespaceClassification(input string) (*NamespaceClassification, error) {
	vals := map[string]NamespaceClassification{
		"custom":   NamespaceClassificationCustom,
		"platform": NamespaceClassificationPlatform,
		"qos":      NamespaceClassificationQos,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NamespaceClassification(input)
	return &out, nil
}
