package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubject struct {
	AltSecId                     *string                               `json:"altSecId,omitempty"`
	CleanupScheduledDateTime     *string                               `json:"cleanupScheduledDateTime,omitempty"`
	ConnectedOrganization        *ConnectedOrganization                `json:"connectedOrganization,omitempty"`
	ConnectedOrganizationId      *string                               `json:"connectedOrganizationId,omitempty"`
	DisplayName                  *string                               `json:"displayName,omitempty"`
	Email                        *string                               `json:"email,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	ODataType                    *string                               `json:"@odata.type,omitempty"`
	ObjectId                     *string                               `json:"objectId,omitempty"`
	OnPremisesSecurityIdentifier *string                               `json:"onPremisesSecurityIdentifier,omitempty"`
	PrincipalName                *string                               `json:"principalName,omitempty"`
	SubjectLifecycle             *AccessPackageSubjectSubjectLifecycle `json:"subjectLifecycle,omitempty"`
	Type                         *string                               `json:"type,omitempty"`
}
