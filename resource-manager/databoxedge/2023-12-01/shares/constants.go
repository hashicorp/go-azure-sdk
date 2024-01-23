package shares

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureContainerDataFormat string

const (
	AzureContainerDataFormatAzureFile AzureContainerDataFormat = "AzureFile"
	AzureContainerDataFormatBlockBlob AzureContainerDataFormat = "BlockBlob"
	AzureContainerDataFormatPageBlob  AzureContainerDataFormat = "PageBlob"
)

func PossibleValuesForAzureContainerDataFormat() []string {
	return []string{
		string(AzureContainerDataFormatAzureFile),
		string(AzureContainerDataFormatBlockBlob),
		string(AzureContainerDataFormatPageBlob),
	}
}

func (s *AzureContainerDataFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureContainerDataFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureContainerDataFormat(input string) (*AzureContainerDataFormat, error) {
	vals := map[string]AzureContainerDataFormat{
		"azurefile": AzureContainerDataFormatAzureFile,
		"blockblob": AzureContainerDataFormatBlockBlob,
		"pageblob":  AzureContainerDataFormatPageBlob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureContainerDataFormat(input)
	return &out, nil
}

type ClientPermissionType string

const (
	ClientPermissionTypeNoAccess  ClientPermissionType = "NoAccess"
	ClientPermissionTypeReadOnly  ClientPermissionType = "ReadOnly"
	ClientPermissionTypeReadWrite ClientPermissionType = "ReadWrite"
)

func PossibleValuesForClientPermissionType() []string {
	return []string{
		string(ClientPermissionTypeNoAccess),
		string(ClientPermissionTypeReadOnly),
		string(ClientPermissionTypeReadWrite),
	}
}

func (s *ClientPermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClientPermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClientPermissionType(input string) (*ClientPermissionType, error) {
	vals := map[string]ClientPermissionType{
		"noaccess":  ClientPermissionTypeNoAccess,
		"readonly":  ClientPermissionTypeReadOnly,
		"readwrite": ClientPermissionTypeReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientPermissionType(input)
	return &out, nil
}

type DataPolicy string

const (
	DataPolicyCloud DataPolicy = "Cloud"
	DataPolicyLocal DataPolicy = "Local"
)

func PossibleValuesForDataPolicy() []string {
	return []string{
		string(DataPolicyCloud),
		string(DataPolicyLocal),
	}
}

func (s *DataPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataPolicy(input string) (*DataPolicy, error) {
	vals := map[string]DataPolicy{
		"cloud": DataPolicyCloud,
		"local": DataPolicyLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataPolicy(input)
	return &out, nil
}

type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

func PossibleValuesForMonitoringStatus() []string {
	return []string{
		string(MonitoringStatusDisabled),
		string(MonitoringStatusEnabled),
	}
}

func (s *MonitoringStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringStatus(input string) (*MonitoringStatus, error) {
	vals := map[string]MonitoringStatus{
		"disabled": MonitoringStatusDisabled,
		"enabled":  MonitoringStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringStatus(input)
	return &out, nil
}

type MountType string

const (
	MountTypeHostPath MountType = "HostPath"
	MountTypeVolume   MountType = "Volume"
)

func PossibleValuesForMountType() []string {
	return []string{
		string(MountTypeHostPath),
		string(MountTypeVolume),
	}
}

func (s *MountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMountType(input string) (*MountType, error) {
	vals := map[string]MountType{
		"hostpath": MountTypeHostPath,
		"volume":   MountTypeVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MountType(input)
	return &out, nil
}

type RoleTypes string

const (
	RoleTypesASA                 RoleTypes = "ASA"
	RoleTypesCloudEdgeManagement RoleTypes = "CloudEdgeManagement"
	RoleTypesCognitive           RoleTypes = "Cognitive"
	RoleTypesFunctions           RoleTypes = "Functions"
	RoleTypesIOT                 RoleTypes = "IOT"
	RoleTypesKubernetes          RoleTypes = "Kubernetes"
	RoleTypesMEC                 RoleTypes = "MEC"
)

func PossibleValuesForRoleTypes() []string {
	return []string{
		string(RoleTypesASA),
		string(RoleTypesCloudEdgeManagement),
		string(RoleTypesCognitive),
		string(RoleTypesFunctions),
		string(RoleTypesIOT),
		string(RoleTypesKubernetes),
		string(RoleTypesMEC),
	}
}

func (s *RoleTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleTypes(input string) (*RoleTypes, error) {
	vals := map[string]RoleTypes{
		"asa":                 RoleTypesASA,
		"cloudedgemanagement": RoleTypesCloudEdgeManagement,
		"cognitive":           RoleTypesCognitive,
		"functions":           RoleTypesFunctions,
		"iot":                 RoleTypesIOT,
		"kubernetes":          RoleTypesKubernetes,
		"mec":                 RoleTypesMEC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleTypes(input)
	return &out, nil
}

type ShareAccessProtocol string

const (
	ShareAccessProtocolNFS ShareAccessProtocol = "NFS"
	ShareAccessProtocolSMB ShareAccessProtocol = "SMB"
)

func PossibleValuesForShareAccessProtocol() []string {
	return []string{
		string(ShareAccessProtocolNFS),
		string(ShareAccessProtocolSMB),
	}
}

func (s *ShareAccessProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShareAccessProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShareAccessProtocol(input string) (*ShareAccessProtocol, error) {
	vals := map[string]ShareAccessProtocol{
		"nfs": ShareAccessProtocolNFS,
		"smb": ShareAccessProtocolSMB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareAccessProtocol(input)
	return &out, nil
}

type ShareAccessType string

const (
	ShareAccessTypeChange ShareAccessType = "Change"
	ShareAccessTypeCustom ShareAccessType = "Custom"
	ShareAccessTypeRead   ShareAccessType = "Read"
)

func PossibleValuesForShareAccessType() []string {
	return []string{
		string(ShareAccessTypeChange),
		string(ShareAccessTypeCustom),
		string(ShareAccessTypeRead),
	}
}

func (s *ShareAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShareAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShareAccessType(input string) (*ShareAccessType, error) {
	vals := map[string]ShareAccessType{
		"change": ShareAccessTypeChange,
		"custom": ShareAccessTypeCustom,
		"read":   ShareAccessTypeRead,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareAccessType(input)
	return &out, nil
}

type ShareStatus string

const (
	ShareStatusNeedsAttention ShareStatus = "NeedsAttention"
	ShareStatusOK             ShareStatus = "OK"
	ShareStatusOffline        ShareStatus = "Offline"
	ShareStatusUnknown        ShareStatus = "Unknown"
	ShareStatusUpdating       ShareStatus = "Updating"
)

func PossibleValuesForShareStatus() []string {
	return []string{
		string(ShareStatusNeedsAttention),
		string(ShareStatusOK),
		string(ShareStatusOffline),
		string(ShareStatusUnknown),
		string(ShareStatusUpdating),
	}
}

func (s *ShareStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShareStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShareStatus(input string) (*ShareStatus, error) {
	vals := map[string]ShareStatus{
		"needsattention": ShareStatusNeedsAttention,
		"ok":             ShareStatusOK,
		"offline":        ShareStatusOffline,
		"unknown":        ShareStatusUnknown,
		"updating":       ShareStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareStatus(input)
	return &out, nil
}
