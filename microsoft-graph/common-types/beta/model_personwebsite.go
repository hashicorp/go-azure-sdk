package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ItemFacet = PersonWebsite{}

type PersonWebsite struct {
	// Contains categories a user has associated with the website (for example, personal, recipes).
	Categories *[]string `json:"categories,omitempty"`

	// Contains a description of the website.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Contains a friendly name for the website.
	DisplayName *string `json:"displayName,omitempty"`

	ThumbnailUrl nullable.Type[string] `json:"thumbnailUrl,omitempty"`

	// Contains a link to the website itself.
	WebUrl *string `json:"webUrl,omitempty"`

	// Fields inherited from ItemFacet

	// The audiences that are able to see the values contained within the associated entity. Possible values are: me,
	// family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
	AllowedAudiences *AllowedAudiences `json:"allowedAudiences,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Provides the dateTimeOffset for when the entity was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains inference detail if the entity is inferred by the creating or modifying application.
	Inference *InferenceData `json:"inference,omitempty"`

	IsSearchable   nullable.Type[bool] `json:"isSearchable,omitempty"`
	LastModifiedBy IdentitySet         `json:"lastModifiedBy"`

	// Provides the dateTimeOffset for when the entity was created.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Where the values within an entity originated if synced from another service.
	Source *PersonDataSources `json:"source,omitempty"`

	// Where the values within an entity originated if synced from another source.
	Sources *[]ProfileSourceAnnotation `json:"sources,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PersonWebsite) ItemFacet() BaseItemFacetImpl {
	return BaseItemFacetImpl{
		AllowedAudiences:     s.AllowedAudiences,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Inference:            s.Inference,
		IsSearchable:         s.IsSearchable,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Source:               s.Source,
		Sources:              s.Sources,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s PersonWebsite) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PersonWebsite{}

func (s PersonWebsite) MarshalJSON() ([]byte, error) {
	type wrapper PersonWebsite
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PersonWebsite: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PersonWebsite: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.personWebsite"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PersonWebsite: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PersonWebsite{}

func (s *PersonWebsite) UnmarshalJSON(bytes []byte) error {
	type alias PersonWebsite
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PersonWebsite: %+v", err)
	}

	s.AllowedAudiences = decoded.AllowedAudiences
	s.Categories = decoded.Categories
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.Inference = decoded.Inference
	s.IsSearchable = decoded.IsSearchable
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source
	s.Sources = decoded.Sources
	s.ThumbnailUrl = decoded.ThumbnailUrl
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PersonWebsite into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PersonWebsite': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PersonWebsite': %+v", err)
		}
		s.LastModifiedBy = impl
	}
	return nil
}
