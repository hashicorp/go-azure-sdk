package vaultcertificates

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceCertificateDetails = ResourceCertificateAndAadDetails{}

type ResourceCertificateAndAadDetails struct {
	AadAudience                     *string `json:"aadAudience,omitempty"`
	AadAuthority                    string  `json:"aadAuthority"`
	AadTenantId                     string  `json:"aadTenantId"`
	AzureManagementEndpointAudience string  `json:"azureManagementEndpointAudience"`
	ServicePrincipalClientId        string  `json:"servicePrincipalClientId"`
	ServicePrincipalObjectId        string  `json:"servicePrincipalObjectId"`
	ServiceResourceId               *string `json:"serviceResourceId,omitempty"`

	// Fields inherited from ResourceCertificateDetails
	Certificate  *string `json:"certificate,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Issuer       *string `json:"issuer,omitempty"`
	ResourceId   *int64  `json:"resourceId,omitempty"`
	Subject      *string `json:"subject,omitempty"`
	Thumbprint   *string `json:"thumbprint,omitempty"`
	ValidFrom    *string `json:"validFrom,omitempty"`
	ValidTo      *string `json:"validTo,omitempty"`
}

var _ json.Marshaler = ResourceCertificateAndAadDetails{}

func (s ResourceCertificateAndAadDetails) MarshalJSON() ([]byte, error) {
	type wrapper ResourceCertificateAndAadDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceCertificateAndAadDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceCertificateAndAadDetails: %+v", err)
	}
	decoded["authType"] = "AzureActiveDirectory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceCertificateAndAadDetails: %+v", err)
	}

	return encoded, nil
}
