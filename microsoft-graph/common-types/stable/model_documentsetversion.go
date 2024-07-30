package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetVersion struct {
	Comment                   *string                   `json:"comment,omitempty"`
	CreatedBy                 *IdentitySet              `json:"createdBy,omitempty"`
	CreatedDateTime           *string                   `json:"createdDateTime,omitempty"`
	Fields                    *FieldValueSet            `json:"fields,omitempty"`
	Id                        *string                   `json:"id,omitempty"`
	Items                     *[]DocumentSetVersionItem `json:"items,omitempty"`
	LastModifiedBy            *IdentitySet              `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime      *string                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string                   `json:"@odata.type,omitempty"`
	Publication               *PublicationFacet         `json:"publication,omitempty"`
	ShouldCaptureMinorVersion *bool                     `json:"shouldCaptureMinorVersion,omitempty"`
}
