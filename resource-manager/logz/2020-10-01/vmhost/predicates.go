package vmhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMResourcesOperationPredicate struct {
	AgentVersion *string
	Id           *string
}

func (p VMResourcesOperationPredicate) Matches(input VMResources) bool {

	if p.AgentVersion != nil && (input.AgentVersion == nil || *p.AgentVersion != *input.AgentVersion) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}
