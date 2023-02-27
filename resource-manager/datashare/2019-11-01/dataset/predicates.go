// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataset

type DataSetOperationPredicate struct {
}

func (p DataSetOperationPredicate) Matches(input DataSet) bool {

	return true
}
