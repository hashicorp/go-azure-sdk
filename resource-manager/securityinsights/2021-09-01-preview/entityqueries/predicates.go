package entityqueries

type EntityQueryOperationPredicate struct {
}

func (p EntityQueryOperationPredicate) Matches(input EntityQuery) bool {

	return true
}

type EntityQueryTemplateOperationPredicate struct {
}

func (p EntityQueryTemplateOperationPredicate) Matches(input EntityQueryTemplate) bool {

	return true
}
