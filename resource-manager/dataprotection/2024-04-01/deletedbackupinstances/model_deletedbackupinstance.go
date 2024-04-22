package deletedbackupinstances

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedBackupInstance struct {
	CurrentProtectionState         *CurrentProtectionState  `json:"currentProtectionState,omitempty"`
	DataSourceInfo                 Datasource               `json:"dataSourceInfo"`
	DataSourceSetInfo              *DatasourceSet           `json:"dataSourceSetInfo,omitempty"`
	DatasourceAuthCredentials      AuthCredentials          `json:"datasourceAuthCredentials"`
	DeletionInfo                   *DeletionInfo            `json:"deletionInfo,omitempty"`
	FriendlyName                   *string                  `json:"friendlyName,omitempty"`
	IdentityDetails                *IdentityDetails         `json:"identityDetails,omitempty"`
	ObjectType                     string                   `json:"objectType"`
	PolicyInfo                     PolicyInfo               `json:"policyInfo"`
	ProtectionErrorDetails         *UserFacingError         `json:"protectionErrorDetails,omitempty"`
	ProtectionStatus               *ProtectionStatusDetails `json:"protectionStatus,omitempty"`
	ProvisioningState              *string                  `json:"provisioningState,omitempty"`
	ResourceGuardOperationRequests *[]string                `json:"resourceGuardOperationRequests,omitempty"`
	ValidationType                 *ValidationType          `json:"validationType,omitempty"`
}

var _ json.Unmarshaler = &DeletedBackupInstance{}

func (s *DeletedBackupInstance) UnmarshalJSON(bytes []byte) error {
	type alias DeletedBackupInstance
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DeletedBackupInstance: %+v", err)
	}

	s.CurrentProtectionState = decoded.CurrentProtectionState
	s.DataSourceInfo = decoded.DataSourceInfo
	s.DataSourceSetInfo = decoded.DataSourceSetInfo
	s.DeletionInfo = decoded.DeletionInfo
	s.FriendlyName = decoded.FriendlyName
	s.IdentityDetails = decoded.IdentityDetails
	s.ObjectType = decoded.ObjectType
	s.PolicyInfo = decoded.PolicyInfo
	s.ProtectionErrorDetails = decoded.ProtectionErrorDetails
	s.ProtectionStatus = decoded.ProtectionStatus
	s.ProvisioningState = decoded.ProvisioningState
	s.ResourceGuardOperationRequests = decoded.ResourceGuardOperationRequests
	s.ValidationType = decoded.ValidationType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeletedBackupInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["datasourceAuthCredentials"]; ok {
		impl, err := unmarshalAuthCredentialsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DatasourceAuthCredentials' for 'DeletedBackupInstance': %+v", err)
		}
		s.DatasourceAuthCredentials = impl
	}
	return nil
}
