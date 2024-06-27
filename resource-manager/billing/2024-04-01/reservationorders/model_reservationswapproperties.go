package reservationorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationSwapProperties struct {
	SwapDestination *string `json:"swapDestination,omitempty"`
	SwapSource      *string `json:"swapSource,omitempty"`
}
