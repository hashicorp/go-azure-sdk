package pools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllocationState string

const (
	AllocationStateResizing AllocationState = "resizing"
	AllocationStateSteady   AllocationState = "steady"
	AllocationStateStopping AllocationState = "stopping"
)

func PossibleValuesForAllocationState() []string {
	return []string{
		string(AllocationStateResizing),
		string(AllocationStateSteady),
		string(AllocationStateStopping),
	}
}

func (s *AllocationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllocationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllocationState(input string) (*AllocationState, error) {
	vals := map[string]AllocationState{
		"resizing": AllocationStateResizing,
		"steady":   AllocationStateSteady,
		"stopping": AllocationStateStopping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllocationState(input)
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

type ComputeNodeDeallocationOption string

const (
	ComputeNodeDeallocationOptionRequeue        ComputeNodeDeallocationOption = "requeue"
	ComputeNodeDeallocationOptionRetaineddata   ComputeNodeDeallocationOption = "retaineddata"
	ComputeNodeDeallocationOptionTaskcompletion ComputeNodeDeallocationOption = "taskcompletion"
	ComputeNodeDeallocationOptionTerminate      ComputeNodeDeallocationOption = "terminate"
)

func PossibleValuesForComputeNodeDeallocationOption() []string {
	return []string{
		string(ComputeNodeDeallocationOptionRequeue),
		string(ComputeNodeDeallocationOptionRetaineddata),
		string(ComputeNodeDeallocationOptionTaskcompletion),
		string(ComputeNodeDeallocationOptionTerminate),
	}
}

func (s *ComputeNodeDeallocationOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeDeallocationOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeDeallocationOption(input string) (*ComputeNodeDeallocationOption, error) {
	vals := map[string]ComputeNodeDeallocationOption{
		"requeue":        ComputeNodeDeallocationOptionRequeue,
		"retaineddata":   ComputeNodeDeallocationOptionRetaineddata,
		"taskcompletion": ComputeNodeDeallocationOptionTaskcompletion,
		"terminate":      ComputeNodeDeallocationOptionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeDeallocationOption(input)
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

type PoolIdentityType string

const (
	PoolIdentityTypeNone         PoolIdentityType = "None"
	PoolIdentityTypeUserAssigned PoolIdentityType = "UserAssigned"
)

func PossibleValuesForPoolIdentityType() []string {
	return []string{
		string(PoolIdentityTypeNone),
		string(PoolIdentityTypeUserAssigned),
	}
}

func (s *PoolIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolIdentityType(input string) (*PoolIdentityType, error) {
	vals := map[string]PoolIdentityType{
		"none":         PoolIdentityTypeNone,
		"userassigned": PoolIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolIdentityType(input)
	return &out, nil
}

type PoolState string

const (
	PoolStateActive   PoolState = "active"
	PoolStateDeleting PoolState = "deleting"
)

func PossibleValuesForPoolState() []string {
	return []string{
		string(PoolStateActive),
		string(PoolStateDeleting),
	}
}

func (s *PoolState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolState(input string) (*PoolState, error) {
	vals := map[string]PoolState{
		"active":   PoolStateActive,
		"deleting": PoolStateDeleting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolState(input)
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
