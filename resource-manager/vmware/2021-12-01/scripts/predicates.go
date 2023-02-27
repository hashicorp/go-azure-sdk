// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scripts

type ScriptCmdletOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptCmdletOperationPredicate) Matches(input ScriptCmdlet) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type ScriptExecutionOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptExecutionOperationPredicate) Matches(input ScriptExecution) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type ScriptPackageOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptPackageOperationPredicate) Matches(input ScriptPackage) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
