package datareference

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataReferenceCredential = AnonymousAccessCredential{}

type AnonymousAccessCredential struct {

	// Fields inherited from DataReferenceCredential
}

var _ json.Marshaler = AnonymousAccessCredential{}

func (s AnonymousAccessCredential) MarshalJSON() ([]byte, error) {
	type wrapper AnonymousAccessCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AnonymousAccessCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AnonymousAccessCredential: %+v", err)
	}
	decoded["credentialType"] = "NoCredentials"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AnonymousAccessCredential: %+v", err)
	}

	return encoded, nil
}
