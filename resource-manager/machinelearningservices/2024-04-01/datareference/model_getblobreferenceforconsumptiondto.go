package datareference

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBlobReferenceForConsumptionDto struct {
	BlobUri             *string                 `json:"blobUri,omitempty"`
	Credential          DataReferenceCredential `json:"credential"`
	StorageAccountArmId *string                 `json:"storageAccountArmId,omitempty"`
}

var _ json.Unmarshaler = &GetBlobReferenceForConsumptionDto{}

func (s *GetBlobReferenceForConsumptionDto) UnmarshalJSON(bytes []byte) error {
	type alias GetBlobReferenceForConsumptionDto
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into GetBlobReferenceForConsumptionDto: %+v", err)
	}

	s.BlobUri = decoded.BlobUri
	s.StorageAccountArmId = decoded.StorageAccountArmId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GetBlobReferenceForConsumptionDto into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["credential"]; ok {
		impl, err := UnmarshalDataReferenceCredentialImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Credential' for 'GetBlobReferenceForConsumptionDto': %+v", err)
		}
		s.Credential = impl
	}
	return nil
}
