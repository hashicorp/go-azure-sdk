package teamscheduleopenshift

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type OpenShiftOperationPredicate struct {
}

func (p OpenShiftOperationPredicate) Matches(input stable.OpenShift) bool {

	return true
}
