package billingpermission

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPermissionOperationPredicate struct {
}

func (p BillingPermissionOperationPredicate) Matches(input BillingPermission) bool {

	return true
}
