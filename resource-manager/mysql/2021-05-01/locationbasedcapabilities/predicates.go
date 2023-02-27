// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationbasedcapabilities

type CapabilityPropertiesOperationPredicate struct {
}

func (p CapabilityPropertiesOperationPredicate) Matches(input CapabilityProperties) bool {

	return true
}
