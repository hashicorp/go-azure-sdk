// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationappliances

type ReplicationApplianceOperationPredicate struct {
}

func (p ReplicationApplianceOperationPredicate) Matches(input ReplicationAppliance) bool {

	return true
}
