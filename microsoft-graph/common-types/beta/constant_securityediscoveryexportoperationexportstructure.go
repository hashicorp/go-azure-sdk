package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationExportStructure string

const (
	SecurityEdiscoveryExportOperationExportStructure_Directory SecurityEdiscoveryExportOperationExportStructure = "directory"
	SecurityEdiscoveryExportOperationExportStructure_None      SecurityEdiscoveryExportOperationExportStructure = "none"
	SecurityEdiscoveryExportOperationExportStructure_Pst       SecurityEdiscoveryExportOperationExportStructure = "pst"
)

func PossibleValuesForSecurityEdiscoveryExportOperationExportStructure() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationExportStructure_Directory),
		string(SecurityEdiscoveryExportOperationExportStructure_None),
		string(SecurityEdiscoveryExportOperationExportStructure_Pst),
	}
}

func (s *SecurityEdiscoveryExportOperationExportStructure) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryExportOperationExportStructure(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryExportOperationExportStructure(input string) (*SecurityEdiscoveryExportOperationExportStructure, error) {
	vals := map[string]SecurityEdiscoveryExportOperationExportStructure{
		"directory": SecurityEdiscoveryExportOperationExportStructure_Directory,
		"none":      SecurityEdiscoveryExportOperationExportStructure_None,
		"pst":       SecurityEdiscoveryExportOperationExportStructure_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationExportStructure(input)
	return &out, nil
}
