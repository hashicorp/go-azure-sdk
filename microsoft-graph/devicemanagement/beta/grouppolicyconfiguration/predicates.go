package grouppolicyconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type GroupPolicyConfigurationOperationPredicate struct {
}

func (p GroupPolicyConfigurationOperationPredicate) Matches(input beta.GroupPolicyConfiguration) bool {

	return true
}

type GroupPolicyConfigurationAssignmentOperationPredicate struct {
}

func (p GroupPolicyConfigurationAssignmentOperationPredicate) Matches(input beta.GroupPolicyConfigurationAssignment) bool {

	return true
}
