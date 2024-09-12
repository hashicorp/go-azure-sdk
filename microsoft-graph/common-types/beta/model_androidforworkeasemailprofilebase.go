package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBase interface {
	Entity
	DeviceConfiguration
	AndroidForWorkEasEmailProfileBase() BaseAndroidForWorkEasEmailProfileBaseImpl
}

var _ AndroidForWorkEasEmailProfileBase = BaseAndroidForWorkEasEmailProfileBaseImpl{}

type BaseAndroidForWorkEasEmailProfileBaseImpl struct {
	// Exchange Active Sync authentication method.
	AuthenticationMethod *EasAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Possible values for email sync duration.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`

	// Possible values for username source or email source.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`

	// Exchange location (URL) that the mail app connects to.
	HostName *string `json:"hostName,omitempty"`

	// Identity certificate.
	IdentityCertificate *AndroidForWorkCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`

	// Android username source.
	UsernameSource *AndroidUsernameSource `json:"usernameSource,omitempty"`

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

func (s BaseAndroidForWorkEasEmailProfileBaseImpl) AndroidForWorkEasEmailProfileBase() BaseAndroidForWorkEasEmailProfileBaseImpl {
	return s
}

func (s BaseAndroidForWorkEasEmailProfileBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s BaseAndroidForWorkEasEmailProfileBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidForWorkEasEmailProfileBase = RawAndroidForWorkEasEmailProfileBaseImpl{}

// RawAndroidForWorkEasEmailProfileBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidForWorkEasEmailProfileBaseImpl struct {
	androidForWorkEasEmailProfileBase BaseAndroidForWorkEasEmailProfileBaseImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawAndroidForWorkEasEmailProfileBaseImpl) AndroidForWorkEasEmailProfileBase() BaseAndroidForWorkEasEmailProfileBaseImpl {
	return s.androidForWorkEasEmailProfileBase
}

func (s RawAndroidForWorkEasEmailProfileBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.androidForWorkEasEmailProfileBase.DeviceConfiguration()
}

func (s RawAndroidForWorkEasEmailProfileBaseImpl) Entity() BaseEntityImpl {
	return s.androidForWorkEasEmailProfileBase.Entity()
}

var _ json.Marshaler = BaseAndroidForWorkEasEmailProfileBaseImpl{}

func (s BaseAndroidForWorkEasEmailProfileBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidForWorkEasEmailProfileBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidForWorkEasEmailProfileBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidForWorkEasEmailProfileBaseImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.androidForWorkEasEmailProfileBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidForWorkEasEmailProfileBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAndroidForWorkEasEmailProfileBaseImpl{}

func (s *BaseAndroidForWorkEasEmailProfileBaseImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseAndroidForWorkEasEmailProfileBaseImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseAndroidForWorkEasEmailProfileBaseImpl: %+v", err)
	}

	s.Assignments = decoded.Assignments
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.DurationOfEmailToSync = decoded.DurationOfEmailToSync
	s.EmailAddressSource = decoded.EmailAddressSource
	s.GroupAssignments = decoded.GroupAssignments
	s.HostName = decoded.HostName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RequireSsl = decoded.RequireSsl
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.UsernameSource = decoded.UsernameSource
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAndroidForWorkEasEmailProfileBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalAndroidForWorkCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'BaseAndroidForWorkEasEmailProfileBaseImpl': %+v", err)
		}
		s.IdentityCertificate = &impl
	}
	return nil
}

func UnmarshalAndroidForWorkEasEmailProfileBaseImplementation(input []byte) (AndroidForWorkEasEmailProfileBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkEasEmailProfileBase into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkGmailEasConfiguration") {
		var out AndroidForWorkGmailEasConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkGmailEasConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkNineWorkEasConfiguration") {
		var out AndroidForWorkNineWorkEasConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkNineWorkEasConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidForWorkEasEmailProfileBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidForWorkEasEmailProfileBaseImpl: %+v", err)
	}

	return RawAndroidForWorkEasEmailProfileBaseImpl{
		androidForWorkEasEmailProfileBase: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
