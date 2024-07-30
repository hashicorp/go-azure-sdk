package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnectionEnabledContentExperiences string

const (
	ExternalConnectorsExternalConnectionEnabledContentExperiences_Compliance ExternalConnectorsExternalConnectionEnabledContentExperiences = "compliance"
	ExternalConnectorsExternalConnectionEnabledContentExperiences_Search     ExternalConnectorsExternalConnectionEnabledContentExperiences = "search"
)

func PossibleValuesForExternalConnectorsExternalConnectionEnabledContentExperiences() []string {
	return []string{
		string(ExternalConnectorsExternalConnectionEnabledContentExperiences_Compliance),
		string(ExternalConnectorsExternalConnectionEnabledContentExperiences_Search),
	}
}

func (s *ExternalConnectorsExternalConnectionEnabledContentExperiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsExternalConnectionEnabledContentExperiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsExternalConnectionEnabledContentExperiences(input string) (*ExternalConnectorsExternalConnectionEnabledContentExperiences, error) {
	vals := map[string]ExternalConnectorsExternalConnectionEnabledContentExperiences{
		"compliance": ExternalConnectorsExternalConnectionEnabledContentExperiences_Compliance,
		"search":     ExternalConnectorsExternalConnectionEnabledContentExperiences_Search,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalConnectionEnabledContentExperiences(input)
	return &out, nil
}
