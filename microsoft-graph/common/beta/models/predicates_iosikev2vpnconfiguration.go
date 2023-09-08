package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationOperationPredicate struct {
	AllowDefaultChildSecurityAssociationParameters *bool
	AllowDefaultSecurityAssociationParameters      *bool
	CloudName                                      *string
	ConnectionName                                 *string
	CreatedDateTime                                *string
	Description                                    *string
	DisableMobilityAndMultihoming                  *bool
	DisableOnDemandUserOverride                    *bool
	DisableRedirect                                *bool
	DisconnectOnIdle                               *bool
	DisconnectOnIdleTimerInSeconds                 *int64
	DisplayName                                    *string
	EnableAlwaysOnConfiguration                    *bool
	EnableCertificateRevocationCheck               *bool
	EnableEAP                                      *bool
	EnablePerApp                                   *bool
	EnablePerfectForwardSecrecy                    *bool
	EnableSplitTunneling                           *bool
	EnableUseInternalSubnetAttributes              *bool
	Id                                             *string
	Identifier                                     *string
	LastModifiedDateTime                           *string
	LoginGroupOrDomain                             *string
	MicrosoftTunnelSiteId                          *string
	MtuSizeInBytes                                 *int64
	ODataType                                      *string
	OptInToDeviceIdSharing                         *bool
	Realm                                          *string
	RemoteIdentifier                               *string
	Role                                           *string
	ServerCertificateCommonName                    *string
	ServerCertificateIssuerCommonName              *string
	SharedSecret                                   *string
	StrictEnforcement                              *bool
	SupportsScopeTags                              *bool
	TlsMaximumVersion                              *string
	TlsMinimumVersion                              *string
	UserDomain                                     *string
	Version                                        *int64
}

func (p IosikEv2VpnConfigurationOperationPredicate) Matches(input IosikEv2VpnConfiguration) bool {

	if p.AllowDefaultChildSecurityAssociationParameters != nil && (input.AllowDefaultChildSecurityAssociationParameters == nil || *p.AllowDefaultChildSecurityAssociationParameters != *input.AllowDefaultChildSecurityAssociationParameters) {
		return false
	}

	if p.AllowDefaultSecurityAssociationParameters != nil && (input.AllowDefaultSecurityAssociationParameters == nil || *p.AllowDefaultSecurityAssociationParameters != *input.AllowDefaultSecurityAssociationParameters) {
		return false
	}

	if p.CloudName != nil && (input.CloudName == nil || *p.CloudName != *input.CloudName) {
		return false
	}

	if p.ConnectionName != nil && (input.ConnectionName == nil || *p.ConnectionName != *input.ConnectionName) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisableMobilityAndMultihoming != nil && (input.DisableMobilityAndMultihoming == nil || *p.DisableMobilityAndMultihoming != *input.DisableMobilityAndMultihoming) {
		return false
	}

	if p.DisableOnDemandUserOverride != nil && (input.DisableOnDemandUserOverride == nil || *p.DisableOnDemandUserOverride != *input.DisableOnDemandUserOverride) {
		return false
	}

	if p.DisableRedirect != nil && (input.DisableRedirect == nil || *p.DisableRedirect != *input.DisableRedirect) {
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

	if p.EnableAlwaysOnConfiguration != nil && (input.EnableAlwaysOnConfiguration == nil || *p.EnableAlwaysOnConfiguration != *input.EnableAlwaysOnConfiguration) {
		return false
	}

	if p.EnableCertificateRevocationCheck != nil && (input.EnableCertificateRevocationCheck == nil || *p.EnableCertificateRevocationCheck != *input.EnableCertificateRevocationCheck) {
		return false
	}

	if p.EnableEAP != nil && (input.EnableEAP == nil || *p.EnableEAP != *input.EnableEAP) {
		return false
	}

	if p.EnablePerApp != nil && (input.EnablePerApp == nil || *p.EnablePerApp != *input.EnablePerApp) {
		return false
	}

	if p.EnablePerfectForwardSecrecy != nil && (input.EnablePerfectForwardSecrecy == nil || *p.EnablePerfectForwardSecrecy != *input.EnablePerfectForwardSecrecy) {
		return false
	}

	if p.EnableSplitTunneling != nil && (input.EnableSplitTunneling == nil || *p.EnableSplitTunneling != *input.EnableSplitTunneling) {
		return false
	}

	if p.EnableUseInternalSubnetAttributes != nil && (input.EnableUseInternalSubnetAttributes == nil || *p.EnableUseInternalSubnetAttributes != *input.EnableUseInternalSubnetAttributes) {
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

	if p.MicrosoftTunnelSiteId != nil && (input.MicrosoftTunnelSiteId == nil || *p.MicrosoftTunnelSiteId != *input.MicrosoftTunnelSiteId) {
		return false
	}

	if p.MtuSizeInBytes != nil && (input.MtuSizeInBytes == nil || *p.MtuSizeInBytes != *input.MtuSizeInBytes) {
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

	if p.RemoteIdentifier != nil && (input.RemoteIdentifier == nil || *p.RemoteIdentifier != *input.RemoteIdentifier) {
		return false
	}

	if p.Role != nil && (input.Role == nil || *p.Role != *input.Role) {
		return false
	}

	if p.ServerCertificateCommonName != nil && (input.ServerCertificateCommonName == nil || *p.ServerCertificateCommonName != *input.ServerCertificateCommonName) {
		return false
	}

	if p.ServerCertificateIssuerCommonName != nil && (input.ServerCertificateIssuerCommonName == nil || *p.ServerCertificateIssuerCommonName != *input.ServerCertificateIssuerCommonName) {
		return false
	}

	if p.SharedSecret != nil && (input.SharedSecret == nil || *p.SharedSecret != *input.SharedSecret) {
		return false
	}

	if p.StrictEnforcement != nil && (input.StrictEnforcement == nil || *p.StrictEnforcement != *input.StrictEnforcement) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.TlsMaximumVersion != nil && (input.TlsMaximumVersion == nil || *p.TlsMaximumVersion != *input.TlsMaximumVersion) {
		return false
	}

	if p.TlsMinimumVersion != nil && (input.TlsMinimumVersion == nil || *p.TlsMinimumVersion != *input.TlsMinimumVersion) {
		return false
	}

	if p.UserDomain != nil && (input.UserDomain == nil || *p.UserDomain != *input.UserDomain) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
