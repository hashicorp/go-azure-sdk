package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBase interface {
	Entity
	DeviceConfiguration
	AndroidDeviceOwnerCertificateProfileBase() BaseAndroidDeviceOwnerCertificateProfileBaseImpl
}

var _ AndroidDeviceOwnerCertificateProfileBase = BaseAndroidDeviceOwnerCertificateProfileBaseImpl{}

type BaseAndroidDeviceOwnerCertificateProfileBaseImpl struct {
	// Certificate Validity Period Options.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`

	// Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int64 `json:"certificateValidityPeriodValue,omitempty"`

	// Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages *[]ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`

	// Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

	// Trusted Root Certificate.
	RootCertificate *AndroidDeviceOwnerTrustedRootCertificate `json:"rootCertificate,omitempty"`

	// Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName,
	// customAzureADAttribute, domainNameService, universalResourceIdentifier.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`

	// Certificate Subject Name Format. Possible values are: commonName, commonNameIncludingEmail, commonNameAsEmail,
	// custom, commonNameAsIMEI, commonNameAsSerialNumber, commonNameAsAadDeviceId, commonNameAsIntuneDeviceId,
	// commonNameAsDurableDeviceId.
	SubjectNameFormat *SubjectNameFormat `json:"subjectNameFormat,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAndroidDeviceOwnerCertificateProfileBaseImpl) AndroidDeviceOwnerCertificateProfileBase() BaseAndroidDeviceOwnerCertificateProfileBaseImpl {
	return s
}

func (s BaseAndroidDeviceOwnerCertificateProfileBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s BaseAndroidDeviceOwnerCertificateProfileBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidDeviceOwnerCertificateProfileBase = RawAndroidDeviceOwnerCertificateProfileBaseImpl{}

// RawAndroidDeviceOwnerCertificateProfileBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceOwnerCertificateProfileBaseImpl struct {
	androidDeviceOwnerCertificateProfileBase BaseAndroidDeviceOwnerCertificateProfileBaseImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawAndroidDeviceOwnerCertificateProfileBaseImpl) AndroidDeviceOwnerCertificateProfileBase() BaseAndroidDeviceOwnerCertificateProfileBaseImpl {
	return s.androidDeviceOwnerCertificateProfileBase
}

func (s RawAndroidDeviceOwnerCertificateProfileBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.androidDeviceOwnerCertificateProfileBase.DeviceConfiguration()
}

func (s RawAndroidDeviceOwnerCertificateProfileBaseImpl) Entity() BaseEntityImpl {
	return s.androidDeviceOwnerCertificateProfileBase.Entity()
}

var _ json.Marshaler = BaseAndroidDeviceOwnerCertificateProfileBaseImpl{}

func (s BaseAndroidDeviceOwnerCertificateProfileBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidDeviceOwnerCertificateProfileBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidDeviceOwnerCertificateProfileBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidDeviceOwnerCertificateProfileBaseImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerCertificateProfileBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidDeviceOwnerCertificateProfileBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAndroidDeviceOwnerCertificateProfileBaseImplementation(input []byte) (AndroidDeviceOwnerCertificateProfileBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerCertificateProfileBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfile") {
		var out AndroidDeviceOwnerImportedPFXCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerImportedPFXCertificateProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerPkcsCertificateProfile") {
		var out AndroidDeviceOwnerPkcsCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerPkcsCertificateProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerScepCertificateProfile") {
		var out AndroidDeviceOwnerScepCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerScepCertificateProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceOwnerCertificateProfileBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceOwnerCertificateProfileBaseImpl: %+v", err)
	}

	return RawAndroidDeviceOwnerCertificateProfileBaseImpl{
		androidDeviceOwnerCertificateProfileBase: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
