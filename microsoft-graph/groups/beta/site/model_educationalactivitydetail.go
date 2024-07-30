package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivityDetail struct {
	Abbreviation  *string   `json:"abbreviation,omitempty"`
	Activities    *[]string `json:"activities,omitempty"`
	Awards        *[]string `json:"awards,omitempty"`
	Description   *string   `json:"description,omitempty"`
	DisplayName   *string   `json:"displayName,omitempty"`
	FieldsOfStudy *[]string `json:"fieldsOfStudy,omitempty"`
	Grade         *string   `json:"grade,omitempty"`
	Notes         *string   `json:"notes,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
	WebUrl        *string   `json:"webUrl,omitempty"`
}
