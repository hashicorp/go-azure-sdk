package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFileSynchronizationVerificationMessage struct {
	Description *string `json:"description,omitempty"`
	FileName    *string `json:"fileName,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Type        *string `json:"type,omitempty"`
}
