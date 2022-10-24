package dataconnectors

type DataConnectorOperationPredicate struct {
}

func (p DataConnectorOperationPredicate) Matches(input DataConnector) bool {

	return true
}
