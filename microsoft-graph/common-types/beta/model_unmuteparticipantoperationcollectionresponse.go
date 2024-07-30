package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root fo
 license information.

type UnmuteParticipantOperationCollectionResponse struct {
	ODataCount    *int64                        `json:"@odata.count,omitempty"`
	ODataNextLink *string                       `json:"@odata.nextLink,omitempty"`
	Value         *[]UnmuteParticipantOperation `json:"value,omitempty"`
}
