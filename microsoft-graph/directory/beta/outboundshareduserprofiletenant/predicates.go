package outboundshareduserprofiletenant

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type TenantReferenceOperationPredicate struct {
}

func (p TenantReferenceOperationPredicate) Matches(input beta.TenantReference) bool {

	return true
}
