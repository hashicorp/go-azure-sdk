package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeConstraint struct {
	// The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
	ActivityDomain *ActivityDomain `json:"activityDomain,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TimeSlots *[]TimeSlot `json:"timeSlots,omitempty"`
}
