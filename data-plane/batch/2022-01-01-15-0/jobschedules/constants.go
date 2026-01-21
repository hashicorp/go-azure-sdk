package jobschedules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessScope string

const (
	AccessScopeJob AccessScope = "job"
)

func PossibleValuesForAccessScope() []string {
	return []string{
		string(AccessScopeJob),
	}
}

func (s *AccessScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessScope(input string) (*AccessScope, error) {
	vals := map[string]AccessScope{
		"job": AccessScopeJob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessScope(input)
	return &out, nil
}

type AutoUserScope string

const (
	AutoUserScopePool AutoUserScope = "pool"
	AutoUserScopeTask AutoUserScope = "task"
)

func PossibleValuesForAutoUserScope() []string {
	return []string{
		string(AutoUserScopePool),
		string(AutoUserScopeTask),
	}
}

func (s *AutoUserScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoUserScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoUserScope(input string) (*AutoUserScope, error) {
	vals := map[string]AutoUserScope{
		"pool": AutoUserScopePool,
		"task": AutoUserScopeTask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoUserScope(input)
	return &out, nil
}

type CachingType string

const (
	CachingTypeNone      CachingType = "none"
	CachingTypeReadonly  CachingType = "readonly"
	CachingTypeReadwrite CachingType = "readwrite"
)

func PossibleValuesForCachingType() []string {
	return []string{
		string(CachingTypeNone),
		string(CachingTypeReadonly),
		string(CachingTypeReadwrite),
	}
}

func (s *CachingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCachingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCachingType(input string) (*CachingType, error) {
	vals := map[string]CachingType{
		"none":      CachingTypeNone,
		"readonly":  CachingTypeReadonly,
		"readwrite": CachingTypeReadwrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CachingType(input)
	return &out, nil
}

type CertificateStoreLocation string

const (
	CertificateStoreLocationCurrentuser  CertificateStoreLocation = "currentuser"
	CertificateStoreLocationLocalmachine CertificateStoreLocation = "localmachine"
)

func PossibleValuesForCertificateStoreLocation() []string {
	return []string{
		string(CertificateStoreLocationCurrentuser),
		string(CertificateStoreLocationLocalmachine),
	}
}

func (s *CertificateStoreLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateStoreLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateStoreLocation(input string) (*CertificateStoreLocation, error) {
	vals := map[string]CertificateStoreLocation{
		"currentuser":  CertificateStoreLocationCurrentuser,
		"localmachine": CertificateStoreLocationLocalmachine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateStoreLocation(input)
	return &out, nil
}

type CertificateVisibility string

const (
	CertificateVisibilityRemoteuser CertificateVisibility = "remoteuser"
	CertificateVisibilityStarttask  CertificateVisibility = "starttask"
	CertificateVisibilityTask       CertificateVisibility = "task"
)

func PossibleValuesForCertificateVisibility() []string {
	return []string{
		string(CertificateVisibilityRemoteuser),
		string(CertificateVisibilityStarttask),
		string(CertificateVisibilityTask),
	}
}

func (s *CertificateVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateVisibility(input string) (*CertificateVisibility, error) {
	vals := map[string]CertificateVisibility{
		"remoteuser": CertificateVisibilityRemoteuser,
		"starttask":  CertificateVisibilityStarttask,
		"task":       CertificateVisibilityTask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateVisibility(input)
	return &out, nil
}

type ComputeNodeFillType string

const (
	ComputeNodeFillTypePack   ComputeNodeFillType = "pack"
	ComputeNodeFillTypeSpread ComputeNodeFillType = "spread"
)

func PossibleValuesForComputeNodeFillType() []string {
	return []string{
		string(ComputeNodeFillTypePack),
		string(ComputeNodeFillTypeSpread),
	}
}

func (s *ComputeNodeFillType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeFillType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeFillType(input string) (*ComputeNodeFillType, error) {
	vals := map[string]ComputeNodeFillType{
		"pack":   ComputeNodeFillTypePack,
		"spread": ComputeNodeFillTypeSpread,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeFillType(input)
	return &out, nil
}

type ContainerType string

const (
	ContainerTypeDockerCompatible ContainerType = "dockerCompatible"
)

func PossibleValuesForContainerType() []string {
	return []string{
		string(ContainerTypeDockerCompatible),
	}
}

func (s *ContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerType(input string) (*ContainerType, error) {
	vals := map[string]ContainerType{
		"dockercompatible": ContainerTypeDockerCompatible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerType(input)
	return &out, nil
}

type ContainerWorkingDirectory string

const (
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = "containerImageDefault"
	ContainerWorkingDirectoryTaskWorkingDirectory  ContainerWorkingDirectory = "taskWorkingDirectory"
)

func PossibleValuesForContainerWorkingDirectory() []string {
	return []string{
		string(ContainerWorkingDirectoryContainerImageDefault),
		string(ContainerWorkingDirectoryTaskWorkingDirectory),
	}
}

func (s *ContainerWorkingDirectory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerWorkingDirectory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerWorkingDirectory(input string) (*ContainerWorkingDirectory, error) {
	vals := map[string]ContainerWorkingDirectory{
		"containerimagedefault": ContainerWorkingDirectoryContainerImageDefault,
		"taskworkingdirectory":  ContainerWorkingDirectoryTaskWorkingDirectory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerWorkingDirectory(input)
	return &out, nil
}

type DiffDiskPlacement string

const (
	DiffDiskPlacementCacheDisk DiffDiskPlacement = "CacheDisk"
)

func PossibleValuesForDiffDiskPlacement() []string {
	return []string{
		string(DiffDiskPlacementCacheDisk),
	}
}

func (s *DiffDiskPlacement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiffDiskPlacement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiffDiskPlacement(input string) (*DiffDiskPlacement, error) {
	vals := map[string]DiffDiskPlacement{
		"cachedisk": DiffDiskPlacementCacheDisk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiffDiskPlacement(input)
	return &out, nil
}

type DiskEncryptionTarget string

const (
	DiskEncryptionTargetOsdisk        DiskEncryptionTarget = "osdisk"
	DiskEncryptionTargetTemporarydisk DiskEncryptionTarget = "temporarydisk"
)

func PossibleValuesForDiskEncryptionTarget() []string {
	return []string{
		string(DiskEncryptionTargetOsdisk),
		string(DiskEncryptionTargetTemporarydisk),
	}
}

func (s *DiskEncryptionTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiskEncryptionTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiskEncryptionTarget(input string) (*DiskEncryptionTarget, error) {
	vals := map[string]DiskEncryptionTarget{
		"osdisk":        DiskEncryptionTargetOsdisk,
		"temporarydisk": DiskEncryptionTargetTemporarydisk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskEncryptionTarget(input)
	return &out, nil
}

type DynamicVNetAssignmentScope string

const (
	DynamicVNetAssignmentScopeJob  DynamicVNetAssignmentScope = "job"
	DynamicVNetAssignmentScopeNone DynamicVNetAssignmentScope = "none"
)

func PossibleValuesForDynamicVNetAssignmentScope() []string {
	return []string{
		string(DynamicVNetAssignmentScopeJob),
		string(DynamicVNetAssignmentScopeNone),
	}
}

func (s *DynamicVNetAssignmentScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDynamicVNetAssignmentScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDynamicVNetAssignmentScope(input string) (*DynamicVNetAssignmentScope, error) {
	vals := map[string]DynamicVNetAssignmentScope{
		"job":  DynamicVNetAssignmentScopeJob,
		"none": DynamicVNetAssignmentScopeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DynamicVNetAssignmentScope(input)
	return &out, nil
}

type ElevationLevel string

const (
	ElevationLevelAdmin    ElevationLevel = "admin"
	ElevationLevelNonadmin ElevationLevel = "nonadmin"
)

func PossibleValuesForElevationLevel() []string {
	return []string{
		string(ElevationLevelAdmin),
		string(ElevationLevelNonadmin),
	}
}

func (s *ElevationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElevationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElevationLevel(input string) (*ElevationLevel, error) {
	vals := map[string]ElevationLevel{
		"admin":    ElevationLevelAdmin,
		"nonadmin": ElevationLevelNonadmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElevationLevel(input)
	return &out, nil
}

type IPAddressProvisioningType string

const (
	IPAddressProvisioningTypeBatchmanaged        IPAddressProvisioningType = "batchmanaged"
	IPAddressProvisioningTypeNopublicipaddresses IPAddressProvisioningType = "nopublicipaddresses"
	IPAddressProvisioningTypeUsermanaged         IPAddressProvisioningType = "usermanaged"
)

func PossibleValuesForIPAddressProvisioningType() []string {
	return []string{
		string(IPAddressProvisioningTypeBatchmanaged),
		string(IPAddressProvisioningTypeNopublicipaddresses),
		string(IPAddressProvisioningTypeUsermanaged),
	}
}

func (s *IPAddressProvisioningType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPAddressProvisioningType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPAddressProvisioningType(input string) (*IPAddressProvisioningType, error) {
	vals := map[string]IPAddressProvisioningType{
		"batchmanaged":        IPAddressProvisioningTypeBatchmanaged,
		"nopublicipaddresses": IPAddressProvisioningTypeNopublicipaddresses,
		"usermanaged":         IPAddressProvisioningTypeUsermanaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAddressProvisioningType(input)
	return &out, nil
}

type InboundEndpointProtocol string

const (
	InboundEndpointProtocolTcp InboundEndpointProtocol = "tcp"
	InboundEndpointProtocolUdp InboundEndpointProtocol = "udp"
)

func PossibleValuesForInboundEndpointProtocol() []string {
	return []string{
		string(InboundEndpointProtocolTcp),
		string(InboundEndpointProtocolUdp),
	}
}

func (s *InboundEndpointProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInboundEndpointProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInboundEndpointProtocol(input string) (*InboundEndpointProtocol, error) {
	vals := map[string]InboundEndpointProtocol{
		"tcp": InboundEndpointProtocolTcp,
		"udp": InboundEndpointProtocolUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InboundEndpointProtocol(input)
	return &out, nil
}

type JobScheduleState string

const (
	JobScheduleStateActive      JobScheduleState = "active"
	JobScheduleStateCompleted   JobScheduleState = "completed"
	JobScheduleStateDeleting    JobScheduleState = "deleting"
	JobScheduleStateDisabled    JobScheduleState = "disabled"
	JobScheduleStateTerminating JobScheduleState = "terminating"
)

func PossibleValuesForJobScheduleState() []string {
	return []string{
		string(JobScheduleStateActive),
		string(JobScheduleStateCompleted),
		string(JobScheduleStateDeleting),
		string(JobScheduleStateDisabled),
		string(JobScheduleStateTerminating),
	}
}

func (s *JobScheduleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobScheduleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobScheduleState(input string) (*JobScheduleState, error) {
	vals := map[string]JobScheduleState{
		"active":      JobScheduleStateActive,
		"completed":   JobScheduleStateCompleted,
		"deleting":    JobScheduleStateDeleting,
		"disabled":    JobScheduleStateDisabled,
		"terminating": JobScheduleStateTerminating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobScheduleState(input)
	return &out, nil
}

type LoginMode string

const (
	LoginModeBatch       LoginMode = "batch"
	LoginModeInteractive LoginMode = "interactive"
)

func PossibleValuesForLoginMode() []string {
	return []string{
		string(LoginModeBatch),
		string(LoginModeInteractive),
	}
}

func (s *LoginMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginMode(input string) (*LoginMode, error) {
	vals := map[string]LoginMode{
		"batch":       LoginModeBatch,
		"interactive": LoginModeInteractive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginMode(input)
	return &out, nil
}

type NetworkSecurityGroupRuleAccess string

const (
	NetworkSecurityGroupRuleAccessAllow NetworkSecurityGroupRuleAccess = "allow"
	NetworkSecurityGroupRuleAccessDeny  NetworkSecurityGroupRuleAccess = "deny"
)

func PossibleValuesForNetworkSecurityGroupRuleAccess() []string {
	return []string{
		string(NetworkSecurityGroupRuleAccessAllow),
		string(NetworkSecurityGroupRuleAccessDeny),
	}
}

func (s *NetworkSecurityGroupRuleAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkSecurityGroupRuleAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkSecurityGroupRuleAccess(input string) (*NetworkSecurityGroupRuleAccess, error) {
	vals := map[string]NetworkSecurityGroupRuleAccess{
		"allow": NetworkSecurityGroupRuleAccessAllow,
		"deny":  NetworkSecurityGroupRuleAccessDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityGroupRuleAccess(input)
	return &out, nil
}

type NodePlacementPolicyType string

const (
	NodePlacementPolicyTypeRegional NodePlacementPolicyType = "regional"
	NodePlacementPolicyTypeZonal    NodePlacementPolicyType = "zonal"
)

func PossibleValuesForNodePlacementPolicyType() []string {
	return []string{
		string(NodePlacementPolicyTypeRegional),
		string(NodePlacementPolicyTypeZonal),
	}
}

func (s *NodePlacementPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNodePlacementPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNodePlacementPolicyType(input string) (*NodePlacementPolicyType, error) {
	vals := map[string]NodePlacementPolicyType{
		"regional": NodePlacementPolicyTypeRegional,
		"zonal":    NodePlacementPolicyTypeZonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodePlacementPolicyType(input)
	return &out, nil
}

type OnAllTasksComplete string

const (
	OnAllTasksCompleteNoaction     OnAllTasksComplete = "noaction"
	OnAllTasksCompleteTerminatejob OnAllTasksComplete = "terminatejob"
)

func PossibleValuesForOnAllTasksComplete() []string {
	return []string{
		string(OnAllTasksCompleteNoaction),
		string(OnAllTasksCompleteTerminatejob),
	}
}

func (s *OnAllTasksComplete) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnAllTasksComplete(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnAllTasksComplete(input string) (*OnAllTasksComplete, error) {
	vals := map[string]OnAllTasksComplete{
		"noaction":     OnAllTasksCompleteNoaction,
		"terminatejob": OnAllTasksCompleteTerminatejob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnAllTasksComplete(input)
	return &out, nil
}

type OnTaskFailure string

const (
	OnTaskFailureNoaction                    OnTaskFailure = "noaction"
	OnTaskFailurePerformexitoptionsjobaction OnTaskFailure = "performexitoptionsjobaction"
)

func PossibleValuesForOnTaskFailure() []string {
	return []string{
		string(OnTaskFailureNoaction),
		string(OnTaskFailurePerformexitoptionsjobaction),
	}
}

func (s *OnTaskFailure) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnTaskFailure(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnTaskFailure(input string) (*OnTaskFailure, error) {
	vals := map[string]OnTaskFailure{
		"noaction":                    OnTaskFailureNoaction,
		"performexitoptionsjobaction": OnTaskFailurePerformexitoptionsjobaction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnTaskFailure(input)
	return &out, nil
}

type OutputFileUploadCondition string

const (
	OutputFileUploadConditionTaskcompletion OutputFileUploadCondition = "taskcompletion"
	OutputFileUploadConditionTaskfailure    OutputFileUploadCondition = "taskfailure"
	OutputFileUploadConditionTasksuccess    OutputFileUploadCondition = "tasksuccess"
)

func PossibleValuesForOutputFileUploadCondition() []string {
	return []string{
		string(OutputFileUploadConditionTaskcompletion),
		string(OutputFileUploadConditionTaskfailure),
		string(OutputFileUploadConditionTasksuccess),
	}
}

func (s *OutputFileUploadCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutputFileUploadCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutputFileUploadCondition(input string) (*OutputFileUploadCondition, error) {
	vals := map[string]OutputFileUploadCondition{
		"taskcompletion": OutputFileUploadConditionTaskcompletion,
		"taskfailure":    OutputFileUploadConditionTaskfailure,
		"tasksuccess":    OutputFileUploadConditionTasksuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputFileUploadCondition(input)
	return &out, nil
}

type PoolLifetimeOption string

const (
	PoolLifetimeOptionJob         PoolLifetimeOption = "job"
	PoolLifetimeOptionJobschedule PoolLifetimeOption = "jobschedule"
)

func PossibleValuesForPoolLifetimeOption() []string {
	return []string{
		string(PoolLifetimeOptionJob),
		string(PoolLifetimeOptionJobschedule),
	}
}

func (s *PoolLifetimeOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolLifetimeOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolLifetimeOption(input string) (*PoolLifetimeOption, error) {
	vals := map[string]PoolLifetimeOption{
		"job":         PoolLifetimeOptionJob,
		"jobschedule": PoolLifetimeOptionJobschedule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolLifetimeOption(input)
	return &out, nil
}

type StorageAccountType string

const (
	StorageAccountTypePremiumLrs  StorageAccountType = "premium_lrs"
	StorageAccountTypeStandardLrs StorageAccountType = "standard_lrs"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypePremiumLrs),
		string(StorageAccountTypeStandardLrs),
	}
}

func (s *StorageAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAccountType(input string) (*StorageAccountType, error) {
	vals := map[string]StorageAccountType{
		"premium_lrs":  StorageAccountTypePremiumLrs,
		"standard_lrs": StorageAccountTypeStandardLrs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountType(input)
	return &out, nil
}
