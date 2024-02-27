package listserviceproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProviderOperationPredicate struct {
}

func (p ServiceProviderOperationPredicate) Matches(input ServiceProvider) bool {

	return true
}
