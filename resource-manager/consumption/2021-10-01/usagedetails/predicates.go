// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package usagedetails

type UsageDetailOperationPredicate struct {
}

func (p UsageDetailOperationPredicate) Matches(input UsageDetail) bool {

	return true
}
