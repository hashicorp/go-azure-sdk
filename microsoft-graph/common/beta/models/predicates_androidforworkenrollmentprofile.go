package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnrollmentProfileOperationPredicate struct {
	AccountId               *string
	CreatedDateTime         *string
	Description             *string
	DisplayName             *string
	EnrolledDeviceCount     *int64
	Id                      *string
	LastModifiedDateTime    *string
	ODataType               *string
	QrCodeContent           *string
	TokenExpirationDateTime *string
	TokenValue              *string
}

func (p AndroidForWorkEnrollmentProfileOperationPredicate) Matches(input AndroidForWorkEnrollmentProfile) bool {

	if p.AccountId != nil && (input.AccountId == nil || *p.AccountId != *input.AccountId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnrolledDeviceCount != nil && (input.EnrolledDeviceCount == nil || *p.EnrolledDeviceCount != *input.EnrolledDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QrCodeContent != nil && (input.QrCodeContent == nil || *p.QrCodeContent != *input.QrCodeContent) {
		return false
	}

	if p.TokenExpirationDateTime != nil && (input.TokenExpirationDateTime == nil || *p.TokenExpirationDateTime != *input.TokenExpirationDateTime) {
		return false
	}

	if p.TokenValue != nil && (input.TokenValue == nil || *p.TokenValue != *input.TokenValue) {
		return false
	}

	return true
}
