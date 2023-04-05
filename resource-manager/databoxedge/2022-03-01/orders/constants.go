package orders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrderState string

const (
	OrderStateArriving               OrderState = "Arriving"
	OrderStateAwaitingDrop           OrderState = "AwaitingDrop"
	OrderStateAwaitingFulfillment    OrderState = "AwaitingFulfillment"
	OrderStateAwaitingPickup         OrderState = "AwaitingPickup"
	OrderStateAwaitingPreparation    OrderState = "AwaitingPreparation"
	OrderStateAwaitingReturnShipment OrderState = "AwaitingReturnShipment"
	OrderStateAwaitingShipment       OrderState = "AwaitingShipment"
	OrderStateCollectedAtMicrosoft   OrderState = "CollectedAtMicrosoft"
	OrderStateDeclined               OrderState = "Declined"
	OrderStateDelivered              OrderState = "Delivered"
	OrderStateLostDevice             OrderState = "LostDevice"
	OrderStatePickupCompleted        OrderState = "PickupCompleted"
	OrderStateReplacementRequested   OrderState = "ReplacementRequested"
	OrderStateReturnInitiated        OrderState = "ReturnInitiated"
	OrderStateShipped                OrderState = "Shipped"
	OrderStateShippedBack            OrderState = "ShippedBack"
	OrderStateUntracked              OrderState = "Untracked"
)

func PossibleValuesForOrderState() []string {
	return []string{
		string(OrderStateArriving),
		string(OrderStateAwaitingDrop),
		string(OrderStateAwaitingFulfillment),
		string(OrderStateAwaitingPickup),
		string(OrderStateAwaitingPreparation),
		string(OrderStateAwaitingReturnShipment),
		string(OrderStateAwaitingShipment),
		string(OrderStateCollectedAtMicrosoft),
		string(OrderStateDeclined),
		string(OrderStateDelivered),
		string(OrderStateLostDevice),
		string(OrderStatePickupCompleted),
		string(OrderStateReplacementRequested),
		string(OrderStateReturnInitiated),
		string(OrderStateShipped),
		string(OrderStateShippedBack),
		string(OrderStateUntracked),
	}
}

type ShipmentType string

const (
	ShipmentTypeNotApplicable     ShipmentType = "NotApplicable"
	ShipmentTypeSelfPickup        ShipmentType = "SelfPickup"
	ShipmentTypeShippedToCustomer ShipmentType = "ShippedToCustomer"
)

func PossibleValuesForShipmentType() []string {
	return []string{
		string(ShipmentTypeNotApplicable),
		string(ShipmentTypeSelfPickup),
		string(ShipmentTypeShippedToCustomer),
	}
}
