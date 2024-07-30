package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedback struct {
	FeedbackBy       *IdentitySet       `json:"feedbackBy,omitempty"`
	FeedbackDateTime *string            `json:"feedbackDateTime,omitempty"`
	ODataType        *string            `json:"@odata.type,omitempty"`
	Text             *EducationItemBody `json:"text,omitempty"`
}
