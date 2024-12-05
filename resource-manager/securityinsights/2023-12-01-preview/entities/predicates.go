package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityOperationPredicate struct {
}

func (p EntityOperationPredicate) Matches(input Entity) bool {

	return true
}
