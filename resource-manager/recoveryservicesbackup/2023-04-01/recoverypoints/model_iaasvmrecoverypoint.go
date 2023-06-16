package recoverypoints

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = IaasVMRecoveryPoint{}

type IaasVMRecoveryPoint struct {
	IsInstantIlrSessionActive       *bool                                      `json:"isInstantIlrSessionActive,omitempty"`
	IsManagedVirtualMachine         *bool                                      `json:"isManagedVirtualMachine,omitempty"`
	IsPrivateAccessEnabledOnAnyDisk *bool                                      `json:"isPrivateAccessEnabledOnAnyDisk,omitempty"`
	IsSourceVMEncrypted             *bool                                      `json:"isSourceVMEncrypted,omitempty"`
	KeyAndSecret                    *KeyAndSecretDetails                       `json:"keyAndSecret,omitempty"`
	OriginalStorageAccountOption    *bool                                      `json:"originalStorageAccountOption,omitempty"`
	OsType                          *string                                    `json:"osType,omitempty"`
	RecoveryPointAdditionalInfo     *string                                    `json:"recoveryPointAdditionalInfo,omitempty"`
	RecoveryPointDiskConfiguration  *RecoveryPointDiskConfiguration            `json:"recoveryPointDiskConfiguration,omitempty"`
	RecoveryPointMoveReadinessInfo  *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties         *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails        *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTime               *string                                    `json:"recoveryPointTime,omitempty"`
	RecoveryPointType               *string                                    `json:"recoveryPointType,omitempty"`
	SecurityType                    *string                                    `json:"securityType,omitempty"`
	SourceVMStorageType             *string                                    `json:"sourceVMStorageType,omitempty"`
	VirtualMachineSize              *string                                    `json:"virtualMachineSize,omitempty"`
	Zones                           *zones.Schema                              `json:"zones,omitempty"`

	// Fields inherited from RecoveryPoint
}

var _ json.Marshaler = IaasVMRecoveryPoint{}

func (s IaasVMRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper IaasVMRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IaasVMRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IaasVMRecoveryPoint: %+v", err)
	}
	decoded["objectType"] = "IaasVMRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IaasVMRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
