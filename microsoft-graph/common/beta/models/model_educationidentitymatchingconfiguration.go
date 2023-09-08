package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityMatchingConfiguration struct {
	MatchingOptions *[]EducationIdentityMatchingOptions `json:"matchingOptions,omitempty"`
	ODataType       *string                             `json:"@odata.type,omitempty"`
}
