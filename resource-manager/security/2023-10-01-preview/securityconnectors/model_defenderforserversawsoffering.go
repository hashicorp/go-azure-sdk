package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForServersAwsOffering{}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `json:"arcAutoProvisioning,omitempty"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `json:"defenderForServers,omitempty"`
	MdeAutoProvisioning *DefenderForServersAwsOfferingMdeAutoProvisioning `json:"mdeAutoProvisioning,omitempty"`
	SubPlan             *DefenderForServersAwsOfferingSubPlan             `json:"subPlan,omitempty"`
	VMScanners          *DefenderForServersAwsOfferingVMScanners          `json:"vmScanners,omitempty"`
	VaAutoProvisioning  *DefenderForServersAwsOfferingVaAutoProvisioning  `json:"vaAutoProvisioning,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s DefenderForServersAwsOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = DefenderForServersAwsOffering{}

func (s DefenderForServersAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForServersAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForServersAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForServersAwsOffering: %+v", err)
	}

	decoded["offeringType"] = "DefenderForServersAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForServersAwsOffering: %+v", err)
	}

	return encoded, nil
}
