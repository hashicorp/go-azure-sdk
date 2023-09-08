package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentIndividualRecipient struct {
	ODataType  *string   `json:"@odata.type,omitempty"`
	Recipients *[]string `json:"recipients,omitempty"`
}
