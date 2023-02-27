// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentscripts

type DeploymentScriptOperationPredicate struct {
}

func (p DeploymentScriptOperationPredicate) Matches(input DeploymentScript) bool {

	return true
}
