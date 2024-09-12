package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = ListItem{}

type ListItem struct {
	// The list of recent activities that took place on this item.
	Activities *[]ItemActivityOLD `json:"activities,omitempty"`

	// Analytics about the view activities that took place on this item.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`

	// The content type of this list item
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`

	Deleted *Deleted `json:"deleted,omitempty"`

	// Version information for a document set version created by a user.
	DocumentSetVersions *[]DocumentSetVersion `json:"documentSetVersions,omitempty"`

	// For document libraries, the driveItem relationship exposes the listItem as a driveItem
	DriveItem *DriveItem `json:"driveItem,omitempty"`

	// The values of the columns set on this list item.
	Fields *FieldValueSet `json:"fields,omitempty"`

	// The set of permissions for the item. Read-only. Nullable.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// The list of previous versions of the list item.
	Versions *[]ListItemVersion `json:"versions,omitempty"`

	// Fields inherited from BaseItem

	// Identity of the user, device, or application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	CreatedByUser *User `json:"createdByUser,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the item.
	Description nullable.Type[string] `json:"description,omitempty"`

	// ETag for the item. Read-only.
	ETag nullable.Type[string] `json:"eTag,omitempty"`

	// Identity of the user, device, and application that last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The name of the item. Read-write.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for
	// other formats). Read-only.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ListItem) BaseItem() BaseBaseItemImpl {
	return BaseBaseItemImpl{
		CreatedBy:            s.CreatedBy,
		CreatedByUser:        s.CreatedByUser,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		ETag:                 s.ETag,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedByUser:   s.LastModifiedByUser,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		ParentReference:      s.ParentReference,
		WebUrl:               s.WebUrl,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ListItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ListItem{}

func (s ListItem) MarshalJSON() ([]byte, error) {
	type wrapper ListItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ListItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ListItem: %+v", err)
	}

	delete(decoded, "permissions")
	delete(decoded, "sharepointIds")
	decoded["@odata.type"] = "#microsoft.graph.listItem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ListItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ListItem{}

func (s *ListItem) UnmarshalJSON(bytes []byte) error {
	type alias ListItem
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ListItem: %+v", err)
	}

	s.Activities = decoded.Activities
	s.Analytics = decoded.Analytics
	s.ContentType = decoded.ContentType
	s.CreatedByUser = decoded.CreatedByUser
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Deleted = decoded.Deleted
	s.Description = decoded.Description
	s.DocumentSetVersions = decoded.DocumentSetVersions
	s.DriveItem = decoded.DriveItem
	s.ETag = decoded.ETag
	s.Fields = decoded.Fields
	s.Id = decoded.Id
	s.LastModifiedByUser = decoded.LastModifiedByUser
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParentReference = decoded.ParentReference
	s.Permissions = decoded.Permissions
	s.SharepointIds = decoded.SharepointIds
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ListItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'ListItem': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ListItem': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	if v, ok := temp["versions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Versions into list []json.RawMessage: %+v", err)
		}

		output := make([]ListItemVersion, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalListItemVersionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Versions' for 'ListItem': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Versions = &output
	}
	return nil
}
