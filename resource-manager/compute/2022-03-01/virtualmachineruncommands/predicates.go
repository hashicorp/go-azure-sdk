// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommands

type RunCommandDocumentBaseOperationPredicate struct {
	Description *string
	Id          *string
	Label       *string
	Schema      *string
}

func (p RunCommandDocumentBaseOperationPredicate) Matches(input RunCommandDocumentBase) bool {

	if p.Description != nil && *p.Description != input.Description {
		return false
	}

	if p.Id != nil && *p.Id != input.Id {
		return false
	}

	if p.Label != nil && *p.Label != input.Label {
		return false
	}

	if p.Schema != nil && *p.Schema != input.Schema {
		return false
	}

	return true
}

type VirtualMachineRunCommandOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p VirtualMachineRunCommandOperationPredicate) Matches(input VirtualMachineRunCommand) bool {

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
