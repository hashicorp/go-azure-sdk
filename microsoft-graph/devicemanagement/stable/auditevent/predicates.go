package auditevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AuditEventOperationPredicate struct {
}

func (p AuditEventOperationPredicate) Matches(input stable.AuditEvent) bool {

	return true
}
