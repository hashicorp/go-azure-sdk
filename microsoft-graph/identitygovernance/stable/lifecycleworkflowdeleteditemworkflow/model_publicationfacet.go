package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicationFacet struct {
	CheckedOutBy *IdentitySet `json:"checkedOutBy,omitempty"`
	Level        *string      `json:"level,omitempty"`
	ODataType    *string      `json:"@odata.type,omitempty"`
	VersionId    *string      `json:"versionId,omitempty"`
}
