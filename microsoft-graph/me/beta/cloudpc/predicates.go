package cloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CloudPCOperationPredicate struct {
}

func (p CloudPCOperationPredicate) Matches(input beta.CloudPC) bool {

	return true
}

type CloudPCRemoteActionResultOperationPredicate struct {
}

func (p CloudPCRemoteActionResultOperationPredicate) Matches(input beta.CloudPCRemoteActionResult) bool {

	return true
}

type CloudPCResizeValidationResultOperationPredicate struct {
}

func (p CloudPCResizeValidationResultOperationPredicate) Matches(input beta.CloudPCResizeValidationResult) bool {

	return true
}
