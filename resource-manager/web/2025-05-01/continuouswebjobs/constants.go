package continuouswebjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousWebJobStatus string

const (
	ContinuousWebJobStatusInitializing   ContinuousWebJobStatus = "Initializing"
	ContinuousWebJobStatusPendingRestart ContinuousWebJobStatus = "PendingRestart"
	ContinuousWebJobStatusRunning        ContinuousWebJobStatus = "Running"
	ContinuousWebJobStatusStarting       ContinuousWebJobStatus = "Starting"
	ContinuousWebJobStatusStopped        ContinuousWebJobStatus = "Stopped"
)

func PossibleValuesForContinuousWebJobStatus() []string {
	return []string{
		string(ContinuousWebJobStatusInitializing),
		string(ContinuousWebJobStatusPendingRestart),
		string(ContinuousWebJobStatusRunning),
		string(ContinuousWebJobStatusStarting),
		string(ContinuousWebJobStatusStopped),
	}
}

func (s *ContinuousWebJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContinuousWebJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContinuousWebJobStatus(input string) (*ContinuousWebJobStatus, error) {
	vals := map[string]ContinuousWebJobStatus{
		"initializing":   ContinuousWebJobStatusInitializing,
		"pendingrestart": ContinuousWebJobStatusPendingRestart,
		"running":        ContinuousWebJobStatusRunning,
		"starting":       ContinuousWebJobStatusStarting,
		"stopped":        ContinuousWebJobStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContinuousWebJobStatus(input)
	return &out, nil
}

type WebJobType string

const (
	WebJobTypeContinuous WebJobType = "Continuous"
	WebJobTypeTriggered  WebJobType = "Triggered"
)

func PossibleValuesForWebJobType() []string {
	return []string{
		string(WebJobTypeContinuous),
		string(WebJobTypeTriggered),
	}
}

func (s *WebJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebJobType(input string) (*WebJobType, error) {
	vals := map[string]WebJobType{
		"continuous": WebJobTypeContinuous,
		"triggered":  WebJobTypeTriggered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebJobType(input)
	return &out, nil
}
