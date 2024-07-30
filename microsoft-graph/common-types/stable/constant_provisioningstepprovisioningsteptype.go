package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStepProvisioningStepType string

const (
	ProvisioningStepProvisioningStepType_Export              ProvisioningStepProvisioningStepType = "export"
	ProvisioningStepProvisioningStepType_Import              ProvisioningStepProvisioningStepType = "import"
	ProvisioningStepProvisioningStepType_Matching            ProvisioningStepProvisioningStepType = "matching"
	ProvisioningStepProvisioningStepType_Processing          ProvisioningStepProvisioningStepType = "processing"
	ProvisioningStepProvisioningStepType_ReferenceResolution ProvisioningStepProvisioningStepType = "referenceResolution"
	ProvisioningStepProvisioningStepType_Scoping             ProvisioningStepProvisioningStepType = "scoping"
)

func PossibleValuesForProvisioningStepProvisioningStepType() []string {
	return []string{
		string(ProvisioningStepProvisioningStepType_Export),
		string(ProvisioningStepProvisioningStepType_Import),
		string(ProvisioningStepProvisioningStepType_Matching),
		string(ProvisioningStepProvisioningStepType_Processing),
		string(ProvisioningStepProvisioningStepType_ReferenceResolution),
		string(ProvisioningStepProvisioningStepType_Scoping),
	}
}

func (s *ProvisioningStepProvisioningStepType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStepProvisioningStepType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStepProvisioningStepType(input string) (*ProvisioningStepProvisioningStepType, error) {
	vals := map[string]ProvisioningStepProvisioningStepType{
		"export":              ProvisioningStepProvisioningStepType_Export,
		"import":              ProvisioningStepProvisioningStepType_Import,
		"matching":            ProvisioningStepProvisioningStepType_Matching,
		"processing":          ProvisioningStepProvisioningStepType_Processing,
		"referenceresolution": ProvisioningStepProvisioningStepType_ReferenceResolution,
		"scoping":             ProvisioningStepProvisioningStepType_Scoping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStepProvisioningStepType(input)
	return &out, nil
}
