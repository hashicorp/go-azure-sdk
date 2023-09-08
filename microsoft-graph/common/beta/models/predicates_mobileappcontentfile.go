package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppContentFileOperationPredicate struct {
	AzureStorageUri                   *string
	AzureStorageUriExpirationDateTime *string
	CreatedDateTime                   *string
	Id                                *string
	IsCommitted                       *bool
	IsDependency                      *bool
	IsFrameworkFile                   *bool
	Manifest                          *string
	Name                              *string
	ODataType                         *string
	Size                              *int64
	SizeEncrypted                     *int64
}

func (p MobileAppContentFileOperationPredicate) Matches(input MobileAppContentFile) bool {

	if p.AzureStorageUri != nil && (input.AzureStorageUri == nil || *p.AzureStorageUri != *input.AzureStorageUri) {
		return false
	}

	if p.AzureStorageUriExpirationDateTime != nil && (input.AzureStorageUriExpirationDateTime == nil || *p.AzureStorageUriExpirationDateTime != *input.AzureStorageUriExpirationDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsCommitted != nil && (input.IsCommitted == nil || *p.IsCommitted != *input.IsCommitted) {
		return false
	}

	if p.IsDependency != nil && (input.IsDependency == nil || *p.IsDependency != *input.IsDependency) {
		return false
	}

	if p.IsFrameworkFile != nil && (input.IsFrameworkFile == nil || *p.IsFrameworkFile != *input.IsFrameworkFile) {
		return false
	}

	if p.Manifest != nil && (input.Manifest == nil || *p.Manifest != *input.Manifest) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	if p.SizeEncrypted != nil && (input.SizeEncrypted == nil || *p.SizeEncrypted != *input.SizeEncrypted) {
		return false
	}

	return true
}
