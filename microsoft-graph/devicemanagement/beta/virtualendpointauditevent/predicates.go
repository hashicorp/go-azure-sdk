package virtualendpointauditevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CloudPCAuditEventOperationPredicate struct {
}

func (p CloudPCAuditEventOperationPredicate) Matches(input beta.CloudPCAuditEvent) bool {

	return true
}
