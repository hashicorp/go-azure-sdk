package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidence struct {
	AppId                    *int64                                             `json:"appId,omitempty"`
	CreatedDateTime          *string                                            `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                          `json:"detailedRoles,omitempty"`
	DisplayName              *string                                            `json:"displayName,omitempty"`
	InstanceId               *int64                                             `json:"instanceId,omitempty"`
	InstanceName             *string                                            `json:"instanceName,omitempty"`
	ODataType                *string                                            `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityCloudApplicationEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                            `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityCloudApplicationEvidenceRoles             `json:"roles,omitempty"`
	SaasAppId                *int64                                             `json:"saasAppId,omitempty"`
	Tags                     *[]string                                          `json:"tags,omitempty"`
	Verdict                  *SecurityCloudApplicationEvidenceVerdict           `json:"verdict,omitempty"`
}
