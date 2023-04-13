package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationOperationPredicate struct {
	Id          *string
	OperationId *string
}

func (p DeploymentOperationOperationPredicate) Matches(input DeploymentOperation) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.OperationId != nil && (input.OperationId == nil && *p.OperationId != *input.OperationId) {
		return false
	}

	return true
}
