package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettingsOperationPredicate struct {
	AppName                  *string
	BadgesEnabled            *bool
	BundleID                 *string
	Enabled                  *bool
	ODataType                *string
	Publisher                *string
	ShowInNotificationCenter *bool
	ShowOnLockScreen         *bool
	SoundsEnabled            *bool
}

func (p IosNotificationSettingsOperationPredicate) Matches(input IosNotificationSettings) bool {

	if p.AppName != nil && (input.AppName == nil || *p.AppName != *input.AppName) {
		return false
	}

	if p.BadgesEnabled != nil && (input.BadgesEnabled == nil || *p.BadgesEnabled != *input.BadgesEnabled) {
		return false
	}

	if p.BundleID != nil && (input.BundleID == nil || *p.BundleID != *input.BundleID) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.ShowInNotificationCenter != nil && (input.ShowInNotificationCenter == nil || *p.ShowInNotificationCenter != *input.ShowInNotificationCenter) {
		return false
	}

	if p.ShowOnLockScreen != nil && (input.ShowOnLockScreen == nil || *p.ShowOnLockScreen != *input.ShowOnLockScreen) {
		return false
	}

	if p.SoundsEnabled != nil && (input.SoundsEnabled == nil || *p.SoundsEnabled != *input.SoundsEnabled) {
		return false
	}

	return true
}
