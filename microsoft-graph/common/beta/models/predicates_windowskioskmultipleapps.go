package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskMultipleAppsOperationPredicate struct {
	AllowAccessToDownloadsFolder *bool
	DisallowDesktopApps          *bool
	ODataType                    *string
	ShowTaskBar                  *bool
	StartMenuLayoutXml           *string
}

func (p WindowsKioskMultipleAppsOperationPredicate) Matches(input WindowsKioskMultipleApps) bool {

	if p.AllowAccessToDownloadsFolder != nil && (input.AllowAccessToDownloadsFolder == nil || *p.AllowAccessToDownloadsFolder != *input.AllowAccessToDownloadsFolder) {
		return false
	}

	if p.DisallowDesktopApps != nil && (input.DisallowDesktopApps == nil || *p.DisallowDesktopApps != *input.DisallowDesktopApps) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShowTaskBar != nil && (input.ShowTaskBar == nil || *p.ShowTaskBar != *input.ShowTaskBar) {
		return false
	}

	if p.StartMenuLayoutXml != nil && (input.StartMenuLayoutXml == nil || *p.StartMenuLayoutXml != *input.StartMenuLayoutXml) {
		return false
	}

	return true
}
