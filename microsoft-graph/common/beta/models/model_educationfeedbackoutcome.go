package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedbackOutcome struct {
	Feedback             *EducationFeedback `json:"feedback,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet       `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string            `json:"@odata.type,omitempty"`
	PublishedFeedback    *EducationFeedback `json:"publishedFeedback,omitempty"`
}
