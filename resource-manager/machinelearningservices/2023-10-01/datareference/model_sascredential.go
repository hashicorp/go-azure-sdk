package datareference

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataReferenceCredential = SASCredential{}

type SASCredential struct {
	SasUri *string `json:"sasUri,omitempty"`

	// Fields inherited from DataReferenceCredential
}

var _ json.Marshaler = SASCredential{}

func (s SASCredential) MarshalJSON() ([]byte, error) {
	type wrapper SASCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SASCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SASCredential: %+v", err)
	}
	decoded["credentialType"] = "SAS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SASCredential: %+v", err)
	}

	return encoded, nil
}
