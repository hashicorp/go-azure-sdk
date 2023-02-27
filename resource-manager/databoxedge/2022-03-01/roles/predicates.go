package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleOperationPredicate struct {
}

func (p RoleOperationPredicate) Matches(input Role) bool {

	return true
}
