package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationOperationPredicate struct {
	ConnectionName                 *string
	CreatedDateTime                *string
	Description                    *string
	DisableOnDemandUserOverride    *bool
	DisconnectOnIdle               *bool
	DisconnectOnIdleTimerInSeconds *int64
	DisplayName                    *string
	EnablePerApp                   *bool
	EnableSplitTunneling           *bool
	Id                             *string
	Identifier                     *string
	LastModifiedDateTime           *string
	LoginGroupOrDomain             *string
	ODataType                      *string
	OptInToDeviceIdSharing         *bool
	Realm                          *string
	Role                           *string
	SupportsScopeTags              *bool
	Version                        *int64
}

func (p AppleVpnConfigurationOperationPredicate) Matches(input AppleVpnConfiguration) bool {

	if p.ConnectionName != nil && (input.ConnectionName == nil || *p.ConnectionName != *input.ConnectionName) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisableOnDemandUserOverride != nil && (input.DisableOnDemandUserOverride == nil || *p.DisableOnDemandUserOverride != *input.DisableOnDemandUserOverride) {
		return false
	}

	if p.DisconnectOnIdle != nil && (input.DisconnectOnIdle == nil || *p.DisconnectOnIdle != *input.DisconnectOnIdle) {
		return false
	}

	if p.DisconnectOnIdleTimerInSeconds != nil && (input.DisconnectOnIdleTimerInSeconds == nil || *p.DisconnectOnIdleTimerInSeconds != *input.DisconnectOnIdleTimerInSeconds) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnablePerApp != nil && (input.EnablePerApp == nil || *p.EnablePerApp != *input.EnablePerApp) {
		return false
	}

	if p.EnableSplitTunneling != nil && (input.EnableSplitTunneling == nil || *p.EnableSplitTunneling != *input.EnableSplitTunneling) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Identifier != nil && (input.Identifier == nil || *p.Identifier != *input.Identifier) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.LoginGroupOrDomain != nil && (input.LoginGroupOrDomain == nil || *p.LoginGroupOrDomain != *input.LoginGroupOrDomain) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OptInToDeviceIdSharing != nil && (input.OptInToDeviceIdSharing == nil || *p.OptInToDeviceIdSharing != *input.OptInToDeviceIdSharing) {
		return false
	}

	if p.Realm != nil && (input.Realm == nil || *p.Realm != *input.Realm) {
		return false
	}

	if p.Role != nil && (input.Role == nil || *p.Role != *input.Role) {
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
