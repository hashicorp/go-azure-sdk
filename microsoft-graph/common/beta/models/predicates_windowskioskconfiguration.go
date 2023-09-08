package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskConfigurationOperationPredicate struct {
	CreatedDateTime                        *string
	Description                            *string
	DisplayName                            *string
	EdgeKioskEnablePublicBrowsing          *bool
	Id                                     *string
	KioskBrowserDefaultUrl                 *string
	KioskBrowserEnableEndSessionButton     *bool
	KioskBrowserEnableHomeButton           *bool
	KioskBrowserEnableNavigationButtons    *bool
	KioskBrowserRestartOnIdleTimeInMinutes *int64
	LastModifiedDateTime                   *string
	ODataType                              *string
	SupportsScopeTags                      *bool
	Version                                *int64
}

func (p WindowsKioskConfigurationOperationPredicate) Matches(input WindowsKioskConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EdgeKioskEnablePublicBrowsing != nil && (input.EdgeKioskEnablePublicBrowsing == nil || *p.EdgeKioskEnablePublicBrowsing != *input.EdgeKioskEnablePublicBrowsing) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.KioskBrowserDefaultUrl != nil && (input.KioskBrowserDefaultUrl == nil || *p.KioskBrowserDefaultUrl != *input.KioskBrowserDefaultUrl) {
		return false
	}

	if p.KioskBrowserEnableEndSessionButton != nil && (input.KioskBrowserEnableEndSessionButton == nil || *p.KioskBrowserEnableEndSessionButton != *input.KioskBrowserEnableEndSessionButton) {
		return false
	}

	if p.KioskBrowserEnableHomeButton != nil && (input.KioskBrowserEnableHomeButton == nil || *p.KioskBrowserEnableHomeButton != *input.KioskBrowserEnableHomeButton) {
		return false
	}

	if p.KioskBrowserEnableNavigationButtons != nil && (input.KioskBrowserEnableNavigationButtons == nil || *p.KioskBrowserEnableNavigationButtons != *input.KioskBrowserEnableNavigationButtons) {
		return false
	}

	if p.KioskBrowserRestartOnIdleTimeInMinutes != nil && (input.KioskBrowserRestartOnIdleTimeInMinutes == nil || *p.KioskBrowserRestartOnIdleTimeInMinutes != *input.KioskBrowserRestartOnIdleTimeInMinutes) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
