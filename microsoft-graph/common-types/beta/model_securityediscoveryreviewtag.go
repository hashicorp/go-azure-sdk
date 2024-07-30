package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryReviewTag struct {
	ChildSelectability   *SecurityEdiscoveryReviewTagChildSelectability `json:"childSelectability,omitempty"`
	ChildTags            *[]SecurityEdiscoveryReviewTag                 `json:"childTags,omitempty"`
	CreatedBy            *IdentitySet                                   `json:"createdBy,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	DisplayName          *string                                        `json:"displayName,omitempty"`
	Id                   *string                                        `json:"id,omitempty"`
	LastModifiedDateTime *string                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                        `json:"@odata.type,omitempty"`
	Parent               *SecurityEdiscoveryReviewTag                   `json:"parent,omitempty"`
}
