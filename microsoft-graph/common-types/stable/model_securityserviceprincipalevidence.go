package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalEvidence struct {
	AppId                    *string                                               `json:"appId,omitempty"`
	AppOwnerTenantId         *string                                               `json:"appOwnerTenantId,omitempty"`
	CreatedDateTime          *string                                               `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                             `json:"detailedRoles,omitempty"`
	ODataType                *string                                               `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityServicePrincipalEvidenceRemediationStatus    `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                               `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityServicePrincipalEvidenceRoles                `json:"roles,omitempty"`
	ServicePrincipalName     *string                                               `json:"servicePrincipalName,omitempty"`
	ServicePrincipalObjectId *string                                               `json:"servicePrincipalObjectId,omitempty"`
	ServicePrincipalType     *SecurityServicePrincipalEvidenceServicePrincipalType `json:"servicePrincipalType,omitempty"`
	Tags                     *[]string                                             `json:"tags,omitempty"`
	TenantId                 *string                                               `json:"tenantId,omitempty"`
	Verdict                  *SecurityServicePrincipalEvidenceVerdict              `json:"verdict,omitempty"`
}
