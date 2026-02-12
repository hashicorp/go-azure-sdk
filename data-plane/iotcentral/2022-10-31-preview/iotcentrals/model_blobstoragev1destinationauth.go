package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobStorageV1DestinationAuth interface {
	BlobStorageV1DestinationAuth() BaseBlobStorageV1DestinationAuthImpl
}

var _ BlobStorageV1DestinationAuth = BaseBlobStorageV1DestinationAuthImpl{}

type BaseBlobStorageV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseBlobStorageV1DestinationAuthImpl) BlobStorageV1DestinationAuth() BaseBlobStorageV1DestinationAuthImpl {
	return s
}

var _ BlobStorageV1DestinationAuth = RawBlobStorageV1DestinationAuthImpl{}

// RawBlobStorageV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBlobStorageV1DestinationAuthImpl struct {
	blobStorageV1DestinationAuth BaseBlobStorageV1DestinationAuthImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawBlobStorageV1DestinationAuthImpl) BlobStorageV1DestinationAuth() BaseBlobStorageV1DestinationAuthImpl {
	return s.blobStorageV1DestinationAuth
}

func UnmarshalBlobStorageV1DestinationAuthImplementation(input []byte) (BlobStorageV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BlobStorageV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "connectionString") {
		var out BlobStorageV1DestinationConnectionStringAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobStorageV1DestinationConnectionStringAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "systemAssignedManagedIdentity") {
		var out BlobStorageV1DestinationSystemAssignedManagedIdentityAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobStorageV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseBlobStorageV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBlobStorageV1DestinationAuthImpl: %+v", err)
	}

	return RawBlobStorageV1DestinationAuthImpl{
		blobStorageV1DestinationAuth: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
