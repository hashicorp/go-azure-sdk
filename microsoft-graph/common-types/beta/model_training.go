package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Training{}

type Training struct {
	// The training availability status. Possible values are: unknown, notAvailable, available, archive, delete,
	// unknownFutureValue.
	AvailabilityStatus *TrainingAvailabilityStatus `json:"availabilityStatus,omitempty"`

	// The identity of the user who created the training.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// The date and time when the training was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The training URL
	CustomUrl nullable.Type[string] `json:"customUrl,omitempty"`

	// The description for the training.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the training.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The duration of the training in minutes.
	DurationInMinutes nullable.Type[int64] `json:"durationInMinutes,omitempty"`

	// Whether the training has any evaluation.
	HasEvaluation nullable.Type[bool] `json:"hasEvaluation,omitempty"`

	// Details about the language used in the training.
	LanguageDetails *[]TrainingLanguageDetail `json:"languageDetails,omitempty"`

	// The identity of the user who last modified the training.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// The date and time when the training was last modified. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The source of the training content. Possible values are: unknown, global, tenant, unknownFutureValue.
	Source *SimulationContentSource `json:"source,omitempty"`

	// The supported locales for content for the associated training.
	SupportedLocales *[]string `json:"supportedLocales,omitempty"`

	// Training tags.
	Tags *[]string `json:"tags,omitempty"`

	// The type of training. Possible values are: unknown, phishing, unknownFutureValue.
	Type *TrainingType `json:"type,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Training) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Training{}

func (s Training) MarshalJSON() ([]byte, error) {
	type wrapper Training
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Training: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Training: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.training"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Training: %+v", err)
	}

	return encoded, nil
}
