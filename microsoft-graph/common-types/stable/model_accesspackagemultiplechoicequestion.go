package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageQuestion = AccessPackageMultipleChoiceQuestion{}

type AccessPackageMultipleChoiceQuestion struct {
	// List of answer choices.
	Choices *[]AccessPackageAnswerChoice `json:"choices,omitempty"`

	// Indicates whether requestor can select multiple choices as their answer.
	IsMultipleSelectionAllowed nullable.Type[bool] `json:"isMultipleSelectionAllowed,omitempty"`

	// Fields inherited from AccessPackageQuestion

	// Specifies whether the requestor is allowed to edit answers to questions for an assignment by posting an update to
	// accessPackageAssignmentRequest.
	IsAnswerEditable nullable.Type[bool] `json:"isAnswerEditable,omitempty"`

	// Whether the requestor is required to supply an answer or not.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

	// The text of the question represented in a format for a specific locale.
	Localizations *[]AccessPackageLocalizedText `json:"localizations,omitempty"`

	// Relative position of this question when displaying a list of questions to the requestor.
	Sequence nullable.Type[int64] `json:"sequence,omitempty"`

	// The text of the question to show to the requestor.
	Text nullable.Type[string] `json:"text,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AccessPackageMultipleChoiceQuestion) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return BaseAccessPackageQuestionImpl{
		IsAnswerEditable: s.IsAnswerEditable,
		IsRequired:       s.IsRequired,
		Localizations:    s.Localizations,
		Sequence:         s.Sequence,
		Text:             s.Text,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s AccessPackageMultipleChoiceQuestion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageMultipleChoiceQuestion{}

func (s AccessPackageMultipleChoiceQuestion) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageMultipleChoiceQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.accessPackageMultipleChoiceQuestion"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	return encoded, nil
}
