package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfiguration struct {
	Assignments                                               *[]DeviceConfigurationAssignment                                  `json:"assignments,omitempty"`
	BackgroundDownloadFromHttpDelayInSeconds                  *int64                                                            `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`
	BandwidthMode                                             *DeliveryOptimizationBandwidth                                    `json:"bandwidthMode,omitempty"`
	CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds *int64                                                            `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerForegroundDownloadFallbackToHttpDelayInSeconds *int64                                                            `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerHostNames                                      *[]string                                                         `json:"cacheServerHostNames,omitempty"`
	CreatedDateTime                                           *string                                                           `json:"createdDateTime,omitempty"`
	DeliveryOptimizationMode                                  *WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode `json:"deliveryOptimizationMode,omitempty"`
	Description                                               *string                                                           `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode               *DeviceManagementApplicabilityRuleDeviceMode                      `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition                *DeviceManagementApplicabilityRuleOsEdition                       `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion                *DeviceManagementApplicabilityRuleOsVersion                       `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                               *[]SettingStateDeviceSummary                                      `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                      *DeviceConfigurationDeviceOverview                                `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                            *[]DeviceConfigurationDeviceStatus                                `json:"deviceStatuses,omitempty"`
	DisplayName                                               *string                                                           `json:"displayName,omitempty"`
	ForegroundDownloadFromHttpDelayInSeconds                  *int64                                                            `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`
	GroupAssignments                                          *[]DeviceConfigurationGroupAssignment                             `json:"groupAssignments,omitempty"`
	GroupIdSource                                             *DeliveryOptimizationGroupIdSource                                `json:"groupIdSource,omitempty"`
	Id                                                        *string                                                           `json:"id,omitempty"`
	LastModifiedDateTime                                      *string                                                           `json:"lastModifiedDateTime,omitempty"`
	MaximumCacheAgeInDays                                     *int64                                                            `json:"maximumCacheAgeInDays,omitempty"`
	MaximumCacheSize                                          *DeliveryOptimizationMaxCacheSize                                 `json:"maximumCacheSize,omitempty"`
	MinimumBatteryPercentageAllowedToUpload                   *int64                                                            `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`
	MinimumDiskSizeAllowedToPeerInGigabytes                   *int64                                                            `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`
	MinimumFileSizeToCacheInMegabytes                         *int64                                                            `json:"minimumFileSizeToCacheInMegabytes,omitempty"`
	MinimumRamAllowedToPeerInGigabytes                        *int64                                                            `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`
	ModifyCacheLocation                                       *string                                                           `json:"modifyCacheLocation,omitempty"`
	ODataType                                                 *string                                                           `json:"@odata.type,omitempty"`
	RestrictPeerSelectionBy                                   *WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy  `json:"restrictPeerSelectionBy,omitempty"`
	RoleScopeTagIds                                           *[]string                                                         `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                                         *bool                                                             `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                                        *DeviceConfigurationUserOverview                                  `json:"userStatusOverview,omitempty"`
	UserStatuses                                              *[]DeviceConfigurationUserStatus                                  `json:"userStatuses,omitempty"`
	Version                                                   *int64                                                            `json:"version,omitempty"`
	VpnPeerCaching                                            *WindowsDeliveryOptimizationConfigurationVpnPeerCaching           `json:"vpnPeerCaching,omitempty"`
}
