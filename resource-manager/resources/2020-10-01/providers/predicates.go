package providers

type ProviderOperationPredicate struct {
	Id                 *string
	Namespace          *string
	RegistrationPolicy *string
	RegistrationState  *string
}

func (p ProviderOperationPredicate) Matches(input Provider) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Namespace != nil && (input.Namespace == nil && *p.Namespace != *input.Namespace) {
		return false
	}

	if p.RegistrationPolicy != nil && (input.RegistrationPolicy == nil && *p.RegistrationPolicy != *input.RegistrationPolicy) {
		return false
	}

	if p.RegistrationState != nil && (input.RegistrationState == nil && *p.RegistrationState != *input.RegistrationState) {
		return false
	}

	return true
}
