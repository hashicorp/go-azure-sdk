package sqlpoolsblobauditing

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedSqlPoolBlobAuditingPolicyOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ExtendedSqlPoolBlobAuditingPolicyOperationPredicate) Matches(input ExtendedSqlPoolBlobAuditingPolicy) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type SqlPoolBlobAuditingPolicyOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p SqlPoolBlobAuditingPolicyOperationPredicate) Matches(input SqlPoolBlobAuditingPolicy) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
