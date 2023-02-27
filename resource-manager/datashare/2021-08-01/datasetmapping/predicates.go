// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasetmapping

type DataSetMappingOperationPredicate struct {
}

func (p DataSetMappingOperationPredicate) Matches(input DataSetMapping) bool {

	return true
}
