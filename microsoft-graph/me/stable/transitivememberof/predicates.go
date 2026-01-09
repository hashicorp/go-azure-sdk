package transitivememberof

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AdministrativeUnitOperationPredicate struct {
}

func (p AdministrativeUnitOperationPredicate) Matches(input stable.AdministrativeUnit) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type DirectoryRoleOperationPredicate struct {
}

func (p DirectoryRoleOperationPredicate) Matches(input stable.DirectoryRole) bool {

	return true
}

type GroupOperationPredicate struct {
}

func (p GroupOperationPredicate) Matches(input stable.Group) bool {

	return true
}
