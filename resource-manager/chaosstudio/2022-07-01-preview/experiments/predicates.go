// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package experiments

type ExperimentOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ExperimentOperationPredicate) Matches(input Experiment) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
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

type ExperimentExecutionDetailsOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ExperimentExecutionDetailsOperationPredicate) Matches(input ExperimentExecutionDetails) bool {

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

type ExperimentStatusOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ExperimentStatusOperationPredicate) Matches(input ExperimentStatus) bool {

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
