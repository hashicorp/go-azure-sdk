package savingsplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanValidateResponse struct {
	Benefits *[]SavingsPlanValidResponseProperty `json:"benefits,omitempty"`
	NextLink *string                             `json:"nextLink,omitempty"`
}
