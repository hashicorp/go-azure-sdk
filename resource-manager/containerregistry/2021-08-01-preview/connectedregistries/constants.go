package connectedregistries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivationStatus string

const (
	ActivationStatusActive   ActivationStatus = "Active"
	ActivationStatusInactive ActivationStatus = "Inactive"
)

func PossibleValuesForActivationStatus() []string {
	return []string{
		string(ActivationStatusActive),
		string(ActivationStatusInactive),
	}
}

func (s *ActivationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForActivationStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ActivationStatus(decoded)
	return nil
}

type AuditLogStatus string

const (
	AuditLogStatusDisabled AuditLogStatus = "Disabled"
	AuditLogStatusEnabled  AuditLogStatus = "Enabled"
)

func PossibleValuesForAuditLogStatus() []string {
	return []string{
		string(AuditLogStatusDisabled),
		string(AuditLogStatusEnabled),
	}
}

func (s *AuditLogStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAuditLogStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AuditLogStatus(decoded)
	return nil
}

type CertificateType string

const (
	CertificateTypeLocalDirectory CertificateType = "LocalDirectory"
)

func PossibleValuesForCertificateType() []string {
	return []string{
		string(CertificateTypeLocalDirectory),
	}
}

func (s *CertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForCertificateType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = CertificateType(decoded)
	return nil
}

type ConnectedRegistryMode string

const (
	ConnectedRegistryModeMirror    ConnectedRegistryMode = "Mirror"
	ConnectedRegistryModeReadOnly  ConnectedRegistryMode = "ReadOnly"
	ConnectedRegistryModeReadWrite ConnectedRegistryMode = "ReadWrite"
	ConnectedRegistryModeRegistry  ConnectedRegistryMode = "Registry"
)

func PossibleValuesForConnectedRegistryMode() []string {
	return []string{
		string(ConnectedRegistryModeMirror),
		string(ConnectedRegistryModeReadOnly),
		string(ConnectedRegistryModeReadWrite),
		string(ConnectedRegistryModeRegistry),
	}
}

func (s *ConnectedRegistryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForConnectedRegistryMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ConnectedRegistryMode(decoded)
	return nil
}

type ConnectionState string

const (
	ConnectionStateOffline   ConnectionState = "Offline"
	ConnectionStateOnline    ConnectionState = "Online"
	ConnectionStateSyncing   ConnectionState = "Syncing"
	ConnectionStateUnhealthy ConnectionState = "Unhealthy"
)

func PossibleValuesForConnectionState() []string {
	return []string{
		string(ConnectionStateOffline),
		string(ConnectionStateOnline),
		string(ConnectionStateSyncing),
		string(ConnectionStateUnhealthy),
	}
}

func (s *ConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForConnectionState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ConnectionState(decoded)
	return nil
}

type LogLevel string

const (
	LogLevelDebug       LogLevel = "Debug"
	LogLevelError       LogLevel = "Error"
	LogLevelInformation LogLevel = "Information"
	LogLevelNone        LogLevel = "None"
	LogLevelWarning     LogLevel = "Warning"
)

func PossibleValuesForLogLevel() []string {
	return []string{
		string(LogLevelDebug),
		string(LogLevelError),
		string(LogLevelInformation),
		string(LogLevelNone),
		string(LogLevelWarning),
	}
}

func (s *LogLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLogLevel() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = LogLevel(decoded)
	return nil
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
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type TlsStatus string

const (
	TlsStatusDisabled TlsStatus = "Disabled"
	TlsStatusEnabled  TlsStatus = "Enabled"
)

func PossibleValuesForTlsStatus() []string {
	return []string{
		string(TlsStatusDisabled),
		string(TlsStatusEnabled),
	}
}

func (s *TlsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTlsStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TlsStatus(decoded)
	return nil
}
