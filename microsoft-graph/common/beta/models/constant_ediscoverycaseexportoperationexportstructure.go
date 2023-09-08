package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationExportStructure string

const (
	EdiscoveryCaseExportOperationExportStructuredirectory EdiscoveryCaseExportOperationExportStructure = "Directory"
	EdiscoveryCaseExportOperationExportStructurenone      EdiscoveryCaseExportOperationExportStructure = "None"
	EdiscoveryCaseExportOperationExportStructurepst       EdiscoveryCaseExportOperationExportStructure = "Pst"
)

func PossibleValuesForEdiscoveryCaseExportOperationExportStructure() []string {
	return []string{
		string(EdiscoveryCaseExportOperationExportStructuredirectory),
		string(EdiscoveryCaseExportOperationExportStructurenone),
		string(EdiscoveryCaseExportOperationExportStructurepst),
	}
}

func parseEdiscoveryCaseExportOperationExportStructure(input string) (*EdiscoveryCaseExportOperationExportStructure, error) {
	vals := map[string]EdiscoveryCaseExportOperationExportStructure{
		"directory": EdiscoveryCaseExportOperationExportStructuredirectory,
		"none":      EdiscoveryCaseExportOperationExportStructurenone,
		"pst":       EdiscoveryCaseExportOperationExportStructurepst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationExportStructure(input)
	return &out, nil
}
