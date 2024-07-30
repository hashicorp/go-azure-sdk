package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppProductIds string

const (
	OfficeSuiteAppProductIds_O365BusinessRetail OfficeSuiteAppProductIds = "o365BusinessRetail"
	OfficeSuiteAppProductIds_O365ProPlusRetail  OfficeSuiteAppProductIds = "o365ProPlusRetail"
	OfficeSuiteAppProductIds_ProjectProRetail   OfficeSuiteAppProductIds = "projectProRetail"
	OfficeSuiteAppProductIds_VisioProRetail     OfficeSuiteAppProductIds = "visioProRetail"
)

func PossibleValuesForOfficeSuiteAppProductIds() []string {
	return []string{
		string(OfficeSuiteAppProductIds_O365BusinessRetail),
		string(OfficeSuiteAppProductIds_O365ProPlusRetail),
		string(OfficeSuiteAppProductIds_ProjectProRetail),
		string(OfficeSuiteAppProductIds_VisioProRetail),
	}
}

func (s *OfficeSuiteAppProductIds) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppProductIds(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppProductIds(input string) (*OfficeSuiteAppProductIds, error) {
	vals := map[string]OfficeSuiteAppProductIds{
		"o365businessretail": OfficeSuiteAppProductIds_O365BusinessRetail,
		"o365proplusretail":  OfficeSuiteAppProductIds_O365ProPlusRetail,
		"projectproretail":   OfficeSuiteAppProductIds_ProjectProRetail,
		"visioproretail":     OfficeSuiteAppProductIds_VisioProRetail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppProductIds(input)
	return &out, nil
}
