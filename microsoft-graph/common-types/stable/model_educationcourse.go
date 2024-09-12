package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationCourse struct {
	// Unique identifier for the course.
	CourseNumber nullable.Type[string] `json:"courseNumber,omitempty"`

	// Description of the course.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the course.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// ID of the course from the syncing system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Subject of the course.
	Subject nullable.Type[string] `json:"subject,omitempty"`
}
