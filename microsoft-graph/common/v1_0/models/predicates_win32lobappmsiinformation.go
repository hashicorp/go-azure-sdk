package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiInformationOperationPredicate struct {
	ODataType      *string
	ProductCode    *string
	ProductName    *string
	ProductVersion *string
	Publisher      *string
	RequiresReboot *bool
	UpgradeCode    *string
}

func (p Win32LobAppMsiInformationOperationPredicate) Matches(input Win32LobAppMsiInformation) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProductCode != nil && (input.ProductCode == nil || *p.ProductCode != *input.ProductCode) {
		return false
	}

	if p.ProductName != nil && (input.ProductName == nil || *p.ProductName != *input.ProductName) {
		return false
	}

	if p.ProductVersion != nil && (input.ProductVersion == nil || *p.ProductVersion != *input.ProductVersion) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.RequiresReboot != nil && (input.RequiresReboot == nil || *p.RequiresReboot != *input.RequiresReboot) {
		return false
	}

	if p.UpgradeCode != nil && (input.UpgradeCode == nil || *p.UpgradeCode != *input.UpgradeCode) {
		return false
	}

	return true
}
