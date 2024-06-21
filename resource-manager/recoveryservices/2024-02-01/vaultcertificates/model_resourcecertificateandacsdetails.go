package vaultcertificates

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceCertificateDetails = ResourceCertificateAndAcsDetails{}

type ResourceCertificateAndAcsDetails struct {
	GlobalAcsHostName  string `json:"globalAcsHostName"`
	GlobalAcsNamespace string `json:"globalAcsNamespace"`
	GlobalAcsRPRealm   string `json:"globalAcsRPRealm"`

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

var _ json.Marshaler = ResourceCertificateAndAcsDetails{}

func (s ResourceCertificateAndAcsDetails) MarshalJSON() ([]byte, error) {
	type wrapper ResourceCertificateAndAcsDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceCertificateAndAcsDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceCertificateAndAcsDetails: %+v", err)
	}
	decoded["authType"] = "AccessControlService"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceCertificateAndAcsDetails: %+v", err)
	}

	return encoded, nil
}
