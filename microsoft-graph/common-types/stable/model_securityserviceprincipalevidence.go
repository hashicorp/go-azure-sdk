package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityServicePrincipalEvidence{}

type SecurityServicePrincipalEvidence struct {
	AppId                    nullable.Type[string]         `json:"appId,omitempty"`
	AppOwnerTenantId         nullable.Type[string]         `json:"appOwnerTenantId,omitempty"`
	ServicePrincipalName     nullable.Type[string]         `json:"servicePrincipalName,omitempty"`
	ServicePrincipalObjectId nullable.Type[string]         `json:"servicePrincipalObjectId,omitempty"`
	ServicePrincipalType     *SecurityServicePrincipalType `json:"servicePrincipalType,omitempty"`
	TenantId                 nullable.Type[string]         `json:"tenantId,omitempty"`

	// Fields inherited from SecurityAlertEvidence

	// The date and time when the evidence was created and added to the alert. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Detailed description of the entity role/s in an alert. Values are free-form.
	DetailedRoles *[]string `json:"detailedRoles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RemediationStatus *SecurityEvidenceRemediationStatus `json:"remediationStatus,omitempty"`

	// Details about the remediation status.
	RemediationStatusDetails nullable.Type[string] `json:"remediationStatusDetails,omitempty"`

	// The role/s that an evidence entity represents in an alert, for example, an IP address that is associated with an
	// attacker has the evidence role Attacker.
	Roles *[]SecurityEvidenceRole `json:"roles,omitempty"`

	// Array of custom tags associated with an evidence instance, for example, to denote a group of devices, high-value
	// assets, etc.
	Tags *[]string `json:"tags,omitempty"`

	Verdict *SecurityEvidenceVerdict `json:"verdict,omitempty"`
}

func (s SecurityServicePrincipalEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
	return BaseSecurityAlertEvidenceImpl{
		CreatedDateTime:          s.CreatedDateTime,
		DetailedRoles:            s.DetailedRoles,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
		RemediationStatus:        s.RemediationStatus,
		RemediationStatusDetails: s.RemediationStatusDetails,
		Roles:                    s.Roles,
		Tags:                     s.Tags,
		Verdict:                  s.Verdict,
	}
}

var _ json.Marshaler = SecurityServicePrincipalEvidence{}

func (s SecurityServicePrincipalEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityServicePrincipalEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityServicePrincipalEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityServicePrincipalEvidence: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.servicePrincipalEvidence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityServicePrincipalEvidence: %+v", err)
	}

	return encoded, nil
}
