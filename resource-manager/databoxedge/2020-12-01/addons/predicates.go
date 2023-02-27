// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package addons

type AddonOperationPredicate struct {
}

func (p AddonOperationPredicate) Matches(input Addon) bool {

	return true
}
