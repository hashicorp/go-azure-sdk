package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionStoreAppOperationPredicate struct {
	Denied        *bool
	Description   *string
	DisplayName   *string
	ODataType     *string
	ProductName   *string
	PublisherName *string
}

func (p WindowsInformationProtectionStoreAppOperationPredicate) Matches(input WindowsInformationProtectionStoreApp) bool {

	if p.Denied != nil && (input.Denied == nil || *p.Denied != *input.Denied) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProductName != nil && (input.ProductName == nil || *p.ProductName != *input.ProductName) {
		return false
	}

	if p.PublisherName != nil && (input.PublisherName == nil || *p.PublisherName != *input.PublisherName) {
		return false
	}

	return true
}
