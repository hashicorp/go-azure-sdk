package directoryroletemplate

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input beta.DirectoryObject) bool {

	return true
}

type DirectoryRoleTemplateOperationPredicate struct {
}

func (p DirectoryRoleTemplateOperationPredicate) Matches(input beta.DirectoryRoleTemplate) bool {

	return true
}
