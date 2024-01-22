package transfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransferDetailsOperationPredicate struct {
}

func (p TransferDetailsOperationPredicate) Matches(input TransferDetails) bool {

	return true
}
