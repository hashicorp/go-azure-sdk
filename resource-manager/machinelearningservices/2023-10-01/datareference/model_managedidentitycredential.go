package datareference

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataReferenceCredential = ManagedIdentityCredential{}

type ManagedIdentityCredential struct {
	ManagedIdentityType            *string `json:"managedIdentityType,omitempty"`
	UserManagedIdentityClientId    *string `json:"userManagedIdentityClientId,omitempty"`
	UserManagedIdentityPrincipalId *string `json:"userManagedIdentityPrincipalId,omitempty"`
	UserManagedIdentityResourceId  *string `json:"userManagedIdentityResourceId,omitempty"`
	UserManagedIdentityTenantId    *string `json:"userManagedIdentityTenantId,omitempty"`

	// Fields inherited from DataReferenceCredential
}

var _ json.Marshaler = ManagedIdentityCredential{}

func (s ManagedIdentityCredential) MarshalJSON() ([]byte, error) {
	type wrapper ManagedIdentityCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedIdentityCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedIdentityCredential: %+v", err)
	}
	decoded["credentialType"] = "ManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedIdentityCredential: %+v", err)
	}

	return encoded, nil
}
