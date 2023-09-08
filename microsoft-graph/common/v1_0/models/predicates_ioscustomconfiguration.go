package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCustomConfigurationOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Payload              *string
	PayloadFileName      *string
	PayloadName          *string
	Version              *int64
}

func (p IosCustomConfigurationOperationPredicate) Matches(input IosCustomConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.Payload != nil && (input.Payload == nil || *p.Payload != *input.Payload) {
		return false
	}

	if p.PayloadFileName != nil && (input.PayloadFileName == nil || *p.PayloadFileName != *input.PayloadFileName) {
		return false
	}

	if p.PayloadName != nil && (input.PayloadName == nil || *p.PayloadName != *input.PayloadName) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
