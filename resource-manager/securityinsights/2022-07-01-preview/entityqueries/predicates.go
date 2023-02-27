// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package entityqueries

type EntityQueryOperationPredicate struct {
}

func (p EntityQueryOperationPredicate) Matches(input EntityQuery) bool {

	return true
}

type EntityQueryTemplateOperationPredicate struct {
}

func (p EntityQueryTemplateOperationPredicate) Matches(input EntityQueryTemplate) bool {

	return true
}
