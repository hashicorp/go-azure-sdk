package dataflowgraph

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphConnectionSchemaSerializationFormat string

const (
	DataflowGraphConnectionSchemaSerializationFormatAvro    DataflowGraphConnectionSchemaSerializationFormat = "Avro"
	DataflowGraphConnectionSchemaSerializationFormatDelta   DataflowGraphConnectionSchemaSerializationFormat = "Delta"
	DataflowGraphConnectionSchemaSerializationFormatJson    DataflowGraphConnectionSchemaSerializationFormat = "Json"
	DataflowGraphConnectionSchemaSerializationFormatParquet DataflowGraphConnectionSchemaSerializationFormat = "Parquet"
)

func PossibleValuesForDataflowGraphConnectionSchemaSerializationFormat() []string {
	return []string{
		string(DataflowGraphConnectionSchemaSerializationFormatAvro),
		string(DataflowGraphConnectionSchemaSerializationFormatDelta),
		string(DataflowGraphConnectionSchemaSerializationFormatJson),
		string(DataflowGraphConnectionSchemaSerializationFormatParquet),
	}
}

func (s *DataflowGraphConnectionSchemaSerializationFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataflowGraphConnectionSchemaSerializationFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataflowGraphConnectionSchemaSerializationFormat(input string) (*DataflowGraphConnectionSchemaSerializationFormat, error) {
	vals := map[string]DataflowGraphConnectionSchemaSerializationFormat{
		"avro":    DataflowGraphConnectionSchemaSerializationFormatAvro,
		"delta":   DataflowGraphConnectionSchemaSerializationFormatDelta,
		"json":    DataflowGraphConnectionSchemaSerializationFormatJson,
		"parquet": DataflowGraphConnectionSchemaSerializationFormatParquet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataflowGraphConnectionSchemaSerializationFormat(input)
	return &out, nil
}

type DataflowGraphDestinationHeaderActionType string

const (
	DataflowGraphDestinationHeaderActionTypeAddIfNotPresent DataflowGraphDestinationHeaderActionType = "AddIfNotPresent"
	DataflowGraphDestinationHeaderActionTypeAddOrReplace    DataflowGraphDestinationHeaderActionType = "AddOrReplace"
	DataflowGraphDestinationHeaderActionTypeRemove          DataflowGraphDestinationHeaderActionType = "Remove"
)

func PossibleValuesForDataflowGraphDestinationHeaderActionType() []string {
	return []string{
		string(DataflowGraphDestinationHeaderActionTypeAddIfNotPresent),
		string(DataflowGraphDestinationHeaderActionTypeAddOrReplace),
		string(DataflowGraphDestinationHeaderActionTypeRemove),
	}
}

func (s *DataflowGraphDestinationHeaderActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataflowGraphDestinationHeaderActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataflowGraphDestinationHeaderActionType(input string) (*DataflowGraphDestinationHeaderActionType, error) {
	vals := map[string]DataflowGraphDestinationHeaderActionType{
		"addifnotpresent": DataflowGraphDestinationHeaderActionTypeAddIfNotPresent,
		"addorreplace":    DataflowGraphDestinationHeaderActionTypeAddOrReplace,
		"remove":          DataflowGraphDestinationHeaderActionTypeRemove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataflowGraphDestinationHeaderActionType(input)
	return &out, nil
}

type DataflowGraphNodeType string

const (
	DataflowGraphNodeTypeDestination DataflowGraphNodeType = "Destination"
	DataflowGraphNodeTypeGraph       DataflowGraphNodeType = "Graph"
	DataflowGraphNodeTypeSource      DataflowGraphNodeType = "Source"
)

func PossibleValuesForDataflowGraphNodeType() []string {
	return []string{
		string(DataflowGraphNodeTypeDestination),
		string(DataflowGraphNodeTypeGraph),
		string(DataflowGraphNodeTypeSource),
	}
}

func (s *DataflowGraphNodeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataflowGraphNodeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataflowGraphNodeType(input string) (*DataflowGraphNodeType, error) {
	vals := map[string]DataflowGraphNodeType{
		"destination": DataflowGraphNodeTypeDestination,
		"graph":       DataflowGraphNodeTypeGraph,
		"source":      DataflowGraphNodeTypeSource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataflowGraphNodeType(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type OperationalMode string

const (
	OperationalModeDisabled OperationalMode = "Disabled"
	OperationalModeEnabled  OperationalMode = "Enabled"
)

func PossibleValuesForOperationalMode() []string {
	return []string{
		string(OperationalModeDisabled),
		string(OperationalModeEnabled),
	}
}

func (s *OperationalMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationalMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationalMode(input string) (*OperationalMode, error) {
	vals := map[string]OperationalMode{
		"disabled": OperationalModeDisabled,
		"enabled":  OperationalModeEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationalMode(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ResourceHealthState string

const (
	ResourceHealthStateAvailable   ResourceHealthState = "Available"
	ResourceHealthStateDegraded    ResourceHealthState = "Degraded"
	ResourceHealthStateUnavailable ResourceHealthState = "Unavailable"
	ResourceHealthStateUnknown     ResourceHealthState = "Unknown"
)

func PossibleValuesForResourceHealthState() []string {
	return []string{
		string(ResourceHealthStateAvailable),
		string(ResourceHealthStateDegraded),
		string(ResourceHealthStateUnavailable),
		string(ResourceHealthStateUnknown),
	}
}

func (s *ResourceHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthState(input string) (*ResourceHealthState, error) {
	vals := map[string]ResourceHealthState{
		"available":   ResourceHealthStateAvailable,
		"degraded":    ResourceHealthStateDegraded,
		"unavailable": ResourceHealthStateUnavailable,
		"unknown":     ResourceHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthState(input)
	return &out, nil
}
