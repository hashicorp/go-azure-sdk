package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityIoTDeviceEvidence{}

type SecurityIoTDeviceEvidence struct {
	DeviceId           nullable.Type[string]            `json:"deviceId,omitempty"`
	DeviceName         nullable.Type[string]            `json:"deviceName,omitempty"`
	DevicePageLink     nullable.Type[string]            `json:"devicePageLink,omitempty"`
	DeviceSubType      nullable.Type[string]            `json:"deviceSubType,omitempty"`
	DeviceType         nullable.Type[string]            `json:"deviceType,omitempty"`
	IPAddress          *SecurityIPEvidence              `json:"ipAddress,omitempty"`
	Importance         *SecurityIoTDeviceImportanceType `json:"importance,omitempty"`
	IoTHub             *SecurityAzureResourceEvidence   `json:"ioTHub,omitempty"`
	IoTSecurityAgentId nullable.Type[string]            `json:"ioTSecurityAgentId,omitempty"`
	IsAuthorized       nullable.Type[bool]              `json:"isAuthorized,omitempty"`
	IsProgramming      nullable.Type[bool]              `json:"isProgramming,omitempty"`
	IsScanner          nullable.Type[bool]              `json:"isScanner,omitempty"`
	MacAddress         nullable.Type[string]            `json:"macAddress,omitempty"`
	Manufacturer       nullable.Type[string]            `json:"manufacturer,omitempty"`
	Model              nullable.Type[string]            `json:"model,omitempty"`
	Nics               *SecurityNicEvidence             `json:"nics,omitempty"`
	OperatingSystem    nullable.Type[string]            `json:"operatingSystem,omitempty"`
	Owners             *[]string                        `json:"owners,omitempty"`
	Protocols          *[]string                        `json:"protocols,omitempty"`
	PurdueLayer        nullable.Type[string]            `json:"purdueLayer,omitempty"`
	Sensor             nullable.Type[string]            `json:"sensor,omitempty"`
	SerialNumber       nullable.Type[string]            `json:"serialNumber,omitempty"`
	Site               nullable.Type[string]            `json:"site,omitempty"`
	Source             nullable.Type[string]            `json:"source,omitempty"`
	SourceRef          *SecurityUrlEvidence             `json:"sourceRef,omitempty"`
	Zone               nullable.Type[string]            `json:"zone,omitempty"`

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

func (s SecurityIoTDeviceEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
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

var _ json.Marshaler = SecurityIoTDeviceEvidence{}

func (s SecurityIoTDeviceEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIoTDeviceEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.ioTDeviceEvidence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIoTDeviceEvidence: %+v", err)
	}

	return encoded, nil
}
