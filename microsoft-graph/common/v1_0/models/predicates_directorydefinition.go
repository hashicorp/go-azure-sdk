package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryDefinitionOperationPredicate struct {
	DiscoveryDateTime *string
	Id                *string
	Name              *string
	ODataType         *string
	ReadOnly          *bool
	Version           *string
}

func (p DirectoryDefinitionOperationPredicate) Matches(input DirectoryDefinition) bool {

	if p.DiscoveryDateTime != nil && (input.DiscoveryDateTime == nil || *p.DiscoveryDateTime != *input.DiscoveryDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReadOnly != nil && (input.ReadOnly == nil || *p.ReadOnly != *input.ReadOnly) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
