package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLearningSummaryApplicationType string

const (
	WindowsInformationProtectionAppLearningSummaryApplicationType_Desktop   WindowsInformationProtectionAppLearningSummaryApplicationType = "desktop"
	WindowsInformationProtectionAppLearningSummaryApplicationType_Universal WindowsInformationProtectionAppLearningSummaryApplicationType = "universal"
)

func PossibleValuesForWindowsInformationProtectionAppLearningSummaryApplicationType() []string {
	return []string{
		string(WindowsInformationProtectionAppLearningSummaryApplicationType_Desktop),
		string(WindowsInformationProtectionAppLearningSummaryApplicationType_Universal),
	}
}

func (s *WindowsInformationProtectionAppLearningSummaryApplicationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionAppLearningSummaryApplicationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionAppLearningSummaryApplicationType(input string) (*WindowsInformationProtectionAppLearningSummaryApplicationType, error) {
	vals := map[string]WindowsInformationProtectionAppLearningSummaryApplicationType{
		"desktop":   WindowsInformationProtectionAppLearningSummaryApplicationType_Desktop,
		"universal": WindowsInformationProtectionAppLearningSummaryApplicationType_Universal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionAppLearningSummaryApplicationType(input)
	return &out, nil
}
