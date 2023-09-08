package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousAccessEvaluationSessionControlMode string

const (
	ContinuousAccessEvaluationSessionControlModedisabled          ContinuousAccessEvaluationSessionControlMode = "Disabled"
	ContinuousAccessEvaluationSessionControlModestrictEnforcement ContinuousAccessEvaluationSessionControlMode = "StrictEnforcement"
	ContinuousAccessEvaluationSessionControlModestrictLocation    ContinuousAccessEvaluationSessionControlMode = "StrictLocation"
)

func PossibleValuesForContinuousAccessEvaluationSessionControlMode() []string {
	return []string{
		string(ContinuousAccessEvaluationSessionControlModedisabled),
		string(ContinuousAccessEvaluationSessionControlModestrictEnforcement),
		string(ContinuousAccessEvaluationSessionControlModestrictLocation),
	}
}

func parseContinuousAccessEvaluationSessionControlMode(input string) (*ContinuousAccessEvaluationSessionControlMode, error) {
	vals := map[string]ContinuousAccessEvaluationSessionControlMode{
		"disabled":          ContinuousAccessEvaluationSessionControlModedisabled,
		"strictenforcement": ContinuousAccessEvaluationSessionControlModestrictEnforcement,
		"strictlocation":    ContinuousAccessEvaluationSessionControlModestrictLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContinuousAccessEvaluationSessionControlMode(input)
	return &out, nil
}
