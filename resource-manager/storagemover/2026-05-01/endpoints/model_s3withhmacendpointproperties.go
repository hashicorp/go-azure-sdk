package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointBaseProperties = S3WithHmacEndpointProperties{}

type S3WithHmacEndpointProperties struct {
	Credentials                *AzureKeyVaultS3WithHmacCredentials `json:"credentials,omitempty"`
	OtherSourceTypeDescription *string                             `json:"otherSourceTypeDescription,omitempty"`
	SourceType                 *S3WithHmacSourceType               `json:"sourceType,omitempty"`
	SourceUri                  *string                             `json:"sourceUri,omitempty"`

	// Fields inherited from EndpointBaseProperties

	Description       *string            `json:"description,omitempty"`
	EndpointKind      *EndpointKind      `json:"endpointKind,omitempty"`
	EndpointType      EndpointType       `json:"endpointType"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

func (s S3WithHmacEndpointProperties) EndpointBaseProperties() BaseEndpointBasePropertiesImpl {
	return BaseEndpointBasePropertiesImpl{
		Description:       s.Description,
		EndpointKind:      s.EndpointKind,
		EndpointType:      s.EndpointType,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = S3WithHmacEndpointProperties{}

func (s S3WithHmacEndpointProperties) MarshalJSON() ([]byte, error) {
	type wrapper S3WithHmacEndpointProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling S3WithHmacEndpointProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling S3WithHmacEndpointProperties: %+v", err)
	}

	decoded["endpointType"] = "S3WithHMAC"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling S3WithHmacEndpointProperties: %+v", err)
	}

	return encoded, nil
}
