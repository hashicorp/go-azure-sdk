package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderFoDatabasesAwsOffering{}

type DefenderFoDatabasesAwsOffering struct {
	ArcAutoProvisioning *DefenderFoDatabasesAwsOfferingArcAutoProvisioning `json:"arcAutoProvisioning,omitempty"`
	DatabasesDspm       *DefenderFoDatabasesAwsOfferingDatabasesDspm       `json:"databasesDspm,omitempty"`
	Rds                 *DefenderFoDatabasesAwsOfferingRds                 `json:"rds,omitempty"`

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = DefenderFoDatabasesAwsOffering{}

func (s DefenderFoDatabasesAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderFoDatabasesAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderFoDatabasesAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderFoDatabasesAwsOffering: %+v", err)
	}
	decoded["offeringType"] = "DefenderForDatabasesAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderFoDatabasesAwsOffering: %+v", err)
	}

	return encoded, nil
}
