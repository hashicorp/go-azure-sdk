package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDesktopAppOperationPredicate struct {
	BinaryName        *string
	BinaryVersionHigh *string
	BinaryVersionLow  *string
	Denied            *bool
	Description       *string
	DisplayName       *string
	ODataType         *string
	ProductName       *string
	PublisherName     *string
}

func (p WindowsInformationProtectionDesktopAppOperationPredicate) Matches(input WindowsInformationProtectionDesktopApp) bool {

	if p.BinaryName != nil && (input.BinaryName == nil || *p.BinaryName != *input.BinaryName) {
		return false
	}

	if p.BinaryVersionHigh != nil && (input.BinaryVersionHigh == nil || *p.BinaryVersionHigh != *input.BinaryVersionHigh) {
		return false
	}

	if p.BinaryVersionLow != nil && (input.BinaryVersionLow == nil || *p.BinaryVersionLow != *input.BinaryVersionLow) {
		return false
	}

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
