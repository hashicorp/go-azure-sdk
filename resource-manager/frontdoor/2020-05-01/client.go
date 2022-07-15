package v2020_05_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailabilitywithsubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/frontdoors"
)

type Client struct {
	CheckFrontDoorNameAvailability                 *checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient
	CheckFrontDoorNameAvailabilityWithSubscription *checkfrontdoornameavailabilitywithsubscription.CheckFrontDoorNameAvailabilityWithSubscriptionClient
	FrontDoors                                     *frontdoors.FrontDoorsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkFrontDoorNameAvailabilityClient := checkfrontdoornameavailability.NewCheckFrontDoorNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityClient.Client)

	checkFrontDoorNameAvailabilityWithSubscriptionClient := checkfrontdoornameavailabilitywithsubscription.NewCheckFrontDoorNameAvailabilityWithSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityWithSubscriptionClient.Client)

	frontDoorsClient := frontdoors.NewFrontDoorsClientWithBaseURI(endpoint)
	configureAuthFunc(&frontDoorsClient.Client)

	return Client{
		CheckFrontDoorNameAvailability:                 &checkFrontDoorNameAvailabilityClient,
		CheckFrontDoorNameAvailabilityWithSubscription: &checkFrontDoorNameAvailabilityWithSubscriptionClient,
		FrontDoors: &frontDoorsClient,
	}
}
