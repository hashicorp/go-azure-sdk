package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingGroupDefinition = DeviceManagementConfigurationSettingGroupCollectionDefinition{}

type DeviceManagementConfigurationSettingGroupCollectionDefinition struct {
	// Maximum number of setting group count in the collection
	MaximumCount *int64 `json:"maximumCount,omitempty"`

	// Minimum number of setting group count in the collection
	MinimumCount *int64 `json:"minimumCount,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingGroupDefinition

	// Dependent child settings to this group of settings.
	ChildIds *[]string `json:"childIds,omitempty"`

	// List of child settings that depend on this setting
	DependedOnBy *[]DeviceManagementConfigurationSettingDependedOnBy `json:"dependedOnBy,omitempty"`

	// List of Dependencies for the setting group
	DependentOn *[]DeviceManagementConfigurationDependentOn `json:"dependentOn,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingDefinition

	AccessTypes *DeviceManagementConfigurationSettingAccessTypes `json:"accessTypes,omitempty"`

	// Details which device setting is applicable on. Supports: $filters.
	Applicability DeviceManagementConfigurationSettingApplicability `json:"applicability"`

	// Base CSP Path
	BaseUri nullable.Type[string] `json:"baseUri,omitempty"`

	// Specify category in which the setting is under. Support $filters.
	CategoryId nullable.Type[string] `json:"categoryId,omitempty"`

	// Description of the setting.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the setting. For example: Allow Toast.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Help text of the setting. Give more details of the setting.
	HelpText nullable.Type[string] `json:"helpText,omitempty"`

	// List of links more info for the setting can be found at.
	InfoUrls *[]string `json:"infoUrls,omitempty"`

	// Tokens which to search settings on
	Keywords *[]string `json:"keywords,omitempty"`

	// Name of the item
	Name nullable.Type[string] `json:"name,omitempty"`

	// Indicates whether the setting is required or not
	Occurrence *DeviceManagementConfigurationSettingOccurrence `json:"occurrence,omitempty"`

	// Offset CSP Path from Base
	OffsetUri nullable.Type[string] `json:"offsetUri,omitempty"`

	// List of referred setting information.
	ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation `json:"referredSettingInformationList,omitempty"`

	// Setting RiskLevel
	RiskLevel *DeviceManagementConfigurationSettingRiskLevel `json:"riskLevel,omitempty"`

	// Root setting definition id if the setting is a child setting.
	RootDefinitionId nullable.Type[string] `json:"rootDefinitionId,omitempty"`

	// Supported setting types
	SettingUsage *DeviceManagementConfigurationSettingUsage `json:"settingUsage,omitempty"`

	// Setting control type representation in the UX
	UxBehavior *DeviceManagementConfigurationControlType `json:"uxBehavior,omitempty"`

	// Item Version
	Version nullable.Type[string] `json:"version,omitempty"`

	// Supported setting types
	Visibility *DeviceManagementConfigurationSettingVisibility `json:"visibility,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationSettingGroupCollectionDefinition) DeviceManagementConfigurationSettingGroupDefinition() BaseDeviceManagementConfigurationSettingGroupDefinitionImpl {
	return BaseDeviceManagementConfigurationSettingGroupDefinitionImpl{
		ChildIds:                       s.ChildIds,
		DependedOnBy:                   s.DependedOnBy,
		DependentOn:                    s.DependentOn,
		AccessTypes:                    s.AccessTypes,
		Applicability:                  s.Applicability,
		BaseUri:                        s.BaseUri,
		CategoryId:                     s.CategoryId,
		Description:                    s.Description,
		DisplayName:                    s.DisplayName,
		HelpText:                       s.HelpText,
		InfoUrls:                       s.InfoUrls,
		Keywords:                       s.Keywords,
		Name:                           s.Name,
		Occurrence:                     s.Occurrence,
		OffsetUri:                      s.OffsetUri,
		ReferredSettingInformationList: s.ReferredSettingInformationList,
		RiskLevel:                      s.RiskLevel,
		RootDefinitionId:               s.RootDefinitionId,
		SettingUsage:                   s.SettingUsage,
		UxBehavior:                     s.UxBehavior,
		Version:                        s.Version,
		Visibility:                     s.Visibility,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s DeviceManagementConfigurationSettingGroupCollectionDefinition) DeviceManagementConfigurationSettingDefinition() BaseDeviceManagementConfigurationSettingDefinitionImpl {
	return BaseDeviceManagementConfigurationSettingDefinitionImpl{
		AccessTypes:                    s.AccessTypes,
		Applicability:                  s.Applicability,
		BaseUri:                        s.BaseUri,
		CategoryId:                     s.CategoryId,
		Description:                    s.Description,
		DisplayName:                    s.DisplayName,
		HelpText:                       s.HelpText,
		InfoUrls:                       s.InfoUrls,
		Keywords:                       s.Keywords,
		Name:                           s.Name,
		Occurrence:                     s.Occurrence,
		OffsetUri:                      s.OffsetUri,
		ReferredSettingInformationList: s.ReferredSettingInformationList,
		RiskLevel:                      s.RiskLevel,
		RootDefinitionId:               s.RootDefinitionId,
		SettingUsage:                   s.SettingUsage,
		UxBehavior:                     s.UxBehavior,
		Version:                        s.Version,
		Visibility:                     s.Visibility,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s DeviceManagementConfigurationSettingGroupCollectionDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSettingGroupCollectionDefinition{}

func (s DeviceManagementConfigurationSettingGroupCollectionDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSettingGroupCollectionDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSettingGroupCollectionDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingGroupCollectionDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSettingGroupCollectionDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSettingGroupCollectionDefinition{}

func (s *DeviceManagementConfigurationSettingGroupCollectionDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		MaximumCount                   *int64                                                     `json:"maximumCount,omitempty"`
		MinimumCount                   *int64                                                     `json:"minimumCount,omitempty"`
		ChildIds                       *[]string                                                  `json:"childIds,omitempty"`
		DependedOnBy                   *[]DeviceManagementConfigurationSettingDependedOnBy        `json:"dependedOnBy,omitempty"`
		DependentOn                    *[]DeviceManagementConfigurationDependentOn                `json:"dependentOn,omitempty"`
		AccessTypes                    *DeviceManagementConfigurationSettingAccessTypes           `json:"accessTypes,omitempty"`
		BaseUri                        nullable.Type[string]                                      `json:"baseUri,omitempty"`
		CategoryId                     nullable.Type[string]                                      `json:"categoryId,omitempty"`
		Description                    nullable.Type[string]                                      `json:"description,omitempty"`
		DisplayName                    nullable.Type[string]                                      `json:"displayName,omitempty"`
		HelpText                       nullable.Type[string]                                      `json:"helpText,omitempty"`
		InfoUrls                       *[]string                                                  `json:"infoUrls,omitempty"`
		Keywords                       *[]string                                                  `json:"keywords,omitempty"`
		Name                           nullable.Type[string]                                      `json:"name,omitempty"`
		Occurrence                     *DeviceManagementConfigurationSettingOccurrence            `json:"occurrence,omitempty"`
		OffsetUri                      nullable.Type[string]                                      `json:"offsetUri,omitempty"`
		ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation `json:"referredSettingInformationList,omitempty"`
		RiskLevel                      *DeviceManagementConfigurationSettingRiskLevel             `json:"riskLevel,omitempty"`
		RootDefinitionId               nullable.Type[string]                                      `json:"rootDefinitionId,omitempty"`
		SettingUsage                   *DeviceManagementConfigurationSettingUsage                 `json:"settingUsage,omitempty"`
		UxBehavior                     *DeviceManagementConfigurationControlType                  `json:"uxBehavior,omitempty"`
		Version                        nullable.Type[string]                                      `json:"version,omitempty"`
		Visibility                     *DeviceManagementConfigurationSettingVisibility            `json:"visibility,omitempty"`
		Id                             *string                                                    `json:"id,omitempty"`
		ODataId                        *string                                                    `json:"@odata.id,omitempty"`
		ODataType                      *string                                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.MaximumCount = decoded.MaximumCount
	s.MinimumCount = decoded.MinimumCount
	s.AccessTypes = decoded.AccessTypes
	s.BaseUri = decoded.BaseUri
	s.CategoryId = decoded.CategoryId
	s.ChildIds = decoded.ChildIds
	s.DependedOnBy = decoded.DependedOnBy
	s.DependentOn = decoded.DependentOn
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.HelpText = decoded.HelpText
	s.Id = decoded.Id
	s.InfoUrls = decoded.InfoUrls
	s.Keywords = decoded.Keywords
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Occurrence = decoded.Occurrence
	s.OffsetUri = decoded.OffsetUri
	s.ReferredSettingInformationList = decoded.ReferredSettingInformationList
	s.RiskLevel = decoded.RiskLevel
	s.RootDefinitionId = decoded.RootDefinitionId
	s.SettingUsage = decoded.SettingUsage
	s.UxBehavior = decoded.UxBehavior
	s.Version = decoded.Version
	s.Visibility = decoded.Visibility

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingGroupCollectionDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["applicability"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingApplicabilityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Applicability' for 'DeviceManagementConfigurationSettingGroupCollectionDefinition': %+v", err)
		}
		s.Applicability = impl
	}

	return nil
}
