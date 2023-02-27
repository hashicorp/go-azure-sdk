package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetMappingOperationPredicate struct {
}

func (p DataSetMappingOperationPredicate) Matches(input DataSetMapping) bool {

	return true
}
