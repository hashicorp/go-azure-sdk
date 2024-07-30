package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskPreviewType string

const (
	BusinessScenarioTaskPreviewType_Automatic   BusinessScenarioTaskPreviewType = "automatic"
	BusinessScenarioTaskPreviewType_Checklist   BusinessScenarioTaskPreviewType = "checklist"
	BusinessScenarioTaskPreviewType_Description BusinessScenarioTaskPreviewType = "description"
	BusinessScenarioTaskPreviewType_NoPreview   BusinessScenarioTaskPreviewType = "noPreview"
	BusinessScenarioTaskPreviewType_Reference   BusinessScenarioTaskPreviewType = "reference"
)

func PossibleValuesForBusinessScenarioTaskPreviewType() []string {
	return []string{
		string(BusinessScenarioTaskPreviewType_Automatic),
		string(BusinessScenarioTaskPreviewType_Checklist),
		string(BusinessScenarioTaskPreviewType_Description),
		string(BusinessScenarioTaskPreviewType_NoPreview),
		string(BusinessScenarioTaskPreviewType_Reference),
	}
}

func (s *BusinessScenarioTaskPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBusinessScenarioTaskPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBusinessScenarioTaskPreviewType(input string) (*BusinessScenarioTaskPreviewType, error) {
	vals := map[string]BusinessScenarioTaskPreviewType{
		"automatic":   BusinessScenarioTaskPreviewType_Automatic,
		"checklist":   BusinessScenarioTaskPreviewType_Checklist,
		"description": BusinessScenarioTaskPreviewType_Description,
		"nopreview":   BusinessScenarioTaskPreviewType_NoPreview,
		"reference":   BusinessScenarioTaskPreviewType_Reference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskPreviewType(input)
	return &out, nil
}
