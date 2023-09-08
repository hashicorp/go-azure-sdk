package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskPreviewType string

const (
	BusinessScenarioTaskPreviewTypeautomatic   BusinessScenarioTaskPreviewType = "Automatic"
	BusinessScenarioTaskPreviewTypechecklist   BusinessScenarioTaskPreviewType = "Checklist"
	BusinessScenarioTaskPreviewTypedescription BusinessScenarioTaskPreviewType = "Description"
	BusinessScenarioTaskPreviewTypenoPreview   BusinessScenarioTaskPreviewType = "NoPreview"
	BusinessScenarioTaskPreviewTypereference   BusinessScenarioTaskPreviewType = "Reference"
)

func PossibleValuesForBusinessScenarioTaskPreviewType() []string {
	return []string{
		string(BusinessScenarioTaskPreviewTypeautomatic),
		string(BusinessScenarioTaskPreviewTypechecklist),
		string(BusinessScenarioTaskPreviewTypedescription),
		string(BusinessScenarioTaskPreviewTypenoPreview),
		string(BusinessScenarioTaskPreviewTypereference),
	}
}

func parseBusinessScenarioTaskPreviewType(input string) (*BusinessScenarioTaskPreviewType, error) {
	vals := map[string]BusinessScenarioTaskPreviewType{
		"automatic":   BusinessScenarioTaskPreviewTypeautomatic,
		"checklist":   BusinessScenarioTaskPreviewTypechecklist,
		"description": BusinessScenarioTaskPreviewTypedescription,
		"nopreview":   BusinessScenarioTaskPreviewTypenoPreview,
		"reference":   BusinessScenarioTaskPreviewTypereference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskPreviewType(input)
	return &out, nil
}
