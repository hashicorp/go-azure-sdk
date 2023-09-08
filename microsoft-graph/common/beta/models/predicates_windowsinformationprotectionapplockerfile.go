package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLockerFileOperationPredicate struct {
	DisplayName *string
	File        *string
	FileHash    *string
	Id          *string
	ODataType   *string
	Version     *string
}

func (p WindowsInformationProtectionAppLockerFileOperationPredicate) Matches(input WindowsInformationProtectionAppLockerFile) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.File != nil && (input.File == nil || *p.File != *input.File) {
		return false
	}

	if p.FileHash != nil && (input.FileHash == nil || *p.FileHash != *input.FileHash) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
