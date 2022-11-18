package replicationappliances

type ReplicationApplianceOperationPredicate struct {
}

func (p ReplicationApplianceOperationPredicate) Matches(input ReplicationAppliance) bool {

	return true
}
