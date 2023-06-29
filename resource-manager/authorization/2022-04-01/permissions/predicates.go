package permissions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionOperationPredicate struct {
}

func (p PermissionOperationPredicate) Matches(input Permission) bool {

	return true
}
