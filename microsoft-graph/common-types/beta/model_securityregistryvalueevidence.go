package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityRegistryValueEvidence{}

type SecurityRegistryValueEvidence struct {
	// A unique identifier assigned to a device by Microsoft Defender for Endpoint.
	MdeDeviceId nullable.Type[string] `json:"mdeDeviceId,omitempty"`

	// Registry hive of the key that the recorded action was applied to.
	RegistryHive nullable.Type[string] `json:"registryHive,omitempty"`

	// Registry key that the recorded action was applied to.
	RegistryKey nullable.Type[string] `json:"registryKey,omitempty"`

	// Data of the registry value that the recorded action was applied to.
	RegistryValue nullable.Type[string] `json:"registryValue,omitempty"`

	// Name of the registry value that the recorded action was applied to.
	RegistryValueName nullable.Type[string] `json:"registryValueName,omitempty"`

	// Data type, such as binary or string, of the registry value that the recorded action was applied to.
	RegistryValueType nullable.Type[string] `json:"registryValueType,omitempty"`

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

func (s SecurityRegistryValueEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityRegistryValueEvidence{}

func (s SecurityRegistryValueEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRegistryValueEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRegistryValueEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRegistryValueEvidence: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.registryValueEvidence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRegistryValueEvidence: %+v", err)
	}

	return encoded, nil
}
