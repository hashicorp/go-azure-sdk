package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaConnectorOperationPredicate struct {
	EnrollmentAuthorizationUrl *string
	EnrollmentToken            *string
	FotaAppsApproved           *bool
	Id                         *string
	LastSyncDateTime           *string
	ODataType                  *string
}

func (p ZebraFotaConnectorOperationPredicate) Matches(input ZebraFotaConnector) bool {

	if p.EnrollmentAuthorizationUrl != nil && (input.EnrollmentAuthorizationUrl == nil || *p.EnrollmentAuthorizationUrl != *input.EnrollmentAuthorizationUrl) {
		return false
	}

	if p.EnrollmentToken != nil && (input.EnrollmentToken == nil || *p.EnrollmentToken != *input.EnrollmentToken) {
		return false
	}

	if p.FotaAppsApproved != nil && (input.FotaAppsApproved == nil || *p.FotaAppsApproved != *input.FotaAppsApproved) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
