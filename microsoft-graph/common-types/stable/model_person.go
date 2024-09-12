package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Person{}

type Person struct {
	// The person's birthday.
	Birthday nullable.Type[string] `json:"birthday,omitempty"`

	// The name of the person's company.
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// The person's department.
	Department nullable.Type[string] `json:"department,omitempty"`

	// The person's display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The person's given name.
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only.
	ImAddress nullable.Type[string] `json:"imAddress,omitempty"`

	// True if the user has flagged this person as a favorite.
	IsFavorite nullable.Type[bool] `json:"isFavorite,omitempty"`

	// The person's job title.
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The location of the person's office.
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// Free-form notes that the user has taken about this person.
	PersonNotes nullable.Type[string] `json:"personNotes,omitempty"`

	// The type of person.
	PersonType *PersonType `json:"personType,omitempty"`

	// The person's phone numbers.
	Phones *[]Phone `json:"phones,omitempty"`

	// The person's addresses.
	PostalAddresses *[]Location `json:"postalAddresses,omitempty"`

	// The person's profession.
	Profession nullable.Type[string] `json:"profession,omitempty"`

	// The person's email addresses.
	ScoredEmailAddresses *[]ScoredEmailAddress `json:"scoredEmailAddresses,omitempty"`

	// The person's surname.
	Surname nullable.Type[string] `json:"surname,omitempty"`

	// The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the
	// Internet standard RFC 822. By convention, this should map to the person's email name. The general format is
	// alias@domain.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The person's websites.
	Websites *[]Website `json:"websites,omitempty"`

	// The phonetic Japanese name of the person's company.
	YomiCompany nullable.Type[string] `json:"yomiCompany,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Person) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Person{}

func (s Person) MarshalJSON() ([]byte, error) {
	type wrapper Person
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Person: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Person: %+v", err)
	}

	delete(decoded, "imAddress")
	decoded["@odata.type"] = "#microsoft.graph.person"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Person: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Person{}

func (s *Person) UnmarshalJSON(bytes []byte) error {
	type alias Person
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into Person: %+v", err)
	}

	s.Birthday = decoded.Birthday
	s.CompanyName = decoded.CompanyName
	s.Department = decoded.Department
	s.DisplayName = decoded.DisplayName
	s.GivenName = decoded.GivenName
	s.Id = decoded.Id
	s.ImAddress = decoded.ImAddress
	s.IsFavorite = decoded.IsFavorite
	s.JobTitle = decoded.JobTitle
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OfficeLocation = decoded.OfficeLocation
	s.PersonNotes = decoded.PersonNotes
	s.PersonType = decoded.PersonType
	s.Phones = decoded.Phones
	s.Profession = decoded.Profession
	s.ScoredEmailAddresses = decoded.ScoredEmailAddresses
	s.Surname = decoded.Surname
	s.UserPrincipalName = decoded.UserPrincipalName
	s.Websites = decoded.Websites
	s.YomiCompany = decoded.YomiCompany

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Person into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["postalAddresses"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PostalAddresses into list []json.RawMessage: %+v", err)
		}

		output := make([]Location, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLocationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PostalAddresses' for 'Person': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PostalAddresses = &output
	}
	return nil
}
