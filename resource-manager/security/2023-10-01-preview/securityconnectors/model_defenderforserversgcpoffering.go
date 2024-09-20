package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForServersGcpOffering{}

type DefenderForServersGcpOffering struct {
	ArcAutoProvisioning *DefenderForServersGcpOfferingArcAutoProvisioning `json:"arcAutoProvisioning,omitempty"`
	DefenderForServers  *DefenderForServersGcpOfferingDefenderForServers  `json:"defenderForServers,omitempty"`
	MdeAutoProvisioning *DefenderForServersGcpOfferingMdeAutoProvisioning `json:"mdeAutoProvisioning,omitempty"`
	SubPlan             *DefenderForServersGcpOfferingSubPlan             `json:"subPlan,omitempty"`
	VMScanners          *DefenderForServersGcpOfferingVMScanners          `json:"vmScanners,omitempty"`
	VaAutoProvisioning  *DefenderForServersGcpOfferingVaAutoProvisioning  `json:"vaAutoProvisioning,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s DefenderForServersGcpOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = DefenderForServersGcpOffering{}

func (s DefenderForServersGcpOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForServersGcpOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForServersGcpOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForServersGcpOffering: %+v", err)
	}

	decoded["offeringType"] = "DefenderForServersGcp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForServersGcpOffering: %+v", err)
	}

	return encoded, nil
}
