package signalr

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ACLAction string

const (
	ACLActionAllow ACLAction = "Allow"
	ACLActionDeny  ACLAction = "Deny"
)

func PossibleValuesForACLAction() []string {
	return []string{
		string(ACLActionAllow),
		string(ACLActionDeny),
	}
}

func (s *ACLAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForACLAction() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ACLAction(decoded)
	return nil
}

type FeatureFlags string

const (
	FeatureFlagsEnableConnectivityLogs FeatureFlags = "EnableConnectivityLogs"
	FeatureFlagsEnableLiveTrace        FeatureFlags = "EnableLiveTrace"
	FeatureFlagsEnableMessagingLogs    FeatureFlags = "EnableMessagingLogs"
	FeatureFlagsServiceMode            FeatureFlags = "ServiceMode"
)

func PossibleValuesForFeatureFlags() []string {
	return []string{
		string(FeatureFlagsEnableConnectivityLogs),
		string(FeatureFlagsEnableLiveTrace),
		string(FeatureFlagsEnableMessagingLogs),
		string(FeatureFlagsServiceMode),
	}
}

func (s *FeatureFlags) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFeatureFlags() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FeatureFlags(decoded)
	return nil
}

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSalt      KeyType = "Salt"
	KeyTypeSecondary KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSalt),
		string(KeyTypeSecondary),
	}
}

func (s *KeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForKeyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = KeyType(decoded)
	return nil
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}

func (s *PrivateLinkServiceConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPrivateLinkServiceConnectionStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PrivateLinkServiceConnectionStatus(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type ScaleType string

const (
	ScaleTypeAutomatic ScaleType = "Automatic"
	ScaleTypeManual    ScaleType = "Manual"
	ScaleTypeNone      ScaleType = "None"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeAutomatic),
		string(ScaleTypeManual),
		string(ScaleTypeNone),
	}
}

func (s *ScaleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForScaleType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ScaleType(decoded)
	return nil
}

type ServiceKind string

const (
	ServiceKindRawWebSockets ServiceKind = "RawWebSockets"
	ServiceKindSignalR       ServiceKind = "SignalR"
)

func PossibleValuesForServiceKind() []string {
	return []string{
		string(ServiceKindRawWebSockets),
		string(ServiceKindSignalR),
	}
}

func (s *ServiceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForServiceKind() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ServiceKind(decoded)
	return nil
}

type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = "Timeout"
)

func PossibleValuesForSharedPrivateLinkResourceStatus() []string {
	return []string{
		string(SharedPrivateLinkResourceStatusApproved),
		string(SharedPrivateLinkResourceStatusDisconnected),
		string(SharedPrivateLinkResourceStatusPending),
		string(SharedPrivateLinkResourceStatusRejected),
		string(SharedPrivateLinkResourceStatusTimeout),
	}
}

func (s *SharedPrivateLinkResourceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSharedPrivateLinkResourceStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SharedPrivateLinkResourceStatus(decoded)
	return nil
}

type SignalRRequestType string

const (
	SignalRRequestTypeClientConnection SignalRRequestType = "ClientConnection"
	SignalRRequestTypeRESTAPI          SignalRRequestType = "RESTAPI"
	SignalRRequestTypeServerConnection SignalRRequestType = "ServerConnection"
	SignalRRequestTypeTrace            SignalRRequestType = "Trace"
)

func PossibleValuesForSignalRRequestType() []string {
	return []string{
		string(SignalRRequestTypeClientConnection),
		string(SignalRRequestTypeRESTAPI),
		string(SignalRRequestTypeServerConnection),
		string(SignalRRequestTypeTrace),
	}
}

func (s *SignalRRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSignalRRequestType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SignalRRequestType(decoded)
	return nil
}

type SignalRSkuTier string

const (
	SignalRSkuTierBasic    SignalRSkuTier = "Basic"
	SignalRSkuTierFree     SignalRSkuTier = "Free"
	SignalRSkuTierPremium  SignalRSkuTier = "Premium"
	SignalRSkuTierStandard SignalRSkuTier = "Standard"
)

func PossibleValuesForSignalRSkuTier() []string {
	return []string{
		string(SignalRSkuTierBasic),
		string(SignalRSkuTierFree),
		string(SignalRSkuTierPremium),
		string(SignalRSkuTierStandard),
	}
}

func (s *SignalRSkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSignalRSkuTier() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SignalRSkuTier(decoded)
	return nil
}

type UpstreamAuthType string

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	UpstreamAuthTypeNone            UpstreamAuthType = "None"
)

func PossibleValuesForUpstreamAuthType() []string {
	return []string{
		string(UpstreamAuthTypeManagedIdentity),
		string(UpstreamAuthTypeNone),
	}
}

func (s *UpstreamAuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForUpstreamAuthType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = UpstreamAuthType(decoded)
	return nil
}
