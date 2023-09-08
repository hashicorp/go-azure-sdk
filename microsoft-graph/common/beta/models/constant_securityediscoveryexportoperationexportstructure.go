package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationExportStructure string

const (
	SecurityEdiscoveryExportOperationExportStructuredirectory SecurityEdiscoveryExportOperationExportStructure = "Directory"
	SecurityEdiscoveryExportOperationExportStructurenone      SecurityEdiscoveryExportOperationExportStructure = "None"
	SecurityEdiscoveryExportOperationExportStructurepst       SecurityEdiscoveryExportOperationExportStructure = "Pst"
)

func PossibleValuesForSecurityEdiscoveryExportOperationExportStructure() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationExportStructuredirectory),
		string(SecurityEdiscoveryExportOperationExportStructurenone),
		string(SecurityEdiscoveryExportOperationExportStructurepst),
	}
}

func parseSecurityEdiscoveryExportOperationExportStructure(input string) (*SecurityEdiscoveryExportOperationExportStructure, error) {
	vals := map[string]SecurityEdiscoveryExportOperationExportStructure{
		"directory": SecurityEdiscoveryExportOperationExportStructuredirectory,
		"none":      SecurityEdiscoveryExportOperationExportStructurenone,
		"pst":       SecurityEdiscoveryExportOperationExportStructurepst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationExportStructure(input)
	return &out, nil
}
