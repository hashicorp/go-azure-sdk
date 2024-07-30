package securityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentInfo struct {
	Format     *ContentInfoFormat `json:"format,omitempty"`
	Identifier *string            `json:"identifier,omitempty"`
	Metadata   *[]KeyValuePair    `json:"metadata,omitempty"`
	ODataType  *string            `json:"@odata.type,omitempty"`
	State      *ContentInfoState  `json:"state,omitempty"`
}
