package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderCspmAwsOffering{}

type DefenderCspmAwsOffering struct {
	Ciem                               *DefenderCspmAwsOfferingCiem                               `json:"ciem,omitempty"`
	DataSensitivityDiscovery           *DefenderCspmAwsOfferingDataSensitivityDiscovery           `json:"dataSensitivityDiscovery,omitempty"`
	DatabasesDspm                      *DefenderCspmAwsOfferingDatabasesDspm                      `json:"databasesDspm,omitempty"`
	MdcContainersAgentlessDiscoveryK8s *DefenderCspmAwsOfferingMdcContainersAgentlessDiscoveryK8s `json:"mdcContainersAgentlessDiscoveryK8s,omitempty"`
	MdcContainersImageAssessment       *DefenderCspmAwsOfferingMdcContainersImageAssessment       `json:"mdcContainersImageAssessment,omitempty"`
	VMScanners                         *DefenderCspmAwsOfferingVMScanners                         `json:"vmScanners,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s DefenderCspmAwsOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = DefenderCspmAwsOffering{}

func (s DefenderCspmAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderCspmAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderCspmAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderCspmAwsOffering: %+v", err)
	}

	decoded["offeringType"] = "DefenderCspmAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderCspmAwsOffering: %+v", err)
	}

	return encoded, nil
}
