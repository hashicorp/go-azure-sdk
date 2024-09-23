package windowsautopilotdeploymentprofile

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type HasPayloadLinkResultItemOperationPredicate struct {
}

func (p HasPayloadLinkResultItemOperationPredicate) Matches(input beta.HasPayloadLinkResultItem) bool {

	return true
}

type WindowsAutopilotDeploymentProfileOperationPredicate struct {
}

func (p WindowsAutopilotDeploymentProfileOperationPredicate) Matches(input beta.WindowsAutopilotDeploymentProfile) bool {

	return true
}
