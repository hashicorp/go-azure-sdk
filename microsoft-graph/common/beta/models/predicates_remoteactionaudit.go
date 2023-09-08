package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAuditOperationPredicate struct {
	DeviceDisplayName            *string
	DeviceIMEI                   *string
	DeviceOwnerUserPrincipalName *string
	Id                           *string
	InitiatedByUserPrincipalName *string
	ManagedDeviceId              *string
	ODataType                    *string
	RequestDateTime              *string
	UserName                     *string
}

func (p RemoteActionAuditOperationPredicate) Matches(input RemoteActionAudit) bool {

	if p.DeviceDisplayName != nil && (input.DeviceDisplayName == nil || *p.DeviceDisplayName != *input.DeviceDisplayName) {
		return false
	}

	if p.DeviceIMEI != nil && (input.DeviceIMEI == nil || *p.DeviceIMEI != *input.DeviceIMEI) {
		return false
	}

	if p.DeviceOwnerUserPrincipalName != nil && (input.DeviceOwnerUserPrincipalName == nil || *p.DeviceOwnerUserPrincipalName != *input.DeviceOwnerUserPrincipalName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InitiatedByUserPrincipalName != nil && (input.InitiatedByUserPrincipalName == nil || *p.InitiatedByUserPrincipalName != *input.InitiatedByUserPrincipalName) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestDateTime != nil && (input.RequestDateTime == nil || *p.RequestDateTime != *input.RequestDateTime) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
