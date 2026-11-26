package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointBaseUpdateProperties = S3WithHmacEndpointUpdateProperties{}

type S3WithHmacEndpointUpdateProperties struct {
	Credentials *AzureKeyVaultS3WithHmacCredentials `json:"credentials,omitempty"`

	// Fields inherited from EndpointBaseUpdateProperties

	Description  *string      `json:"description,omitempty"`
	EndpointType EndpointType `json:"endpointType"`
}

func (s S3WithHmacEndpointUpdateProperties) EndpointBaseUpdateProperties() BaseEndpointBaseUpdatePropertiesImpl {
	return BaseEndpointBaseUpdatePropertiesImpl{
		Description:  s.Description,
		EndpointType: s.EndpointType,
	}
}

var _ json.Marshaler = S3WithHmacEndpointUpdateProperties{}

func (s S3WithHmacEndpointUpdateProperties) MarshalJSON() ([]byte, error) {
	type wrapper S3WithHmacEndpointUpdateProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling S3WithHmacEndpointUpdateProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling S3WithHmacEndpointUpdateProperties: %+v", err)
	}

	decoded["endpointType"] = "S3WithHMAC"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling S3WithHmacEndpointUpdateProperties: %+v", err)
	}

	return encoded, nil
}
