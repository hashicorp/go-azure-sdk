package vmwares

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DhcpTypeEnum string

const (
	DhcpTypeEnumRELAY  DhcpTypeEnum = "RELAY"
	DhcpTypeEnumSERVER DhcpTypeEnum = "SERVER"
)

func PossibleValuesForDhcpTypeEnum() []string {
	return []string{
		string(DhcpTypeEnumRELAY),
		string(DhcpTypeEnumSERVER),
	}
}

func (s *DhcpTypeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDhcpTypeEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDhcpTypeEnum(input string) (*DhcpTypeEnum, error) {
	vals := map[string]DhcpTypeEnum{
		"relay":  DhcpTypeEnumRELAY,
		"server": DhcpTypeEnumSERVER,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DhcpTypeEnum(input)
	return &out, nil
}

type DnsServiceLogLevelEnum string

const (
	DnsServiceLogLevelEnumDEBUG   DnsServiceLogLevelEnum = "DEBUG"
	DnsServiceLogLevelEnumERROR   DnsServiceLogLevelEnum = "ERROR"
	DnsServiceLogLevelEnumFATAL   DnsServiceLogLevelEnum = "FATAL"
	DnsServiceLogLevelEnumINFO    DnsServiceLogLevelEnum = "INFO"
	DnsServiceLogLevelEnumWARNING DnsServiceLogLevelEnum = "WARNING"
)

func PossibleValuesForDnsServiceLogLevelEnum() []string {
	return []string{
		string(DnsServiceLogLevelEnumDEBUG),
		string(DnsServiceLogLevelEnumERROR),
		string(DnsServiceLogLevelEnumFATAL),
		string(DnsServiceLogLevelEnumINFO),
		string(DnsServiceLogLevelEnumWARNING),
	}
}

func (s *DnsServiceLogLevelEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDnsServiceLogLevelEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDnsServiceLogLevelEnum(input string) (*DnsServiceLogLevelEnum, error) {
	vals := map[string]DnsServiceLogLevelEnum{
		"debug":   DnsServiceLogLevelEnumDEBUG,
		"error":   DnsServiceLogLevelEnumERROR,
		"fatal":   DnsServiceLogLevelEnumFATAL,
		"info":    DnsServiceLogLevelEnumINFO,
		"warning": DnsServiceLogLevelEnumWARNING,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsServiceLogLevelEnum(input)
	return &out, nil
}

type DnsServiceStatusEnum string

const (
	DnsServiceStatusEnumFAILURE DnsServiceStatusEnum = "FAILURE"
	DnsServiceStatusEnumSUCCESS DnsServiceStatusEnum = "SUCCESS"
)

func PossibleValuesForDnsServiceStatusEnum() []string {
	return []string{
		string(DnsServiceStatusEnumFAILURE),
		string(DnsServiceStatusEnumSUCCESS),
	}
}

func (s *DnsServiceStatusEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDnsServiceStatusEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDnsServiceStatusEnum(input string) (*DnsServiceStatusEnum, error) {
	vals := map[string]DnsServiceStatusEnum{
		"failure": DnsServiceStatusEnumFAILURE,
		"success": DnsServiceStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsServiceStatusEnum(input)
	return &out, nil
}

type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnumBIDIRECTIONAL PortMirroringDirectionEnum = "BIDIRECTIONAL"
	PortMirroringDirectionEnumEGRESS        PortMirroringDirectionEnum = "EGRESS"
	PortMirroringDirectionEnumINGRESS       PortMirroringDirectionEnum = "INGRESS"
)

func PossibleValuesForPortMirroringDirectionEnum() []string {
	return []string{
		string(PortMirroringDirectionEnumBIDIRECTIONAL),
		string(PortMirroringDirectionEnumEGRESS),
		string(PortMirroringDirectionEnumINGRESS),
	}
}

func (s *PortMirroringDirectionEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePortMirroringDirectionEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePortMirroringDirectionEnum(input string) (*PortMirroringDirectionEnum, error) {
	vals := map[string]PortMirroringDirectionEnum{
		"bidirectional": PortMirroringDirectionEnumBIDIRECTIONAL,
		"egress":        PortMirroringDirectionEnumEGRESS,
		"ingress":       PortMirroringDirectionEnumINGRESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PortMirroringDirectionEnum(input)
	return &out, nil
}

type PortMirroringStatusEnum string

const (
	PortMirroringStatusEnumFAILURE PortMirroringStatusEnum = "FAILURE"
	PortMirroringStatusEnumSUCCESS PortMirroringStatusEnum = "SUCCESS"
)

func PossibleValuesForPortMirroringStatusEnum() []string {
	return []string{
		string(PortMirroringStatusEnumFAILURE),
		string(PortMirroringStatusEnumSUCCESS),
	}
}

func (s *PortMirroringStatusEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePortMirroringStatusEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePortMirroringStatusEnum(input string) (*PortMirroringStatusEnum, error) {
	vals := map[string]PortMirroringStatusEnum{
		"failure": PortMirroringStatusEnumFAILURE,
		"success": PortMirroringStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PortMirroringStatusEnum(input)
	return &out, nil
}

type VMGroupStatusEnum string

const (
	VMGroupStatusEnumFAILURE VMGroupStatusEnum = "FAILURE"
	VMGroupStatusEnumSUCCESS VMGroupStatusEnum = "SUCCESS"
)

func PossibleValuesForVMGroupStatusEnum() []string {
	return []string{
		string(VMGroupStatusEnumFAILURE),
		string(VMGroupStatusEnumSUCCESS),
	}
}

func (s *VMGroupStatusEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVMGroupStatusEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVMGroupStatusEnum(input string) (*VMGroupStatusEnum, error) {
	vals := map[string]VMGroupStatusEnum{
		"failure": VMGroupStatusEnumFAILURE,
		"success": VMGroupStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGroupStatusEnum(input)
	return &out, nil
}

type WorkloadNetworkDhcpProvisioningState string

const (
	WorkloadNetworkDhcpProvisioningStateBuilding  WorkloadNetworkDhcpProvisioningState = "Building"
	WorkloadNetworkDhcpProvisioningStateCanceled  WorkloadNetworkDhcpProvisioningState = "Canceled"
	WorkloadNetworkDhcpProvisioningStateDeleting  WorkloadNetworkDhcpProvisioningState = "Deleting"
	WorkloadNetworkDhcpProvisioningStateFailed    WorkloadNetworkDhcpProvisioningState = "Failed"
	WorkloadNetworkDhcpProvisioningStateSucceeded WorkloadNetworkDhcpProvisioningState = "Succeeded"
	WorkloadNetworkDhcpProvisioningStateUpdating  WorkloadNetworkDhcpProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDhcpProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDhcpProvisioningStateBuilding),
		string(WorkloadNetworkDhcpProvisioningStateCanceled),
		string(WorkloadNetworkDhcpProvisioningStateDeleting),
		string(WorkloadNetworkDhcpProvisioningStateFailed),
		string(WorkloadNetworkDhcpProvisioningStateSucceeded),
		string(WorkloadNetworkDhcpProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkDhcpProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkDhcpProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkDhcpProvisioningState(input string) (*WorkloadNetworkDhcpProvisioningState, error) {
	vals := map[string]WorkloadNetworkDhcpProvisioningState{
		"building":  WorkloadNetworkDhcpProvisioningStateBuilding,
		"canceled":  WorkloadNetworkDhcpProvisioningStateCanceled,
		"deleting":  WorkloadNetworkDhcpProvisioningStateDeleting,
		"failed":    WorkloadNetworkDhcpProvisioningStateFailed,
		"succeeded": WorkloadNetworkDhcpProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDhcpProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDhcpProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkDnsServiceProvisioningState string

const (
	WorkloadNetworkDnsServiceProvisioningStateBuilding  WorkloadNetworkDnsServiceProvisioningState = "Building"
	WorkloadNetworkDnsServiceProvisioningStateCanceled  WorkloadNetworkDnsServiceProvisioningState = "Canceled"
	WorkloadNetworkDnsServiceProvisioningStateDeleting  WorkloadNetworkDnsServiceProvisioningState = "Deleting"
	WorkloadNetworkDnsServiceProvisioningStateFailed    WorkloadNetworkDnsServiceProvisioningState = "Failed"
	WorkloadNetworkDnsServiceProvisioningStateSucceeded WorkloadNetworkDnsServiceProvisioningState = "Succeeded"
	WorkloadNetworkDnsServiceProvisioningStateUpdating  WorkloadNetworkDnsServiceProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsServiceProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsServiceProvisioningStateBuilding),
		string(WorkloadNetworkDnsServiceProvisioningStateCanceled),
		string(WorkloadNetworkDnsServiceProvisioningStateDeleting),
		string(WorkloadNetworkDnsServiceProvisioningStateFailed),
		string(WorkloadNetworkDnsServiceProvisioningStateSucceeded),
		string(WorkloadNetworkDnsServiceProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkDnsServiceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkDnsServiceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkDnsServiceProvisioningState(input string) (*WorkloadNetworkDnsServiceProvisioningState, error) {
	vals := map[string]WorkloadNetworkDnsServiceProvisioningState{
		"building":  WorkloadNetworkDnsServiceProvisioningStateBuilding,
		"canceled":  WorkloadNetworkDnsServiceProvisioningStateCanceled,
		"deleting":  WorkloadNetworkDnsServiceProvisioningStateDeleting,
		"failed":    WorkloadNetworkDnsServiceProvisioningStateFailed,
		"succeeded": WorkloadNetworkDnsServiceProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDnsServiceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDnsServiceProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkDnsZoneProvisioningState string

const (
	WorkloadNetworkDnsZoneProvisioningStateBuilding  WorkloadNetworkDnsZoneProvisioningState = "Building"
	WorkloadNetworkDnsZoneProvisioningStateCanceled  WorkloadNetworkDnsZoneProvisioningState = "Canceled"
	WorkloadNetworkDnsZoneProvisioningStateDeleting  WorkloadNetworkDnsZoneProvisioningState = "Deleting"
	WorkloadNetworkDnsZoneProvisioningStateFailed    WorkloadNetworkDnsZoneProvisioningState = "Failed"
	WorkloadNetworkDnsZoneProvisioningStateSucceeded WorkloadNetworkDnsZoneProvisioningState = "Succeeded"
	WorkloadNetworkDnsZoneProvisioningStateUpdating  WorkloadNetworkDnsZoneProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsZoneProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsZoneProvisioningStateBuilding),
		string(WorkloadNetworkDnsZoneProvisioningStateCanceled),
		string(WorkloadNetworkDnsZoneProvisioningStateDeleting),
		string(WorkloadNetworkDnsZoneProvisioningStateFailed),
		string(WorkloadNetworkDnsZoneProvisioningStateSucceeded),
		string(WorkloadNetworkDnsZoneProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkDnsZoneProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkDnsZoneProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkDnsZoneProvisioningState(input string) (*WorkloadNetworkDnsZoneProvisioningState, error) {
	vals := map[string]WorkloadNetworkDnsZoneProvisioningState{
		"building":  WorkloadNetworkDnsZoneProvisioningStateBuilding,
		"canceled":  WorkloadNetworkDnsZoneProvisioningStateCanceled,
		"deleting":  WorkloadNetworkDnsZoneProvisioningStateDeleting,
		"failed":    WorkloadNetworkDnsZoneProvisioningStateFailed,
		"succeeded": WorkloadNetworkDnsZoneProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDnsZoneProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDnsZoneProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkPortMirroringProvisioningState string

const (
	WorkloadNetworkPortMirroringProvisioningStateBuilding  WorkloadNetworkPortMirroringProvisioningState = "Building"
	WorkloadNetworkPortMirroringProvisioningStateCanceled  WorkloadNetworkPortMirroringProvisioningState = "Canceled"
	WorkloadNetworkPortMirroringProvisioningStateDeleting  WorkloadNetworkPortMirroringProvisioningState = "Deleting"
	WorkloadNetworkPortMirroringProvisioningStateFailed    WorkloadNetworkPortMirroringProvisioningState = "Failed"
	WorkloadNetworkPortMirroringProvisioningStateSucceeded WorkloadNetworkPortMirroringProvisioningState = "Succeeded"
	WorkloadNetworkPortMirroringProvisioningStateUpdating  WorkloadNetworkPortMirroringProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPortMirroringProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPortMirroringProvisioningStateBuilding),
		string(WorkloadNetworkPortMirroringProvisioningStateCanceled),
		string(WorkloadNetworkPortMirroringProvisioningStateDeleting),
		string(WorkloadNetworkPortMirroringProvisioningStateFailed),
		string(WorkloadNetworkPortMirroringProvisioningStateSucceeded),
		string(WorkloadNetworkPortMirroringProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkPortMirroringProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkPortMirroringProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkPortMirroringProvisioningState(input string) (*WorkloadNetworkPortMirroringProvisioningState, error) {
	vals := map[string]WorkloadNetworkPortMirroringProvisioningState{
		"building":  WorkloadNetworkPortMirroringProvisioningStateBuilding,
		"canceled":  WorkloadNetworkPortMirroringProvisioningStateCanceled,
		"deleting":  WorkloadNetworkPortMirroringProvisioningStateDeleting,
		"failed":    WorkloadNetworkPortMirroringProvisioningStateFailed,
		"succeeded": WorkloadNetworkPortMirroringProvisioningStateSucceeded,
		"updating":  WorkloadNetworkPortMirroringProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkPortMirroringProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkPublicIPProvisioningState string

const (
	WorkloadNetworkPublicIPProvisioningStateBuilding  WorkloadNetworkPublicIPProvisioningState = "Building"
	WorkloadNetworkPublicIPProvisioningStateCanceled  WorkloadNetworkPublicIPProvisioningState = "Canceled"
	WorkloadNetworkPublicIPProvisioningStateDeleting  WorkloadNetworkPublicIPProvisioningState = "Deleting"
	WorkloadNetworkPublicIPProvisioningStateFailed    WorkloadNetworkPublicIPProvisioningState = "Failed"
	WorkloadNetworkPublicIPProvisioningStateSucceeded WorkloadNetworkPublicIPProvisioningState = "Succeeded"
	WorkloadNetworkPublicIPProvisioningStateUpdating  WorkloadNetworkPublicIPProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPublicIPProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPublicIPProvisioningStateBuilding),
		string(WorkloadNetworkPublicIPProvisioningStateCanceled),
		string(WorkloadNetworkPublicIPProvisioningStateDeleting),
		string(WorkloadNetworkPublicIPProvisioningStateFailed),
		string(WorkloadNetworkPublicIPProvisioningStateSucceeded),
		string(WorkloadNetworkPublicIPProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkPublicIPProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkPublicIPProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkPublicIPProvisioningState(input string) (*WorkloadNetworkPublicIPProvisioningState, error) {
	vals := map[string]WorkloadNetworkPublicIPProvisioningState{
		"building":  WorkloadNetworkPublicIPProvisioningStateBuilding,
		"canceled":  WorkloadNetworkPublicIPProvisioningStateCanceled,
		"deleting":  WorkloadNetworkPublicIPProvisioningStateDeleting,
		"failed":    WorkloadNetworkPublicIPProvisioningStateFailed,
		"succeeded": WorkloadNetworkPublicIPProvisioningStateSucceeded,
		"updating":  WorkloadNetworkPublicIPProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkPublicIPProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkVMGroupProvisioningState string

const (
	WorkloadNetworkVMGroupProvisioningStateBuilding  WorkloadNetworkVMGroupProvisioningState = "Building"
	WorkloadNetworkVMGroupProvisioningStateCanceled  WorkloadNetworkVMGroupProvisioningState = "Canceled"
	WorkloadNetworkVMGroupProvisioningStateDeleting  WorkloadNetworkVMGroupProvisioningState = "Deleting"
	WorkloadNetworkVMGroupProvisioningStateFailed    WorkloadNetworkVMGroupProvisioningState = "Failed"
	WorkloadNetworkVMGroupProvisioningStateSucceeded WorkloadNetworkVMGroupProvisioningState = "Succeeded"
	WorkloadNetworkVMGroupProvisioningStateUpdating  WorkloadNetworkVMGroupProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkVMGroupProvisioningState() []string {
	return []string{
		string(WorkloadNetworkVMGroupProvisioningStateBuilding),
		string(WorkloadNetworkVMGroupProvisioningStateCanceled),
		string(WorkloadNetworkVMGroupProvisioningStateDeleting),
		string(WorkloadNetworkVMGroupProvisioningStateFailed),
		string(WorkloadNetworkVMGroupProvisioningStateSucceeded),
		string(WorkloadNetworkVMGroupProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkVMGroupProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkVMGroupProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkVMGroupProvisioningState(input string) (*WorkloadNetworkVMGroupProvisioningState, error) {
	vals := map[string]WorkloadNetworkVMGroupProvisioningState{
		"building":  WorkloadNetworkVMGroupProvisioningStateBuilding,
		"canceled":  WorkloadNetworkVMGroupProvisioningStateCanceled,
		"deleting":  WorkloadNetworkVMGroupProvisioningStateDeleting,
		"failed":    WorkloadNetworkVMGroupProvisioningStateFailed,
		"succeeded": WorkloadNetworkVMGroupProvisioningStateSucceeded,
		"updating":  WorkloadNetworkVMGroupProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkVMGroupProvisioningState(input)
	return &out, nil
}
