package containerapps

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple ActiveRevisionsMode = "multiple"
	ActiveRevisionsModeSingle   ActiveRevisionsMode = "single"
)

func PossibleValuesForActiveRevisionsMode() []string {
	return []string{
		string(ActiveRevisionsModeMultiple),
		string(ActiveRevisionsModeSingle),
	}
}

func (s *ActiveRevisionsMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActiveRevisionsMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActiveRevisionsMode(input string) (*ActiveRevisionsMode, error) {
	vals := map[string]ActiveRevisionsMode{
		"multiple": ActiveRevisionsModeMultiple,
		"single":   ActiveRevisionsModeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActiveRevisionsMode(input)
	return &out, nil
}

type ContainerAppProvisioningState string

const (
	ContainerAppProvisioningStateCanceled   ContainerAppProvisioningState = "Canceled"
	ContainerAppProvisioningStateFailed     ContainerAppProvisioningState = "Failed"
	ContainerAppProvisioningStateInProgress ContainerAppProvisioningState = "InProgress"
	ContainerAppProvisioningStateSucceeded  ContainerAppProvisioningState = "Succeeded"
)

func PossibleValuesForContainerAppProvisioningState() []string {
	return []string{
		string(ContainerAppProvisioningStateCanceled),
		string(ContainerAppProvisioningStateFailed),
		string(ContainerAppProvisioningStateInProgress),
		string(ContainerAppProvisioningStateSucceeded),
	}
}

func (s *ContainerAppProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerAppProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerAppProvisioningState(input string) (*ContainerAppProvisioningState, error) {
	vals := map[string]ContainerAppProvisioningState{
		"canceled":   ContainerAppProvisioningStateCanceled,
		"failed":     ContainerAppProvisioningStateFailed,
		"inprogress": ContainerAppProvisioningStateInProgress,
		"succeeded":  ContainerAppProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerAppProvisioningState(input)
	return &out, nil
}

type IngressTransportMethod string

const (
	IngressTransportMethodAuto    IngressTransportMethod = "auto"
	IngressTransportMethodHTTP    IngressTransportMethod = "http"
	IngressTransportMethodHTTPTwo IngressTransportMethod = "http2"
)

func PossibleValuesForIngressTransportMethod() []string {
	return []string{
		string(IngressTransportMethodAuto),
		string(IngressTransportMethodHTTP),
		string(IngressTransportMethodHTTPTwo),
	}
}

func (s *IngressTransportMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIngressTransportMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIngressTransportMethod(input string) (*IngressTransportMethod, error) {
	vals := map[string]IngressTransportMethod{
		"auto":  IngressTransportMethodAuto,
		"http":  IngressTransportMethodHTTP,
		"http2": IngressTransportMethodHTTPTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IngressTransportMethod(input)
	return &out, nil
}
