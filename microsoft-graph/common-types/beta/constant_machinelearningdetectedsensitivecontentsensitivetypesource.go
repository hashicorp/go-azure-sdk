package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentSensitiveTypeSource string

const (
	MachineLearningDetectedSensitiveContentSensitiveTypeSource_OutOfBox MachineLearningDetectedSensitiveContentSensitiveTypeSource = "outOfBox"
	MachineLearningDetectedSensitiveContentSensitiveTypeSource_Tenant   MachineLearningDetectedSensitiveContentSensitiveTypeSource = "tenant"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentSensitiveTypeSource() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentSensitiveTypeSource_OutOfBox),
		string(MachineLearningDetectedSensitiveContentSensitiveTypeSource_Tenant),
	}
}

func (s *MachineLearningDetectedSensitiveContentSensitiveTypeSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineLearningDetectedSensitiveContentSensitiveTypeSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineLearningDetectedSensitiveContentSensitiveTypeSource(input string) (*MachineLearningDetectedSensitiveContentSensitiveTypeSource, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentSensitiveTypeSource{
		"outofbox": MachineLearningDetectedSensitiveContentSensitiveTypeSource_OutOfBox,
		"tenant":   MachineLearningDetectedSensitiveContentSensitiveTypeSource_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentSensitiveTypeSource(input)
	return &out, nil
}
