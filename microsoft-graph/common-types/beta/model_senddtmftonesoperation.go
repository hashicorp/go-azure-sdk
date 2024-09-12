package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CommsOperation = SendDtmfTonesOperation{}

type SendDtmfTonesOperation struct {
	// The results of the action. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled,
	// unknownfutureValue.
	CompletionReason *SendDtmfCompletionReason `json:"completionReason,omitempty"`

	// Fields inherited from CommsOperation

	// Unique Client Context string. Max limit is 256 chars.
	ClientContext nullable.Type[string] `json:"clientContext,omitempty"`

	// The result information. Read-only.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	Status *OperationStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SendDtmfTonesOperation) CommsOperation() BaseCommsOperationImpl {
	return BaseCommsOperationImpl{
		ClientContext: s.ClientContext,
		ResultInfo:    s.ResultInfo,
		Status:        s.Status,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s SendDtmfTonesOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SendDtmfTonesOperation{}

func (s SendDtmfTonesOperation) MarshalJSON() ([]byte, error) {
	type wrapper SendDtmfTonesOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SendDtmfTonesOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SendDtmfTonesOperation: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.sendDtmfTonesOperation"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SendDtmfTonesOperation: %+v", err)
	}

	return encoded, nil
}
