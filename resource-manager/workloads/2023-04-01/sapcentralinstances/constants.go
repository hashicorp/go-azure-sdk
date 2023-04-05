package sapcentralinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CentralServerVirtualMachineType string

const (
	CentralServerVirtualMachineTypeASCS        CentralServerVirtualMachineType = "ASCS"
	CentralServerVirtualMachineTypeERS         CentralServerVirtualMachineType = "ERS"
	CentralServerVirtualMachineTypeERSInactive CentralServerVirtualMachineType = "ERSInactive"
	CentralServerVirtualMachineTypePrimary     CentralServerVirtualMachineType = "Primary"
	CentralServerVirtualMachineTypeSecondary   CentralServerVirtualMachineType = "Secondary"
	CentralServerVirtualMachineTypeStandby     CentralServerVirtualMachineType = "Standby"
	CentralServerVirtualMachineTypeUnknown     CentralServerVirtualMachineType = "Unknown"
)

func PossibleValuesForCentralServerVirtualMachineType() []string {
	return []string{
		string(CentralServerVirtualMachineTypeASCS),
		string(CentralServerVirtualMachineTypeERS),
		string(CentralServerVirtualMachineTypeERSInactive),
		string(CentralServerVirtualMachineTypePrimary),
		string(CentralServerVirtualMachineTypeSecondary),
		string(CentralServerVirtualMachineTypeStandby),
		string(CentralServerVirtualMachineTypeUnknown),
	}
}

func (s *CentralServerVirtualMachineType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForCentralServerVirtualMachineType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = CentralServerVirtualMachineType(decoded)
	return nil
}

type EnqueueReplicationServerType string

const (
	EnqueueReplicationServerTypeEnqueueReplicatorOne EnqueueReplicationServerType = "EnqueueReplicator1"
	EnqueueReplicationServerTypeEnqueueReplicatorTwo EnqueueReplicationServerType = "EnqueueReplicator2"
)

func PossibleValuesForEnqueueReplicationServerType() []string {
	return []string{
		string(EnqueueReplicationServerTypeEnqueueReplicatorOne),
		string(EnqueueReplicationServerTypeEnqueueReplicatorTwo),
	}
}

func (s *EnqueueReplicationServerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEnqueueReplicationServerType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EnqueueReplicationServerType(decoded)
	return nil
}

type SAPHealthState string

const (
	SAPHealthStateDegraded  SAPHealthState = "Degraded"
	SAPHealthStateHealthy   SAPHealthState = "Healthy"
	SAPHealthStateUnhealthy SAPHealthState = "Unhealthy"
	SAPHealthStateUnknown   SAPHealthState = "Unknown"
)

func PossibleValuesForSAPHealthState() []string {
	return []string{
		string(SAPHealthStateDegraded),
		string(SAPHealthStateHealthy),
		string(SAPHealthStateUnhealthy),
		string(SAPHealthStateUnknown),
	}
}

func (s *SAPHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPHealthState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPHealthState(decoded)
	return nil
}

type SAPVirtualInstanceStatus string

const (
	SAPVirtualInstanceStatusOffline          SAPVirtualInstanceStatus = "Offline"
	SAPVirtualInstanceStatusPartiallyRunning SAPVirtualInstanceStatus = "PartiallyRunning"
	SAPVirtualInstanceStatusRunning          SAPVirtualInstanceStatus = "Running"
	SAPVirtualInstanceStatusSoftShutdown     SAPVirtualInstanceStatus = "SoftShutdown"
	SAPVirtualInstanceStatusStarting         SAPVirtualInstanceStatus = "Starting"
	SAPVirtualInstanceStatusStopping         SAPVirtualInstanceStatus = "Stopping"
	SAPVirtualInstanceStatusUnavailable      SAPVirtualInstanceStatus = "Unavailable"
)

func PossibleValuesForSAPVirtualInstanceStatus() []string {
	return []string{
		string(SAPVirtualInstanceStatusOffline),
		string(SAPVirtualInstanceStatusPartiallyRunning),
		string(SAPVirtualInstanceStatusRunning),
		string(SAPVirtualInstanceStatusSoftShutdown),
		string(SAPVirtualInstanceStatusStarting),
		string(SAPVirtualInstanceStatusStopping),
		string(SAPVirtualInstanceStatusUnavailable),
	}
}

func (s *SAPVirtualInstanceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPVirtualInstanceStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPVirtualInstanceStatus(decoded)
	return nil
}

type SapVirtualInstanceProvisioningState string

const (
	SapVirtualInstanceProvisioningStateCreating  SapVirtualInstanceProvisioningState = "Creating"
	SapVirtualInstanceProvisioningStateDeleting  SapVirtualInstanceProvisioningState = "Deleting"
	SapVirtualInstanceProvisioningStateFailed    SapVirtualInstanceProvisioningState = "Failed"
	SapVirtualInstanceProvisioningStateSucceeded SapVirtualInstanceProvisioningState = "Succeeded"
	SapVirtualInstanceProvisioningStateUpdating  SapVirtualInstanceProvisioningState = "Updating"
)

func PossibleValuesForSapVirtualInstanceProvisioningState() []string {
	return []string{
		string(SapVirtualInstanceProvisioningStateCreating),
		string(SapVirtualInstanceProvisioningStateDeleting),
		string(SapVirtualInstanceProvisioningStateFailed),
		string(SapVirtualInstanceProvisioningStateSucceeded),
		string(SapVirtualInstanceProvisioningStateUpdating),
	}
}

func (s *SapVirtualInstanceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSapVirtualInstanceProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SapVirtualInstanceProvisioningState(decoded)
	return nil
}
