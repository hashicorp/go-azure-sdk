package addons

type AddonOperationPredicate struct {
}

func (p AddonOperationPredicate) Matches(input Addon) bool {

	return true
}
