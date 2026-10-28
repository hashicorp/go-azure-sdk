package origingroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthProbeRequestType string

const (
	HealthProbeRequestTypeGET    HealthProbeRequestType = "GET"
	HealthProbeRequestTypeHEAD   HealthProbeRequestType = "HEAD"
	HealthProbeRequestTypeNotSet HealthProbeRequestType = "NotSet"
)

func PossibleValuesForHealthProbeRequestType() []string {
	return []string{
		string(HealthProbeRequestTypeGET),
		string(HealthProbeRequestTypeHEAD),
		string(HealthProbeRequestTypeNotSet),
	}
}

func (s *HealthProbeRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthProbeRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthProbeRequestType(input string) (*HealthProbeRequestType, error) {
	vals := map[string]HealthProbeRequestType{
		"get":    HealthProbeRequestTypeGET,
		"head":   HealthProbeRequestTypeHEAD,
		"notset": HealthProbeRequestTypeNotSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthProbeRequestType(input)
	return &out, nil
}

type OriginGroupProvisioningState string

const (
	OriginGroupProvisioningStateCreating  OriginGroupProvisioningState = "Creating"
	OriginGroupProvisioningStateDeleting  OriginGroupProvisioningState = "Deleting"
	OriginGroupProvisioningStateFailed    OriginGroupProvisioningState = "Failed"
	OriginGroupProvisioningStateSucceeded OriginGroupProvisioningState = "Succeeded"
	OriginGroupProvisioningStateUpdating  OriginGroupProvisioningState = "Updating"
)

func PossibleValuesForOriginGroupProvisioningState() []string {
	return []string{
		string(OriginGroupProvisioningStateCreating),
		string(OriginGroupProvisioningStateDeleting),
		string(OriginGroupProvisioningStateFailed),
		string(OriginGroupProvisioningStateSucceeded),
		string(OriginGroupProvisioningStateUpdating),
	}
}

func (s *OriginGroupProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOriginGroupProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOriginGroupProvisioningState(input string) (*OriginGroupProvisioningState, error) {
	vals := map[string]OriginGroupProvisioningState{
		"creating":  OriginGroupProvisioningStateCreating,
		"deleting":  OriginGroupProvisioningStateDeleting,
		"failed":    OriginGroupProvisioningStateFailed,
		"succeeded": OriginGroupProvisioningStateSucceeded,
		"updating":  OriginGroupProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OriginGroupProvisioningState(input)
	return &out, nil
}

type OriginGroupResourceState string

const (
	OriginGroupResourceStateActive   OriginGroupResourceState = "Active"
	OriginGroupResourceStateCreating OriginGroupResourceState = "Creating"
	OriginGroupResourceStateDeleting OriginGroupResourceState = "Deleting"
)

func PossibleValuesForOriginGroupResourceState() []string {
	return []string{
		string(OriginGroupResourceStateActive),
		string(OriginGroupResourceStateCreating),
		string(OriginGroupResourceStateDeleting),
	}
}

func (s *OriginGroupResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOriginGroupResourceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOriginGroupResourceState(input string) (*OriginGroupResourceState, error) {
	vals := map[string]OriginGroupResourceState{
		"active":   OriginGroupResourceStateActive,
		"creating": OriginGroupResourceStateCreating,
		"deleting": OriginGroupResourceStateDeleting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OriginGroupResourceState(input)
	return &out, nil
}

type ProbeProtocol string

const (
	ProbeProtocolHTTP   ProbeProtocol = "Http"
	ProbeProtocolHTTPS  ProbeProtocol = "Https"
	ProbeProtocolNotSet ProbeProtocol = "NotSet"
)

func PossibleValuesForProbeProtocol() []string {
	return []string{
		string(ProbeProtocolHTTP),
		string(ProbeProtocolHTTPS),
		string(ProbeProtocolNotSet),
	}
}

func (s *ProbeProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProbeProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProbeProtocol(input string) (*ProbeProtocol, error) {
	vals := map[string]ProbeProtocol{
		"http":   ProbeProtocolHTTP,
		"https":  ProbeProtocolHTTPS,
		"notset": ProbeProtocolNotSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProbeProtocol(input)
	return &out, nil
}

type ResponseBasedDetectedErrorTypes string

const (
	ResponseBasedDetectedErrorTypesNone             ResponseBasedDetectedErrorTypes = "None"
	ResponseBasedDetectedErrorTypesTcpAndHTTPErrors ResponseBasedDetectedErrorTypes = "TcpAndHttpErrors"
	ResponseBasedDetectedErrorTypesTcpErrorsOnly    ResponseBasedDetectedErrorTypes = "TcpErrorsOnly"
)

func PossibleValuesForResponseBasedDetectedErrorTypes() []string {
	return []string{
		string(ResponseBasedDetectedErrorTypesNone),
		string(ResponseBasedDetectedErrorTypesTcpAndHTTPErrors),
		string(ResponseBasedDetectedErrorTypesTcpErrorsOnly),
	}
}

func (s *ResponseBasedDetectedErrorTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseBasedDetectedErrorTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseBasedDetectedErrorTypes(input string) (*ResponseBasedDetectedErrorTypes, error) {
	vals := map[string]ResponseBasedDetectedErrorTypes{
		"none":             ResponseBasedDetectedErrorTypesNone,
		"tcpandhttperrors": ResponseBasedDetectedErrorTypesTcpAndHTTPErrors,
		"tcperrorsonly":    ResponseBasedDetectedErrorTypesTcpErrorsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseBasedDetectedErrorTypes(input)
	return &out, nil
}
