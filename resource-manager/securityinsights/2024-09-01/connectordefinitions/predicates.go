package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorDefinitionOperationPredicate struct {
}

func (p DataConnectorDefinitionOperationPredicate) Matches(input DataConnectorDefinition) bool {

	return true
}
