package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnersBillingOperation = PartnersBillingRunningOperation{}

type PartnersBillingRunningOperation struct {

	// Fields inherited from PartnersBillingOperation

	// The start time of the operation. The timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The time of the last action of the operation. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`

	// The status of the operation. Possible values are: notStarted, running, completed, failed, unknownFutureValue.
	Status *LongRunningOperationStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PartnersBillingRunningOperation) PartnersBillingOperation() BasePartnersBillingOperationImpl {
	return BasePartnersBillingOperationImpl{
		CreatedDateTime:    s.CreatedDateTime,
		LastActionDateTime: s.LastActionDateTime,
		Status:             s.Status,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s PartnersBillingRunningOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnersBillingRunningOperation{}

func (s PartnersBillingRunningOperation) MarshalJSON() ([]byte, error) {
	type wrapper PartnersBillingRunningOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnersBillingRunningOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnersBillingRunningOperation: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.partners.billing.runningOperation"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnersBillingRunningOperation: %+v", err)
	}

	return encoded, nil
}
