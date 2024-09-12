package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExactMatchSessionBase = ExactMatchSession{}

type ExactMatchSession struct {
	Checksum      nullable.Type[string]  `json:"checksum,omitempty"`
	DataUploadURI nullable.Type[string]  `json:"dataUploadURI,omitempty"`
	Fields        *[]string              `json:"fields,omitempty"`
	FileName      nullable.Type[string]  `json:"fileName,omitempty"`
	RowsPerBlock  nullable.Type[int64]   `json:"rowsPerBlock,omitempty"`
	Salt          nullable.Type[string]  `json:"salt,omitempty"`
	UploadAgent   *ExactMatchUploadAgent `json:"uploadAgent,omitempty"`
	UploadAgentId nullable.Type[string]  `json:"uploadAgentId,omitempty"`

	// Fields inherited from ExactMatchSessionBase

	DataStoreId                  nullable.Type[string] `json:"dataStoreId,omitempty"`
	ProcessingCompletionDateTime nullable.Type[string] `json:"processingCompletionDateTime,omitempty"`
	RemainingBlockCount          nullable.Type[int64]  `json:"remainingBlockCount,omitempty"`
	RemainingJobCount            nullable.Type[int64]  `json:"remainingJobCount,omitempty"`
	State                        nullable.Type[string] `json:"state,omitempty"`
	TotalBlockCount              nullable.Type[int64]  `json:"totalBlockCount,omitempty"`
	TotalJobCount                nullable.Type[int64]  `json:"totalJobCount,omitempty"`
	UploadCompletionDateTime     nullable.Type[string] `json:"uploadCompletionDateTime,omitempty"`

	// Fields inherited from ExactMatchJobBase

	CompletionDateTime  nullable.Type[string] `json:"completionDateTime,omitempty"`
	CreationDateTime    nullable.Type[string] `json:"creationDateTime,omitempty"`
	Error               *ClassificationError  `json:"error,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	StartDateTime       nullable.Type[string] `json:"startDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ExactMatchSession) ExactMatchSessionBase() BaseExactMatchSessionBaseImpl {
	return BaseExactMatchSessionBaseImpl{
		DataStoreId:                  s.DataStoreId,
		ProcessingCompletionDateTime: s.ProcessingCompletionDateTime,
		RemainingBlockCount:          s.RemainingBlockCount,
		RemainingJobCount:            s.RemainingJobCount,
		State:                        s.State,
		TotalBlockCount:              s.TotalBlockCount,
		TotalJobCount:                s.TotalJobCount,
		UploadCompletionDateTime:     s.UploadCompletionDateTime,
		CompletionDateTime:           s.CompletionDateTime,
		CreationDateTime:             s.CreationDateTime,
		Error:                        s.Error,
		LastUpdatedDateTime:          s.LastUpdatedDateTime,
		StartDateTime:                s.StartDateTime,
		Id:                           s.Id,
		ODataId:                      s.ODataId,
		ODataType:                    s.ODataType,
	}
}

func (s ExactMatchSession) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return BaseExactMatchJobBaseImpl{
		CompletionDateTime:  s.CompletionDateTime,
		CreationDateTime:    s.CreationDateTime,
		Error:               s.Error,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		StartDateTime:       s.StartDateTime,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s ExactMatchSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExactMatchSession{}

func (s ExactMatchSession) MarshalJSON() ([]byte, error) {
	type wrapper ExactMatchSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExactMatchSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchSession: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.exactMatchSession"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExactMatchSession: %+v", err)
	}

	return encoded, nil
}
