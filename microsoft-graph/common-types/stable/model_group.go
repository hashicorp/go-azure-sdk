package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Group struct {
	AcceptedSenders               *[]DirectoryObject                 `json:"acceptedSenders,omitempty"`
	AllowExternalSenders          *bool                              `json:"allowExternalSenders,omitempty"`
	AppRoleAssignments            *[]AppRoleAssignment               `json:"appRoleAssignments,omitempty"`
	AssignedLabels                *[]AssignedLabel                   `json:"assignedLabels,omitempty"`
	AssignedLicenses              *[]AssignedLicense                 `json:"assignedLicenses,omitempty"`
	AutoSubscribeNewMembers       *bool                              `json:"autoSubscribeNewMembers,omitempty"`
	Calendar                      *Calendar                          `json:"calendar,omitempty"`
	CalendarView                  *[]Event                           `json:"calendarView,omitempty"`
	Classification                *string                            `json:"classification,omitempty"`
	Conversations                 *[]Conversation                    `json:"conversations,omitempty"`
	CreatedDateTime               *string                            `json:"createdDateTime,omitempty"`
	CreatedOnBehalfOf             *DirectoryObject                   `json:"createdOnBehalfOf,omitempty"`
	DeletedDateTime               *string                            `json:"deletedDateTime,omitempty"`
	Description                   *string                            `json:"description,omitempty"`
	DisplayName                   *string                            `json:"displayName,omitempty"`
	Drive                         *Drive                             `json:"drive,omitempty"`
	Drives                        *[]Drive                           `json:"drives,omitempty"`
	Events                        *[]Event                           `json:"events,omitempty"`
	ExpirationDateTime            *string                            `json:"expirationDateTime,omitempty"`
	Extensions                    *[]Extension                       `json:"extensions,omitempty"`
	GroupLifecyclePolicies        *[]GroupLifecyclePolicy            `json:"groupLifecyclePolicies,omitempty"`
	GroupTypes                    *[]string                          `json:"groupTypes,omitempty"`
	HasMembersWithLicenseErrors   *bool                              `json:"hasMembersWithLicenseErrors,omitempty"`
	HideFromAddressLists          *bool                              `json:"hideFromAddressLists,omitempty"`
	HideFromOutlookClients        *bool                              `json:"hideFromOutlookClients,omitempty"`
	Id                            *string                            `json:"id,omitempty"`
	IsArchived                    *bool                              `json:"isArchived,omitempty"`
	IsAssignableToRole            *bool                              `json:"isAssignableToRole,omitempty"`
	IsSubscribedByMail            *bool                              `json:"isSubscribedByMail,omitempty"`
	LicenseProcessingState        *LicenseProcessingState            `json:"licenseProcessingState,omitempty"`
	Mail                          *string                            `json:"mail,omitempty"`
	MailEnabled                   *bool                              `json:"mailEnabled,omitempty"`
	MailNickname                  *string                            `json:"mailNickname,omitempty"`
	MemberOf                      *[]DirectoryObject                 `json:"memberOf,omitempty"`
	Members                       *[]DirectoryObject                 `json:"members,omitempty"`
	MembersWithLicenseErrors      *[]DirectoryObject                 `json:"membersWithLicenseErrors,omitempty"`
	MembershipRule                *string                            `json:"membershipRule,omitempty"`
	MembershipRuleProcessingState *string                            `json:"membershipRuleProcessingState,omitempty"`
	ODataType                     *string                            `json:"@odata.type,omitempty"`
	OnPremisesDomainName          *string                            `json:"onPremisesDomainName,omitempty"`
	OnPremisesLastSyncDateTime    *string                            `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesNetBiosName         *string                            `json:"onPremisesNetBiosName,omitempty"`
	OnPremisesProvisioningErrors  *[]OnPremisesProvisioningError     `json:"onPremisesProvisioningErrors,omitempty"`
	OnPremisesSamAccountName      *string                            `json:"onPremisesSamAccountName,omitempty"`
	OnPremisesSecurityIdentifier  *string                            `json:"onPremisesSecurityIdentifier,omitempty"`
	OnPremisesSyncEnabled         *bool                              `json:"onPremisesSyncEnabled,omitempty"`
	Onenote                       *Onenote                           `json:"onenote,omitempty"`
	Owners                        *[]DirectoryObject                 `json:"owners,omitempty"`
	PermissionGrants              *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	Photo                         *ProfilePhoto                      `json:"photo,omitempty"`
	Photos                        *[]ProfilePhoto                    `json:"photos,omitempty"`
	Planner                       *PlannerGroup                      `json:"planner,omitempty"`
	PreferredDataLocation         *string                            `json:"preferredDataLocation,omitempty"`
	PreferredLanguage             *string                            `json:"preferredLanguage,omitempty"`
	ProxyAddresses                *[]string                          `json:"proxyAddresses,omitempty"`
	RejectedSenders               *[]DirectoryObject                 `json:"rejectedSenders,omitempty"`
	RenewedDateTime               *string                            `json:"renewedDateTime,omitempty"`
	SecurityEnabled               *bool                              `json:"securityEnabled,omitempty"`
	SecurityIdentifier            *string                            `json:"securityIdentifier,omitempty"`
	ServiceProvisioningErrors     *[]ServiceProvisioningError        `json:"serviceProvisioningErrors,omitempty"`
	Settings                      *[]GroupSetting                    `json:"settings,omitempty"`
	Sites                         *[]Site                            `json:"sites,omitempty"`
	Team                          *Team                              `json:"team,omitempty"`
	Theme                         *string                            `json:"theme,omitempty"`
	Threads                       *[]ConversationThread              `json:"threads,omitempty"`
	TransitiveMemberOf            *[]DirectoryObject                 `json:"transitiveMemberOf,omitempty"`
	TransitiveMembers             *[]DirectoryObject                 `json:"transitiveMembers,omitempty"`
	UnseenCount                   *int64                             `json:"unseenCount,omitempty"`
	Visibility                    *string                            `json:"visibility,omitempty"`
}
