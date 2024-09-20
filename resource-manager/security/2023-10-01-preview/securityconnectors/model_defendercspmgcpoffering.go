package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderCspmGcpOffering{}

type DefenderCspmGcpOffering struct {
	CiemDiscovery                      *DefenderCspmGcpOfferingCiemDiscovery                      `json:"ciemDiscovery,omitempty"`
	DataSensitivityDiscovery           *DefenderCspmGcpOfferingDataSensitivityDiscovery           `json:"dataSensitivityDiscovery,omitempty"`
	MdcContainersAgentlessDiscoveryK8s *DefenderCspmGcpOfferingMdcContainersAgentlessDiscoveryK8s `json:"mdcContainersAgentlessDiscoveryK8s,omitempty"`
	MdcContainersImageAssessment       *DefenderCspmGcpOfferingMdcContainersImageAssessment       `json:"mdcContainersImageAssessment,omitempty"`
	VMScanners                         *DefenderCspmGcpOfferingVMScanners                         `json:"vmScanners,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s DefenderCspmGcpOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = DefenderCspmGcpOffering{}

func (s DefenderCspmGcpOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderCspmGcpOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderCspmGcpOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderCspmGcpOffering: %+v", err)
	}

	decoded["offeringType"] = "DefenderCspmGcp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderCspmGcpOffering: %+v", err)
	}

	return encoded, nil
}
