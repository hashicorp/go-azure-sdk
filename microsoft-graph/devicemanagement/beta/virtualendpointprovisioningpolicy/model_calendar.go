package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Calendar struct {
	AllowedOnlineMeetingProviders *CalendarAllowedOnlineMeetingProviders `json:"allowedOnlineMeetingProviders,omitempty"`
	CalendarGroupId               *string                                `json:"calendarGroupId,omitempty"`
	CalendarPermissions           *[]CalendarPermission                  `json:"calendarPermissions,omitempty"`
	CalendarView                  *[]Event                               `json:"calendarView,omitempty"`
	CanEdit                       *bool                                  `json:"canEdit,omitempty"`
	CanShare                      *bool                                  `json:"canShare,omitempty"`
	CanViewPrivateItems           *bool                                  `json:"canViewPrivateItems,omitempty"`
	ChangeKey                     *string                                `json:"changeKey,omitempty"`
	Color                         *CalendarColor                         `json:"color,omitempty"`
	DefaultOnlineMeetingProvider  *CalendarDefaultOnlineMeetingProvider  `json:"defaultOnlineMeetingProvider,omitempty"`
	Events                        *[]Event                               `json:"events,omitempty"`
	HexColor                      *string                                `json:"hexColor,omitempty"`
	Id                            *string                                `json:"id,omitempty"`
	IsDefaultCalendar             *bool                                  `json:"isDefaultCalendar,omitempty"`
	IsRemovable                   *bool                                  `json:"isRemovable,omitempty"`
	IsShared                      *bool                                  `json:"isShared,omitempty"`
	IsSharedWithMe                *bool                                  `json:"isSharedWithMe,omitempty"`
	IsTallyingResponses           *bool                                  `json:"isTallyingResponses,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty    `json:"multiValueExtendedProperties,omitempty"`
	Name                          *string                                `json:"name,omitempty"`
	ODataType                     *string                                `json:"@odata.type,omitempty"`
	Owner                         *EmailAddress                          `json:"owner,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty   `json:"singleValueExtendedProperties,omitempty"`
}
