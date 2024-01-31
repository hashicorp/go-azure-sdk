package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = GetTdeCertificatesSqlTaskProperties{}

type GetTdeCertificatesSqlTaskProperties struct {
	Input  *GetTdeCertificatesSqlTaskInput    `json:"input,omitempty"`
	Output *[]GetTdeCertificatesSqlTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = GetTdeCertificatesSqlTaskProperties{}

func (s GetTdeCertificatesSqlTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper GetTdeCertificatesSqlTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GetTdeCertificatesSqlTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GetTdeCertificatesSqlTaskProperties: %+v", err)
	}
	decoded["taskType"] = "GetTDECertificates.Sql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GetTdeCertificatesSqlTaskProperties: %+v", err)
	}

	return encoded, nil
}
