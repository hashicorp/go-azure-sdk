package loadtest

type OutboundEnvironmentEndpointOperationPredicate struct {
	Category *string
}

func (p OutboundEnvironmentEndpointOperationPredicate) Matches(input OutboundEnvironmentEndpoint) bool {

	if p.Category != nil && (input.Category == nil && *p.Category != *input.Category) {
		return false
	}

	return true
}
