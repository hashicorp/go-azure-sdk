package redisenterprise

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessKeyType string

const (
	AccessKeyTypePrimary   AccessKeyType = "Primary"
	AccessKeyTypeSecondary AccessKeyType = "Secondary"
)

func PossibleValuesForAccessKeyType() []string {
	return []string{
		string(AccessKeyTypePrimary),
		string(AccessKeyTypeSecondary),
	}
}

func (s *AccessKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessKeyType(input string) (*AccessKeyType, error) {
	vals := map[string]AccessKeyType{
		"primary":   AccessKeyTypePrimary,
		"secondary": AccessKeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessKeyType(input)
	return &out, nil
}

type AccessKeysAuthentication string

const (
	AccessKeysAuthenticationDisabled AccessKeysAuthentication = "Disabled"
	AccessKeysAuthenticationEnabled  AccessKeysAuthentication = "Enabled"
)

func PossibleValuesForAccessKeysAuthentication() []string {
	return []string{
		string(AccessKeysAuthenticationDisabled),
		string(AccessKeysAuthenticationEnabled),
	}
}

func (s *AccessKeysAuthentication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessKeysAuthentication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessKeysAuthentication(input string) (*AccessKeysAuthentication, error) {
	vals := map[string]AccessKeysAuthentication{
		"disabled": AccessKeysAuthenticationDisabled,
		"enabled":  AccessKeysAuthenticationEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessKeysAuthentication(input)
	return &out, nil
}

type AofFrequency string

const (
	AofFrequencyAlways AofFrequency = "always"
	AofFrequencyOnes   AofFrequency = "1s"
)

func PossibleValuesForAofFrequency() []string {
	return []string{
		string(AofFrequencyAlways),
		string(AofFrequencyOnes),
	}
}

func (s *AofFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAofFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAofFrequency(input string) (*AofFrequency, error) {
	vals := map[string]AofFrequency{
		"always": AofFrequencyAlways,
		"1s":     AofFrequencyOnes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AofFrequency(input)
	return &out, nil
}

type ClusteringPolicy string

const (
	ClusteringPolicyEnterpriseCluster ClusteringPolicy = "EnterpriseCluster"
	ClusteringPolicyOSSCluster        ClusteringPolicy = "OSSCluster"
)

func PossibleValuesForClusteringPolicy() []string {
	return []string{
		string(ClusteringPolicyEnterpriseCluster),
		string(ClusteringPolicyOSSCluster),
	}
}

func (s *ClusteringPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusteringPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusteringPolicy(input string) (*ClusteringPolicy, error) {
	vals := map[string]ClusteringPolicy{
		"enterprisecluster": ClusteringPolicyEnterpriseCluster,
		"osscluster":        ClusteringPolicyOSSCluster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusteringPolicy(input)
	return &out, nil
}

type CmkIdentityType string

const (
	CmkIdentityTypeSystemAssignedIdentity CmkIdentityType = "systemAssignedIdentity"
	CmkIdentityTypeUserAssignedIdentity   CmkIdentityType = "userAssignedIdentity"
)

func PossibleValuesForCmkIdentityType() []string {
	return []string{
		string(CmkIdentityTypeSystemAssignedIdentity),
		string(CmkIdentityTypeUserAssignedIdentity),
	}
}

func (s *CmkIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCmkIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCmkIdentityType(input string) (*CmkIdentityType, error) {
	vals := map[string]CmkIdentityType{
		"systemassignedidentity": CmkIdentityTypeSystemAssignedIdentity,
		"userassignedidentity":   CmkIdentityTypeUserAssignedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CmkIdentityType(input)
	return &out, nil
}

type DeferUpgradeSetting string

const (
	DeferUpgradeSettingDeferred    DeferUpgradeSetting = "Deferred"
	DeferUpgradeSettingNotDeferred DeferUpgradeSetting = "NotDeferred"
)

func PossibleValuesForDeferUpgradeSetting() []string {
	return []string{
		string(DeferUpgradeSettingDeferred),
		string(DeferUpgradeSettingNotDeferred),
	}
}

func (s *DeferUpgradeSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeferUpgradeSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeferUpgradeSetting(input string) (*DeferUpgradeSetting, error) {
	vals := map[string]DeferUpgradeSetting{
		"deferred":    DeferUpgradeSettingDeferred,
		"notdeferred": DeferUpgradeSettingNotDeferred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeferUpgradeSetting(input)
	return &out, nil
}

type EvictionPolicy string

const (
	EvictionPolicyAllKeysLFU     EvictionPolicy = "AllKeysLFU"
	EvictionPolicyAllKeysLRU     EvictionPolicy = "AllKeysLRU"
	EvictionPolicyAllKeysRandom  EvictionPolicy = "AllKeysRandom"
	EvictionPolicyNoEviction     EvictionPolicy = "NoEviction"
	EvictionPolicyVolatileLFU    EvictionPolicy = "VolatileLFU"
	EvictionPolicyVolatileLRU    EvictionPolicy = "VolatileLRU"
	EvictionPolicyVolatileRandom EvictionPolicy = "VolatileRandom"
	EvictionPolicyVolatileTTL    EvictionPolicy = "VolatileTTL"
)

func PossibleValuesForEvictionPolicy() []string {
	return []string{
		string(EvictionPolicyAllKeysLFU),
		string(EvictionPolicyAllKeysLRU),
		string(EvictionPolicyAllKeysRandom),
		string(EvictionPolicyNoEviction),
		string(EvictionPolicyVolatileLFU),
		string(EvictionPolicyVolatileLRU),
		string(EvictionPolicyVolatileRandom),
		string(EvictionPolicyVolatileTTL),
	}
}

func (s *EvictionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEvictionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEvictionPolicy(input string) (*EvictionPolicy, error) {
	vals := map[string]EvictionPolicy{
		"allkeyslfu":     EvictionPolicyAllKeysLFU,
		"allkeyslru":     EvictionPolicyAllKeysLRU,
		"allkeysrandom":  EvictionPolicyAllKeysRandom,
		"noeviction":     EvictionPolicyNoEviction,
		"volatilelfu":    EvictionPolicyVolatileLFU,
		"volatilelru":    EvictionPolicyVolatileLRU,
		"volatilerandom": EvictionPolicyVolatileRandom,
		"volatilettl":    EvictionPolicyVolatileTTL,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EvictionPolicy(input)
	return &out, nil
}

type HighAvailability string

const (
	HighAvailabilityDisabled HighAvailability = "Disabled"
	HighAvailabilityEnabled  HighAvailability = "Enabled"
)

func PossibleValuesForHighAvailability() []string {
	return []string{
		string(HighAvailabilityDisabled),
		string(HighAvailabilityEnabled),
	}
}

func (s *HighAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHighAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHighAvailability(input string) (*HighAvailability, error) {
	vals := map[string]HighAvailability{
		"disabled": HighAvailabilityDisabled,
		"enabled":  HighAvailabilityEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HighAvailability(input)
	return &out, nil
}

type Kind string

const (
	KindVOne Kind = "v1"
	KindVTwo Kind = "v2"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindVOne),
		string(KindVTwo),
	}
}

func (s *Kind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKind(input string) (*Kind, error) {
	vals := map[string]Kind{
		"v1": KindVOne,
		"v2": KindVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Kind(input)
	return &out, nil
}

type LinkState string

const (
	LinkStateLinkFailed   LinkState = "LinkFailed"
	LinkStateLinked       LinkState = "Linked"
	LinkStateLinking      LinkState = "Linking"
	LinkStateUnlinkFailed LinkState = "UnlinkFailed"
	LinkStateUnlinking    LinkState = "Unlinking"
)

func PossibleValuesForLinkState() []string {
	return []string{
		string(LinkStateLinkFailed),
		string(LinkStateLinked),
		string(LinkStateLinking),
		string(LinkStateUnlinkFailed),
		string(LinkStateUnlinking),
	}
}

func (s *LinkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLinkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLinkState(input string) (*LinkState, error) {
	vals := map[string]LinkState{
		"linkfailed":   LinkStateLinkFailed,
		"linked":       LinkStateLinked,
		"linking":      LinkStateLinking,
		"unlinkfailed": LinkStateUnlinkFailed,
		"unlinking":    LinkStateUnlinking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinkState(input)
	return &out, nil
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
	}
}

func (s *PrivateEndpointConnectionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateEndpointConnectionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateEndpointConnectionProvisioningState(input string) (*PrivateEndpointConnectionProvisioningState, error) {
	vals := map[string]PrivateEndpointConnectionProvisioningState{
		"creating":  PrivateEndpointConnectionProvisioningStateCreating,
		"deleting":  PrivateEndpointConnectionProvisioningStateDeleting,
		"failed":    PrivateEndpointConnectionProvisioningStateFailed,
		"succeeded": PrivateEndpointConnectionProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointConnectionProvisioningState(input)
	return &out, nil
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

func (s *PrivateEndpointServiceConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateEndpointServiceConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateEndpointServiceConnectionStatus(input string) (*PrivateEndpointServiceConnectionStatus, error) {
	vals := map[string]PrivateEndpointServiceConnectionStatus{
		"approved": PrivateEndpointServiceConnectionStatusApproved,
		"pending":  PrivateEndpointServiceConnectionStatusPending,
		"rejected": PrivateEndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointServiceConnectionStatus(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolEncrypted Protocol = "Encrypted"
	ProtocolPlaintext Protocol = "Plaintext"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolEncrypted),
		string(ProtocolPlaintext),
	}
}

func (s *Protocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"encrypted": ProtocolEncrypted,
		"plaintext": ProtocolPlaintext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
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
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type RdbFrequency string

const (
	RdbFrequencyOneTwoh RdbFrequency = "12h"
	RdbFrequencyOneh    RdbFrequency = "1h"
	RdbFrequencySixh    RdbFrequency = "6h"
)

func PossibleValuesForRdbFrequency() []string {
	return []string{
		string(RdbFrequencyOneTwoh),
		string(RdbFrequencyOneh),
		string(RdbFrequencySixh),
	}
}

func (s *RdbFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRdbFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRdbFrequency(input string) (*RdbFrequency, error) {
	vals := map[string]RdbFrequency{
		"12h": RdbFrequencyOneTwoh,
		"1h":  RdbFrequencyOneh,
		"6h":  RdbFrequencySixh,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RdbFrequency(input)
	return &out, nil
}

type RedundancyMode string

const (
	RedundancyModeLR   RedundancyMode = "LR"
	RedundancyModeNone RedundancyMode = "None"
	RedundancyModeZR   RedundancyMode = "ZR"
)

func PossibleValuesForRedundancyMode() []string {
	return []string{
		string(RedundancyModeLR),
		string(RedundancyModeNone),
		string(RedundancyModeZR),
	}
}

func (s *RedundancyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRedundancyMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRedundancyMode(input string) (*RedundancyMode, error) {
	vals := map[string]RedundancyMode{
		"lr":   RedundancyModeLR,
		"none": RedundancyModeNone,
		"zr":   RedundancyModeZR,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RedundancyMode(input)
	return &out, nil
}

type ResourceState string

const (
	ResourceStateCreateFailed  ResourceState = "CreateFailed"
	ResourceStateCreating      ResourceState = "Creating"
	ResourceStateDeleteFailed  ResourceState = "DeleteFailed"
	ResourceStateDeleting      ResourceState = "Deleting"
	ResourceStateDisableFailed ResourceState = "DisableFailed"
	ResourceStateDisabled      ResourceState = "Disabled"
	ResourceStateDisabling     ResourceState = "Disabling"
	ResourceStateEnableFailed  ResourceState = "EnableFailed"
	ResourceStateEnabling      ResourceState = "Enabling"
	ResourceStateMoving        ResourceState = "Moving"
	ResourceStateRunning       ResourceState = "Running"
	ResourceStateScaling       ResourceState = "Scaling"
	ResourceStateScalingFailed ResourceState = "ScalingFailed"
	ResourceStateUpdateFailed  ResourceState = "UpdateFailed"
	ResourceStateUpdating      ResourceState = "Updating"
)

func PossibleValuesForResourceState() []string {
	return []string{
		string(ResourceStateCreateFailed),
		string(ResourceStateCreating),
		string(ResourceStateDeleteFailed),
		string(ResourceStateDeleting),
		string(ResourceStateDisableFailed),
		string(ResourceStateDisabled),
		string(ResourceStateDisabling),
		string(ResourceStateEnableFailed),
		string(ResourceStateEnabling),
		string(ResourceStateMoving),
		string(ResourceStateRunning),
		string(ResourceStateScaling),
		string(ResourceStateScalingFailed),
		string(ResourceStateUpdateFailed),
		string(ResourceStateUpdating),
	}
}

func (s *ResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceState(input string) (*ResourceState, error) {
	vals := map[string]ResourceState{
		"createfailed":  ResourceStateCreateFailed,
		"creating":      ResourceStateCreating,
		"deletefailed":  ResourceStateDeleteFailed,
		"deleting":      ResourceStateDeleting,
		"disablefailed": ResourceStateDisableFailed,
		"disabled":      ResourceStateDisabled,
		"disabling":     ResourceStateDisabling,
		"enablefailed":  ResourceStateEnableFailed,
		"enabling":      ResourceStateEnabling,
		"moving":        ResourceStateMoving,
		"running":       ResourceStateRunning,
		"scaling":       ResourceStateScaling,
		"scalingfailed": ResourceStateScalingFailed,
		"updatefailed":  ResourceStateUpdateFailed,
		"updating":      ResourceStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceState(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameBalancedBFive                  SkuName = "Balanced_B5"
	SkuNameBalancedBFiveHundred           SkuName = "Balanced_B500"
	SkuNameBalancedBFiveZero              SkuName = "Balanced_B50"
	SkuNameBalancedBOne                   SkuName = "Balanced_B1"
	SkuNameBalancedBOneFiveZero           SkuName = "Balanced_B150"
	SkuNameBalancedBOneHundred            SkuName = "Balanced_B100"
	SkuNameBalancedBOneThousand           SkuName = "Balanced_B1000"
	SkuNameBalancedBOneZero               SkuName = "Balanced_B10"
	SkuNameBalancedBSevenHundred          SkuName = "Balanced_B700"
	SkuNameBalancedBThree                 SkuName = "Balanced_B3"
	SkuNameBalancedBThreeFiveZero         SkuName = "Balanced_B350"
	SkuNameBalancedBTwoFiveZero           SkuName = "Balanced_B250"
	SkuNameBalancedBTwoZero               SkuName = "Balanced_B20"
	SkuNameBalancedBZero                  SkuName = "Balanced_B0"
	SkuNameComputeOptimizedXFive          SkuName = "ComputeOptimized_X5"
	SkuNameComputeOptimizedXFiveHundred   SkuName = "ComputeOptimized_X500"
	SkuNameComputeOptimizedXFiveZero      SkuName = "ComputeOptimized_X50"
	SkuNameComputeOptimizedXOneFiveZero   SkuName = "ComputeOptimized_X150"
	SkuNameComputeOptimizedXOneHundred    SkuName = "ComputeOptimized_X100"
	SkuNameComputeOptimizedXOneZero       SkuName = "ComputeOptimized_X10"
	SkuNameComputeOptimizedXSevenHundred  SkuName = "ComputeOptimized_X700"
	SkuNameComputeOptimizedXThree         SkuName = "ComputeOptimized_X3"
	SkuNameComputeOptimizedXThreeFiveZero SkuName = "ComputeOptimized_X350"
	SkuNameComputeOptimizedXTwoFiveZero   SkuName = "ComputeOptimized_X250"
	SkuNameComputeOptimizedXTwoZero       SkuName = "ComputeOptimized_X20"
	SkuNameEnterpriseEFive                SkuName = "Enterprise_E5"
	SkuNameEnterpriseEFiveZero            SkuName = "Enterprise_E50"
	SkuNameEnterpriseEFourHundred         SkuName = "Enterprise_E400"
	SkuNameEnterpriseEOne                 SkuName = "Enterprise_E1"
	SkuNameEnterpriseEOneHundred          SkuName = "Enterprise_E100"
	SkuNameEnterpriseEOneZero             SkuName = "Enterprise_E10"
	SkuNameEnterpriseETwoHundred          SkuName = "Enterprise_E200"
	SkuNameEnterpriseETwoZero             SkuName = "Enterprise_E20"
	SkuNameEnterpriseFlashFOneFiveHundred SkuName = "EnterpriseFlash_F1500"
	SkuNameEnterpriseFlashFSevenHundred   SkuName = "EnterpriseFlash_F700"
	SkuNameEnterpriseFlashFThreeHundred   SkuName = "EnterpriseFlash_F300"
	SkuNameFlashOptimizedAFiveHundred     SkuName = "FlashOptimized_A500"
	SkuNameFlashOptimizedAFourFiveHundred SkuName = "FlashOptimized_A4500"
	SkuNameFlashOptimizedAOneFiveHundred  SkuName = "FlashOptimized_A1500"
	SkuNameFlashOptimizedAOneThousand     SkuName = "FlashOptimized_A1000"
	SkuNameFlashOptimizedASevenHundred    SkuName = "FlashOptimized_A700"
	SkuNameFlashOptimizedATwoFiveZero     SkuName = "FlashOptimized_A250"
	SkuNameFlashOptimizedATwoThousand     SkuName = "FlashOptimized_A2000"
	SkuNameMemoryOptimizedMFiveHundred    SkuName = "MemoryOptimized_M500"
	SkuNameMemoryOptimizedMFiveZero       SkuName = "MemoryOptimized_M50"
	SkuNameMemoryOptimizedMOneFiveHundred SkuName = "MemoryOptimized_M1500"
	SkuNameMemoryOptimizedMOneFiveZero    SkuName = "MemoryOptimized_M150"
	SkuNameMemoryOptimizedMOneHundred     SkuName = "MemoryOptimized_M100"
	SkuNameMemoryOptimizedMOneThousand    SkuName = "MemoryOptimized_M1000"
	SkuNameMemoryOptimizedMOneZero        SkuName = "MemoryOptimized_M10"
	SkuNameMemoryOptimizedMSevenHundred   SkuName = "MemoryOptimized_M700"
	SkuNameMemoryOptimizedMThreeFiveZero  SkuName = "MemoryOptimized_M350"
	SkuNameMemoryOptimizedMTwoFiveZero    SkuName = "MemoryOptimized_M250"
	SkuNameMemoryOptimizedMTwoThousand    SkuName = "MemoryOptimized_M2000"
	SkuNameMemoryOptimizedMTwoZero        SkuName = "MemoryOptimized_M20"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBalancedBFive),
		string(SkuNameBalancedBFiveHundred),
		string(SkuNameBalancedBFiveZero),
		string(SkuNameBalancedBOne),
		string(SkuNameBalancedBOneFiveZero),
		string(SkuNameBalancedBOneHundred),
		string(SkuNameBalancedBOneThousand),
		string(SkuNameBalancedBOneZero),
		string(SkuNameBalancedBSevenHundred),
		string(SkuNameBalancedBThree),
		string(SkuNameBalancedBThreeFiveZero),
		string(SkuNameBalancedBTwoFiveZero),
		string(SkuNameBalancedBTwoZero),
		string(SkuNameBalancedBZero),
		string(SkuNameComputeOptimizedXFive),
		string(SkuNameComputeOptimizedXFiveHundred),
		string(SkuNameComputeOptimizedXFiveZero),
		string(SkuNameComputeOptimizedXOneFiveZero),
		string(SkuNameComputeOptimizedXOneHundred),
		string(SkuNameComputeOptimizedXOneZero),
		string(SkuNameComputeOptimizedXSevenHundred),
		string(SkuNameComputeOptimizedXThree),
		string(SkuNameComputeOptimizedXThreeFiveZero),
		string(SkuNameComputeOptimizedXTwoFiveZero),
		string(SkuNameComputeOptimizedXTwoZero),
		string(SkuNameEnterpriseEFive),
		string(SkuNameEnterpriseEFiveZero),
		string(SkuNameEnterpriseEFourHundred),
		string(SkuNameEnterpriseEOne),
		string(SkuNameEnterpriseEOneHundred),
		string(SkuNameEnterpriseEOneZero),
		string(SkuNameEnterpriseETwoHundred),
		string(SkuNameEnterpriseETwoZero),
		string(SkuNameEnterpriseFlashFOneFiveHundred),
		string(SkuNameEnterpriseFlashFSevenHundred),
		string(SkuNameEnterpriseFlashFThreeHundred),
		string(SkuNameFlashOptimizedAFiveHundred),
		string(SkuNameFlashOptimizedAFourFiveHundred),
		string(SkuNameFlashOptimizedAOneFiveHundred),
		string(SkuNameFlashOptimizedAOneThousand),
		string(SkuNameFlashOptimizedASevenHundred),
		string(SkuNameFlashOptimizedATwoFiveZero),
		string(SkuNameFlashOptimizedATwoThousand),
		string(SkuNameMemoryOptimizedMFiveHundred),
		string(SkuNameMemoryOptimizedMFiveZero),
		string(SkuNameMemoryOptimizedMOneFiveHundred),
		string(SkuNameMemoryOptimizedMOneFiveZero),
		string(SkuNameMemoryOptimizedMOneHundred),
		string(SkuNameMemoryOptimizedMOneThousand),
		string(SkuNameMemoryOptimizedMOneZero),
		string(SkuNameMemoryOptimizedMSevenHundred),
		string(SkuNameMemoryOptimizedMThreeFiveZero),
		string(SkuNameMemoryOptimizedMTwoFiveZero),
		string(SkuNameMemoryOptimizedMTwoThousand),
		string(SkuNameMemoryOptimizedMTwoZero),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"balanced_b5":           SkuNameBalancedBFive,
		"balanced_b500":         SkuNameBalancedBFiveHundred,
		"balanced_b50":          SkuNameBalancedBFiveZero,
		"balanced_b1":           SkuNameBalancedBOne,
		"balanced_b150":         SkuNameBalancedBOneFiveZero,
		"balanced_b100":         SkuNameBalancedBOneHundred,
		"balanced_b1000":        SkuNameBalancedBOneThousand,
		"balanced_b10":          SkuNameBalancedBOneZero,
		"balanced_b700":         SkuNameBalancedBSevenHundred,
		"balanced_b3":           SkuNameBalancedBThree,
		"balanced_b350":         SkuNameBalancedBThreeFiveZero,
		"balanced_b250":         SkuNameBalancedBTwoFiveZero,
		"balanced_b20":          SkuNameBalancedBTwoZero,
		"balanced_b0":           SkuNameBalancedBZero,
		"computeoptimized_x5":   SkuNameComputeOptimizedXFive,
		"computeoptimized_x500": SkuNameComputeOptimizedXFiveHundred,
		"computeoptimized_x50":  SkuNameComputeOptimizedXFiveZero,
		"computeoptimized_x150": SkuNameComputeOptimizedXOneFiveZero,
		"computeoptimized_x100": SkuNameComputeOptimizedXOneHundred,
		"computeoptimized_x10":  SkuNameComputeOptimizedXOneZero,
		"computeoptimized_x700": SkuNameComputeOptimizedXSevenHundred,
		"computeoptimized_x3":   SkuNameComputeOptimizedXThree,
		"computeoptimized_x350": SkuNameComputeOptimizedXThreeFiveZero,
		"computeoptimized_x250": SkuNameComputeOptimizedXTwoFiveZero,
		"computeoptimized_x20":  SkuNameComputeOptimizedXTwoZero,
		"enterprise_e5":         SkuNameEnterpriseEFive,
		"enterprise_e50":        SkuNameEnterpriseEFiveZero,
		"enterprise_e400":       SkuNameEnterpriseEFourHundred,
		"enterprise_e1":         SkuNameEnterpriseEOne,
		"enterprise_e100":       SkuNameEnterpriseEOneHundred,
		"enterprise_e10":        SkuNameEnterpriseEOneZero,
		"enterprise_e200":       SkuNameEnterpriseETwoHundred,
		"enterprise_e20":        SkuNameEnterpriseETwoZero,
		"enterpriseflash_f1500": SkuNameEnterpriseFlashFOneFiveHundred,
		"enterpriseflash_f700":  SkuNameEnterpriseFlashFSevenHundred,
		"enterpriseflash_f300":  SkuNameEnterpriseFlashFThreeHundred,
		"flashoptimized_a500":   SkuNameFlashOptimizedAFiveHundred,
		"flashoptimized_a4500":  SkuNameFlashOptimizedAFourFiveHundred,
		"flashoptimized_a1500":  SkuNameFlashOptimizedAOneFiveHundred,
		"flashoptimized_a1000":  SkuNameFlashOptimizedAOneThousand,
		"flashoptimized_a700":   SkuNameFlashOptimizedASevenHundred,
		"flashoptimized_a250":   SkuNameFlashOptimizedATwoFiveZero,
		"flashoptimized_a2000":  SkuNameFlashOptimizedATwoThousand,
		"memoryoptimized_m500":  SkuNameMemoryOptimizedMFiveHundred,
		"memoryoptimized_m50":   SkuNameMemoryOptimizedMFiveZero,
		"memoryoptimized_m1500": SkuNameMemoryOptimizedMOneFiveHundred,
		"memoryoptimized_m150":  SkuNameMemoryOptimizedMOneFiveZero,
		"memoryoptimized_m100":  SkuNameMemoryOptimizedMOneHundred,
		"memoryoptimized_m1000": SkuNameMemoryOptimizedMOneThousand,
		"memoryoptimized_m10":   SkuNameMemoryOptimizedMOneZero,
		"memoryoptimized_m700":  SkuNameMemoryOptimizedMSevenHundred,
		"memoryoptimized_m350":  SkuNameMemoryOptimizedMThreeFiveZero,
		"memoryoptimized_m250":  SkuNameMemoryOptimizedMTwoFiveZero,
		"memoryoptimized_m2000": SkuNameMemoryOptimizedMTwoThousand,
		"memoryoptimized_m20":   SkuNameMemoryOptimizedMTwoZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type TlsVersion string

const (
	TlsVersionOnePointOne  TlsVersion = "1.1"
	TlsVersionOnePointTwo  TlsVersion = "1.2"
	TlsVersionOnePointZero TlsVersion = "1.0"
)

func PossibleValuesForTlsVersion() []string {
	return []string{
		string(TlsVersionOnePointOne),
		string(TlsVersionOnePointTwo),
		string(TlsVersionOnePointZero),
	}
}

func (s *TlsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTlsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTlsVersion(input string) (*TlsVersion, error) {
	vals := map[string]TlsVersion{
		"1.1": TlsVersionOnePointOne,
		"1.2": TlsVersionOnePointTwo,
		"1.0": TlsVersionOnePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TlsVersion(input)
	return &out, nil
}
