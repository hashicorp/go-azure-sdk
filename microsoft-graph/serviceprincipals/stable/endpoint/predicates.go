package endpoint

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type EndpointOperationPredicate struct {
}

func (p EndpointOperationPredicate) Matches(input stable.Endpoint) bool {

	return true
}
