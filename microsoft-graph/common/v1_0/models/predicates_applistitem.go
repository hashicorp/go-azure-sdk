package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppListItemOperationPredicate struct {
	AppId       *string
	AppStoreUrl *string
	Name        *string
	ODataType   *string
	Publisher   *string
}

func (p AppListItemOperationPredicate) Matches(input AppListItem) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.AppStoreUrl != nil && (input.AppStoreUrl == nil || *p.AppStoreUrl != *input.AppStoreUrl) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	return true
}
