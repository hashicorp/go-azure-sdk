package replicationappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationApplianceOperationPredicate struct {
}

func (p ReplicationApplianceOperationPredicate) Matches(input ReplicationAppliance) bool {

	return true
}
