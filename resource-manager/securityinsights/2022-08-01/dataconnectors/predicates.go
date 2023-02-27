// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataconnectors

type DataConnectorOperationPredicate struct {
}

func (p DataConnectorOperationPredicate) Matches(input DataConnector) bool {

	return true
}
