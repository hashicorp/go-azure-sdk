package memberof

// Copyright (c) HashiCorp Inc. All rights reserved.
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

type GroupOperationPredicate struct {
}

func (p GroupOperationPredicate) Matches(input stable.Group) bool {

	return true
}
