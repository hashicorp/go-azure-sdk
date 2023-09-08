package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantExecutionErrorOperationPredicate struct {
	Error          *string
	ErrorDetails   *string
	NodeId         *int64
	ODataType      *string
	RawToken       *string
	StatementIndex *int64
	TenantId       *string
}

func (p ManagedTenantsManagedTenantExecutionErrorOperationPredicate) Matches(input ManagedTenantsManagedTenantExecutionError) bool {

	if p.Error != nil && (input.Error == nil || *p.Error != *input.Error) {
		return false
	}

	if p.ErrorDetails != nil && (input.ErrorDetails == nil || *p.ErrorDetails != *input.ErrorDetails) {
		return false
	}

	if p.NodeId != nil && (input.NodeId == nil || *p.NodeId != *input.NodeId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RawToken != nil && (input.RawToken == nil || *p.RawToken != *input.RawToken) {
		return false
	}

	if p.StatementIndex != nil && (input.StatementIndex == nil || *p.StatementIndex != *input.StatementIndex) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
