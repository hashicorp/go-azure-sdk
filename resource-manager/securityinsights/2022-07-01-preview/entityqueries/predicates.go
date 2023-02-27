package entityqueries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
