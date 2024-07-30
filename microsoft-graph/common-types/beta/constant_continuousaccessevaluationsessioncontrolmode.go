package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousAccessEvaluationSessionControlMode string

const (
	ContinuousAccessEvaluationSessionControlMode_Disabled          ContinuousAccessEvaluationSessionControlMode = "disabled"
	ContinuousAccessEvaluationSessionControlMode_StrictEnforcement ContinuousAccessEvaluationSessionControlMode = "strictEnforcement"
	ContinuousAccessEvaluationSessionControlMode_StrictLocation    ContinuousAccessEvaluationSessionControlMode = "strictLocation"
)

func PossibleValuesForContinuousAccessEvaluationSessionControlMode() []string {
	return []string{
		string(ContinuousAccessEvaluationSessionControlMode_Disabled),
		string(ContinuousAccessEvaluationSessionControlMode_StrictEnforcement),
		string(ContinuousAccessEvaluationSessionControlMode_StrictLocation),
	}
}

func (s *ContinuousAccessEvaluationSessionControlMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContinuousAccessEvaluationSessionControlMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContinuousAccessEvaluationSessionControlMode(input string) (*ContinuousAccessEvaluationSessionControlMode, error) {
	vals := map[string]ContinuousAccessEvaluationSessionControlMode{
		"disabled":          ContinuousAccessEvaluationSessionControlMode_Disabled,
		"strictenforcement": ContinuousAccessEvaluationSessionControlMode_StrictEnforcement,
		"strictlocation":    ContinuousAccessEvaluationSessionControlMode_StrictLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContinuousAccessEvaluationSessionControlMode(input)
	return &out, nil
}
