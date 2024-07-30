package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubject struct {
	ConnectedOrganization        *ConnectedOrganization           `json:"connectedOrganization,omitempty"`
	DisplayName                  *string                          `json:"displayName,omitempty"`
	Email                        *string                          `json:"email,omitempty"`
	Id                           *string                          `json:"id,omitempty"`
	ODataType                    *string                          `json:"@odata.type,omitempty"`
	ObjectId                     *string                          `json:"objectId,omitempty"`
	OnPremisesSecurityIdentifier *string                          `json:"onPremisesSecurityIdentifier,omitempty"`
	PrincipalName                *string                          `json:"principalName,omitempty"`
	SubjectType                  *AccessPackageSubjectSubjectType `json:"subjectType,omitempty"`
}
