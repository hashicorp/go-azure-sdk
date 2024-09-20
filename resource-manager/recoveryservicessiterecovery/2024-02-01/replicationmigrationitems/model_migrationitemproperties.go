package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationItemProperties struct {
	AllowedOperations           *[]MigrationItemOperation         `json:"allowedOperations,omitempty"`
	CriticalJobHistory          *[]CriticalJobHistoryDetails      `json:"criticalJobHistory,omitempty"`
	CurrentJob                  *CurrentJobDetails                `json:"currentJob,omitempty"`
	EventCorrelationId          *string                           `json:"eventCorrelationId,omitempty"`
	Health                      *ProtectionHealth                 `json:"health,omitempty"`
	HealthErrors                *[]HealthError                    `json:"healthErrors,omitempty"`
	LastMigrationStatus         *string                           `json:"lastMigrationStatus,omitempty"`
	LastMigrationTime           *string                           `json:"lastMigrationTime,omitempty"`
	LastTestMigrationStatus     *string                           `json:"lastTestMigrationStatus,omitempty"`
	LastTestMigrationTime       *string                           `json:"lastTestMigrationTime,omitempty"`
	MachineName                 *string                           `json:"machineName,omitempty"`
	MigrationState              *MigrationState                   `json:"migrationState,omitempty"`
	MigrationStateDescription   *string                           `json:"migrationStateDescription,omitempty"`
	PolicyFriendlyName          *string                           `json:"policyFriendlyName,omitempty"`
	PolicyId                    *string                           `json:"policyId,omitempty"`
	ProviderSpecificDetails     MigrationProviderSpecificSettings `json:"providerSpecificDetails"`
	RecoveryServicesProviderId  *string                           `json:"recoveryServicesProviderId,omitempty"`
	ReplicationStatus           *string                           `json:"replicationStatus,omitempty"`
	TestMigrateState            *TestMigrationState               `json:"testMigrateState,omitempty"`
	TestMigrateStateDescription *string                           `json:"testMigrateStateDescription,omitempty"`
}

func (o *MigrationItemProperties) GetLastMigrationTimeAsTime() (*time.Time, error) {
	if o.LastMigrationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastMigrationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationItemProperties) SetLastMigrationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastMigrationTime = &formatted
}

func (o *MigrationItemProperties) GetLastTestMigrationTimeAsTime() (*time.Time, error) {
	if o.LastTestMigrationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastTestMigrationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationItemProperties) SetLastTestMigrationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastTestMigrationTime = &formatted
}

var _ json.Unmarshaler = &MigrationItemProperties{}

func (s *MigrationItemProperties) UnmarshalJSON(bytes []byte) error {
	type alias MigrationItemProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MigrationItemProperties: %+v", err)
	}

	s.AllowedOperations = decoded.AllowedOperations
	s.CriticalJobHistory = decoded.CriticalJobHistory
	s.CurrentJob = decoded.CurrentJob
	s.EventCorrelationId = decoded.EventCorrelationId
	s.Health = decoded.Health
	s.HealthErrors = decoded.HealthErrors
	s.LastMigrationStatus = decoded.LastMigrationStatus
	s.LastMigrationTime = decoded.LastMigrationTime
	s.LastTestMigrationStatus = decoded.LastTestMigrationStatus
	s.LastTestMigrationTime = decoded.LastTestMigrationTime
	s.MachineName = decoded.MachineName
	s.MigrationState = decoded.MigrationState
	s.MigrationStateDescription = decoded.MigrationStateDescription
	s.PolicyFriendlyName = decoded.PolicyFriendlyName
	s.PolicyId = decoded.PolicyId
	s.RecoveryServicesProviderId = decoded.RecoveryServicesProviderId
	s.ReplicationStatus = decoded.ReplicationStatus
	s.TestMigrateState = decoded.TestMigrateState
	s.TestMigrateStateDescription = decoded.TestMigrateStateDescription

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MigrationItemProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := UnmarshalMigrationProviderSpecificSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'MigrationItemProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
