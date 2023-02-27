// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apirevision

type ApiRevisionContractOperationPredicate struct {
	ApiId           *string
	ApiRevision     *string
	CreatedDateTime *string
	Description     *string
	IsCurrent       *bool
	IsOnline        *bool
	PrivateUrl      *string
	UpdatedDateTime *string
}

func (p ApiRevisionContractOperationPredicate) Matches(input ApiRevisionContract) bool {

	if p.ApiId != nil && (input.ApiId == nil && *p.ApiId != *input.ApiId) {
		return false
	}

	if p.ApiRevision != nil && (input.ApiRevision == nil && *p.ApiRevision != *input.ApiRevision) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil && *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil && *p.Description != *input.Description) {
		return false
	}

	if p.IsCurrent != nil && (input.IsCurrent == nil && *p.IsCurrent != *input.IsCurrent) {
		return false
	}

	if p.IsOnline != nil && (input.IsOnline == nil && *p.IsOnline != *input.IsOnline) {
		return false
	}

	if p.PrivateUrl != nil && (input.PrivateUrl == nil && *p.PrivateUrl != *input.PrivateUrl) {
		return false
	}

	if p.UpdatedDateTime != nil && (input.UpdatedDateTime == nil && *p.UpdatedDateTime != *input.UpdatedDateTime) {
		return false
	}

	return true
}
