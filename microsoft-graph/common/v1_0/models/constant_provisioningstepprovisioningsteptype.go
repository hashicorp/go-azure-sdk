package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStepProvisioningStepType string

const (
	ProvisioningStepProvisioningStepTypeexport              ProvisioningStepProvisioningStepType = "Export"
	ProvisioningStepProvisioningStepTypeimport              ProvisioningStepProvisioningStepType = "Import"
	ProvisioningStepProvisioningStepTypematching            ProvisioningStepProvisioningStepType = "Matching"
	ProvisioningStepProvisioningStepTypeprocessing          ProvisioningStepProvisioningStepType = "Processing"
	ProvisioningStepProvisioningStepTypereferenceResolution ProvisioningStepProvisioningStepType = "ReferenceResolution"
	ProvisioningStepProvisioningStepTypescoping             ProvisioningStepProvisioningStepType = "Scoping"
)

func PossibleValuesForProvisioningStepProvisioningStepType() []string {
	return []string{
		string(ProvisioningStepProvisioningStepTypeexport),
		string(ProvisioningStepProvisioningStepTypeimport),
		string(ProvisioningStepProvisioningStepTypematching),
		string(ProvisioningStepProvisioningStepTypeprocessing),
		string(ProvisioningStepProvisioningStepTypereferenceResolution),
		string(ProvisioningStepProvisioningStepTypescoping),
	}
}

func parseProvisioningStepProvisioningStepType(input string) (*ProvisioningStepProvisioningStepType, error) {
	vals := map[string]ProvisioningStepProvisioningStepType{
		"export":              ProvisioningStepProvisioningStepTypeexport,
		"import":              ProvisioningStepProvisioningStepTypeimport,
		"matching":            ProvisioningStepProvisioningStepTypematching,
		"processing":          ProvisioningStepProvisioningStepTypeprocessing,
		"referenceresolution": ProvisioningStepProvisioningStepTypereferenceResolution,
		"scoping":             ProvisioningStepProvisioningStepTypescoping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStepProvisioningStepType(input)
	return &out, nil
}
