package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForDatabasesGcpOffering{}

type DefenderForDatabasesGcpOffering struct {
	ArcAutoProvisioning                     *DefenderForDatabasesGcpOfferingArcAutoProvisioning                     `json:"arcAutoProvisioning,omitempty"`
	DefenderForDatabasesArcAutoProvisioning *DefenderForDatabasesGcpOfferingDefenderForDatabasesArcAutoProvisioning `json:"defenderForDatabasesArcAutoProvisioning,omitempty"`

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = DefenderForDatabasesGcpOffering{}

func (s DefenderForDatabasesGcpOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForDatabasesGcpOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForDatabasesGcpOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForDatabasesGcpOffering: %+v", err)
	}
	decoded["offeringType"] = "DefenderForDatabasesGcp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForDatabasesGcpOffering: %+v", err)
	}

	return encoded, nil
}
