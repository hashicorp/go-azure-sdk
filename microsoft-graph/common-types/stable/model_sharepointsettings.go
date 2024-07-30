package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettings struct {
	AllowedDomainGuidsForSyncApp                    *[]string                                       `json:"allowedDomainGuidsForSyncApp,omitempty"`
	AvailableManagedPathsForSiteCreation            *[]string                                       `json:"availableManagedPathsForSiteCreation,omitempty"`
	DeletedUserPersonalSiteRetentionPeriodInDays    *int64                                          `json:"deletedUserPersonalSiteRetentionPeriodInDays,omitempty"`
	ExcludedFileExtensionsForSyncApp                *[]string                                       `json:"excludedFileExtensionsForSyncApp,omitempty"`
	Id                                              *string                                         `json:"id,omitempty"`
	IdleSessionSignOut                              *IdleSessionSignOut                             `json:"idleSessionSignOut,omitempty"`
	ImageTaggingOption                              *SharepointSettingsImageTaggingOption           `json:"imageTaggingOption,omitempty"`
	IsCommentingOnSitePagesEnabled                  *bool                                           `json:"isCommentingOnSitePagesEnabled,omitempty"`
	IsFileActivityNotificationEnabled               *bool                                           `json:"isFileActivityNotificationEnabled,omitempty"`
	IsLegacyAuthProtocolsEnabled                    *bool                                           `json:"isLegacyAuthProtocolsEnabled,omitempty"`
	IsLoopEnabled                                   *bool                                           `json:"isLoopEnabled,omitempty"`
	IsMacSyncAppEnabled                             *bool                                           `json:"isMacSyncAppEnabled,omitempty"`
	IsRequireAcceptingUserToMatchInvitedUserEnabled *bool                                           `json:"isRequireAcceptingUserToMatchInvitedUserEnabled,omitempty"`
	IsResharingByExternalUsersEnabled               *bool                                           `json:"isResharingByExternalUsersEnabled,omitempty"`
	IsSharePointMobileNotificationEnabled           *bool                                           `json:"isSharePointMobileNotificationEnabled,omitempty"`
	IsSharePointNewsfeedEnabled                     *bool                                           `json:"isSharePointNewsfeedEnabled,omitempty"`
	IsSiteCreationEnabled                           *bool                                           `json:"isSiteCreationEnabled,omitempty"`
	IsSiteCreationUIEnabled                         *bool                                           `json:"isSiteCreationUIEnabled,omitempty"`
	IsSitePagesCreationEnabled                      *bool                                           `json:"isSitePagesCreationEnabled,omitempty"`
	IsSitesStorageLimitAutomatic                    *bool                                           `json:"isSitesStorageLimitAutomatic,omitempty"`
	IsSyncButtonHiddenOnPersonalSite                *bool                                           `json:"isSyncButtonHiddenOnPersonalSite,omitempty"`
	IsUnmanagedSyncAppForTenantRestricted           *bool                                           `json:"isUnmanagedSyncAppForTenantRestricted,omitempty"`
	ODataType                                       *string                                         `json:"@odata.type,omitempty"`
	PersonalSiteDefaultStorageLimitInMB             *int64                                          `json:"personalSiteDefaultStorageLimitInMB,omitempty"`
	SharingAllowedDomainList                        *[]string                                       `json:"sharingAllowedDomainList,omitempty"`
	SharingBlockedDomainList                        *[]string                                       `json:"sharingBlockedDomainList,omitempty"`
	SharingCapability                               *SharepointSettingsSharingCapability            `json:"sharingCapability,omitempty"`
	SharingDomainRestrictionMode                    *SharepointSettingsSharingDomainRestrictionMode `json:"sharingDomainRestrictionMode,omitempty"`
	SiteCreationDefaultManagedPath                  *string                                         `json:"siteCreationDefaultManagedPath,omitempty"`
	SiteCreationDefaultStorageLimitInMB             *int64                                          `json:"siteCreationDefaultStorageLimitInMB,omitempty"`
	TenantDefaultTimezone                           *string                                         `json:"tenantDefaultTimezone,omitempty"`
}
