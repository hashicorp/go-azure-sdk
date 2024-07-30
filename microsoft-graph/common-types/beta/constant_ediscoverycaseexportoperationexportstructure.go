package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationExportStructure string

const (
	EdiscoveryCaseExportOperationExportStructure_Directory EdiscoveryCaseExportOperationExportStructure = "directory"
	EdiscoveryCaseExportOperationExportStructure_None      EdiscoveryCaseExportOperationExportStructure = "none"
	EdiscoveryCaseExportOperationExportStructure_Pst       EdiscoveryCaseExportOperationExportStructure = "pst"
)

func PossibleValuesForEdiscoveryCaseExportOperationExportStructure() []string {
	return []string{
		string(EdiscoveryCaseExportOperationExportStructure_Directory),
		string(EdiscoveryCaseExportOperationExportStructure_None),
		string(EdiscoveryCaseExportOperationExportStructure_Pst),
	}
}

func (s *EdiscoveryCaseExportOperationExportStructure) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseExportOperationExportStructure(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseExportOperationExportStructure(input string) (*EdiscoveryCaseExportOperationExportStructure, error) {
	vals := map[string]EdiscoveryCaseExportOperationExportStructure{
		"directory": EdiscoveryCaseExportOperationExportStructure_Directory,
		"none":      EdiscoveryCaseExportOperationExportStructure_None,
		"pst":       EdiscoveryCaseExportOperationExportStructure_Pst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationExportStructure(input)
	return &out, nil
}
