package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForContainersGcpOffering{}

type DefenderForContainersGcpOffering struct {
	AuditLogsAutoProvisioningFlag      *bool                                                               `json:"auditLogsAutoProvisioningFlag,omitempty"`
	DataPipelineNativeCloudConnection  *DefenderForContainersGcpOfferingDataPipelineNativeCloudConnection  `json:"dataPipelineNativeCloudConnection,omitempty"`
	DefenderAgentAutoProvisioningFlag  *bool                                                               `json:"defenderAgentAutoProvisioningFlag,omitempty"`
	MdcContainersAgentlessDiscoveryK8s *DefenderForContainersGcpOfferingMdcContainersAgentlessDiscoveryK8s `json:"mdcContainersAgentlessDiscoveryK8s,omitempty"`
	MdcContainersImageAssessment       *DefenderForContainersGcpOfferingMdcContainersImageAssessment       `json:"mdcContainersImageAssessment,omitempty"`
	NativeCloudConnection              *DefenderForContainersGcpOfferingNativeCloudConnection              `json:"nativeCloudConnection,omitempty"`
	PolicyAgentAutoProvisioningFlag    *bool                                                               `json:"policyAgentAutoProvisioningFlag,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s DefenderForContainersGcpOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = DefenderForContainersGcpOffering{}

func (s DefenderForContainersGcpOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForContainersGcpOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForContainersGcpOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForContainersGcpOffering: %+v", err)
	}

	decoded["offeringType"] = "DefenderForContainersGcp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForContainersGcpOffering: %+v", err)
	}

	return encoded, nil
}
