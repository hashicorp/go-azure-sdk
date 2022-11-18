package settings

type SettingOperationPredicate struct {
}

func (p SettingOperationPredicate) Matches(input Setting) bool {

	return true
}
