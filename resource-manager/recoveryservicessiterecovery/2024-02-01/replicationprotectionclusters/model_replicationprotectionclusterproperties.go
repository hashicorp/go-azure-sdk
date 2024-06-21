package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionClusterProperties struct {
	ActiveLocation                          *string                                    `json:"activeLocation,omitempty"`
	AgentClusterId                          *string                                    `json:"agentClusterId,omitempty"`
	AllowedOperations                       *[]string                                  `json:"allowedOperations,omitempty"`
	AreAllClusterNodesRegistered            *bool                                      `json:"areAllClusterNodesRegistered,omitempty"`
	ClusterFqdn                             *string                                    `json:"clusterFqdn,omitempty"`
	ClusterNodeFqdns                        *[]string                                  `json:"clusterNodeFqdns,omitempty"`
	ClusterProtectedItemIds                 *[]string                                  `json:"clusterProtectedItemIds,omitempty"`
	ClusterRegisteredNodes                  *[]RegisteredClusterNodes                  `json:"clusterRegisteredNodes,omitempty"`
	CurrentScenario                         *CurrentScenarioDetails                    `json:"currentScenario,omitempty"`
	HealthErrors                            *[]HealthError                             `json:"healthErrors,omitempty"`
	LastSuccessfulFailoverTime              *string                                    `json:"lastSuccessfulFailoverTime,omitempty"`
	LastSuccessfulTestFailoverTime          *string                                    `json:"lastSuccessfulTestFailoverTime,omitempty"`
	PolicyFriendlyName                      *string                                    `json:"policyFriendlyName,omitempty"`
	PolicyId                                *string                                    `json:"policyId,omitempty"`
	PrimaryFabricFriendlyName               *string                                    `json:"primaryFabricFriendlyName,omitempty"`
	PrimaryFabricProvider                   *string                                    `json:"primaryFabricProvider,omitempty"`
	PrimaryProtectionContainerFriendlyName  *string                                    `json:"primaryProtectionContainerFriendlyName,omitempty"`
	ProtectionClusterType                   *string                                    `json:"protectionClusterType,omitempty"`
	ProtectionState                         *string                                    `json:"protectionState,omitempty"`
	ProtectionStateDescription              *string                                    `json:"protectionStateDescription,omitempty"`
	ProviderSpecificDetails                 ReplicationClusterProviderSpecificSettings `json:"providerSpecificDetails"`
	ProvisioningState                       *string                                    `json:"provisioningState,omitempty"`
	RecoveryContainerId                     *string                                    `json:"recoveryContainerId,omitempty"`
	RecoveryFabricFriendlyName              *string                                    `json:"recoveryFabricFriendlyName,omitempty"`
	RecoveryFabricId                        *string                                    `json:"recoveryFabricId,omitempty"`
	RecoveryProtectionContainerFriendlyName *string                                    `json:"recoveryProtectionContainerFriendlyName,omitempty"`
	ReplicationHealth                       *string                                    `json:"replicationHealth,omitempty"`
	SharedDiskProperties                    *SharedDiskReplicationItemProperties       `json:"sharedDiskProperties,omitempty"`
	TestFailoverState                       *string                                    `json:"testFailoverState,omitempty"`
	TestFailoverStateDescription            *string                                    `json:"testFailoverStateDescription,omitempty"`
}

var _ json.Unmarshaler = &ReplicationProtectionClusterProperties{}

func (s *ReplicationProtectionClusterProperties) UnmarshalJSON(bytes []byte) error {
	type alias ReplicationProtectionClusterProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ReplicationProtectionClusterProperties: %+v", err)
	}

	s.ActiveLocation = decoded.ActiveLocation
	s.AgentClusterId = decoded.AgentClusterId
	s.AllowedOperations = decoded.AllowedOperations
	s.AreAllClusterNodesRegistered = decoded.AreAllClusterNodesRegistered
	s.ClusterFqdn = decoded.ClusterFqdn
	s.ClusterNodeFqdns = decoded.ClusterNodeFqdns
	s.ClusterProtectedItemIds = decoded.ClusterProtectedItemIds
	s.ClusterRegisteredNodes = decoded.ClusterRegisteredNodes
	s.CurrentScenario = decoded.CurrentScenario
	s.HealthErrors = decoded.HealthErrors
	s.LastSuccessfulFailoverTime = decoded.LastSuccessfulFailoverTime
	s.LastSuccessfulTestFailoverTime = decoded.LastSuccessfulTestFailoverTime
	s.PolicyFriendlyName = decoded.PolicyFriendlyName
	s.PolicyId = decoded.PolicyId
	s.PrimaryFabricFriendlyName = decoded.PrimaryFabricFriendlyName
	s.PrimaryFabricProvider = decoded.PrimaryFabricProvider
	s.PrimaryProtectionContainerFriendlyName = decoded.PrimaryProtectionContainerFriendlyName
	s.ProtectionClusterType = decoded.ProtectionClusterType
	s.ProtectionState = decoded.ProtectionState
	s.ProtectionStateDescription = decoded.ProtectionStateDescription
	s.ProvisioningState = decoded.ProvisioningState
	s.RecoveryContainerId = decoded.RecoveryContainerId
	s.RecoveryFabricFriendlyName = decoded.RecoveryFabricFriendlyName
	s.RecoveryFabricId = decoded.RecoveryFabricId
	s.RecoveryProtectionContainerFriendlyName = decoded.RecoveryProtectionContainerFriendlyName
	s.ReplicationHealth = decoded.ReplicationHealth
	s.SharedDiskProperties = decoded.SharedDiskProperties
	s.TestFailoverState = decoded.TestFailoverState
	s.TestFailoverStateDescription = decoded.TestFailoverStateDescription

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ReplicationProtectionClusterProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalReplicationClusterProviderSpecificSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ReplicationProtectionClusterProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
