package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantUserSyncInboundOperationPredicate struct {
	IsSyncAllowed *bool
	ODataType     *string
}

func (p CrossTenantUserSyncInboundOperationPredicate) Matches(input CrossTenantUserSyncInbound) bool {

	if p.IsSyncAllowed != nil && (input.IsSyncAllowed == nil || *p.IsSyncAllowed != *input.IsSyncAllowed) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
