package v2015_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/analyticsitemsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentannotationsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentapikeysapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentcontinuousexportapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentfeaturesandpricingapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentproactivedetectionapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentworkitemconfigsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/favoritesapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/myworkbooksapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/webtestlocationsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/webtestsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/workbooksapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AnalyticsItemsAPIs              *analyticsitemsapis.AnalyticsItemsAPIsClient
	ComponentAnnotationsAPIs        *componentannotationsapis.ComponentAnnotationsAPIsClient
	ComponentApiKeysAPIs            *componentapikeysapis.ComponentApiKeysAPIsClient
	ComponentContinuousExportAPIs   *componentcontinuousexportapis.ComponentContinuousExportAPIsClient
	ComponentFeaturesAndPricingAPIs *componentfeaturesandpricingapis.ComponentFeaturesAndPricingAPIsClient
	ComponentProactiveDetectionAPIs *componentproactivedetectionapis.ComponentProactiveDetectionAPIsClient
	ComponentWorkItemConfigsAPIs    *componentworkitemconfigsapis.ComponentWorkItemConfigsAPIsClient
	ComponentsAPIs                  *componentsapis.ComponentsAPIsClient
	FavoritesAPIs                   *favoritesapis.FavoritesAPIsClient
	MyworkbooksAPIs                 *myworkbooksapis.MyworkbooksAPIsClient
	WebTestLocationsAPIs            *webtestlocationsapis.WebTestLocationsAPIsClient
	WebTestsAPIs                    *webtestsapis.WebTestsAPIsClient
	WorkbooksAPIs                   *workbooksapis.WorkbooksAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	analyticsItemsAPIsClient, err := analyticsitemsapis.NewAnalyticsItemsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AnalyticsItemsAPIs client: %+v", err)
	}
	configureFunc(analyticsItemsAPIsClient.Client)

	componentAnnotationsAPIsClient, err := componentannotationsapis.NewComponentAnnotationsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentAnnotationsAPIs client: %+v", err)
	}
	configureFunc(componentAnnotationsAPIsClient.Client)

	componentApiKeysAPIsClient, err := componentapikeysapis.NewComponentApiKeysAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentApiKeysAPIs client: %+v", err)
	}
	configureFunc(componentApiKeysAPIsClient.Client)

	componentContinuousExportAPIsClient, err := componentcontinuousexportapis.NewComponentContinuousExportAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentContinuousExportAPIs client: %+v", err)
	}
	configureFunc(componentContinuousExportAPIsClient.Client)

	componentFeaturesAndPricingAPIsClient, err := componentfeaturesandpricingapis.NewComponentFeaturesAndPricingAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentFeaturesAndPricingAPIs client: %+v", err)
	}
	configureFunc(componentFeaturesAndPricingAPIsClient.Client)

	componentProactiveDetectionAPIsClient, err := componentproactivedetectionapis.NewComponentProactiveDetectionAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentProactiveDetectionAPIs client: %+v", err)
	}
	configureFunc(componentProactiveDetectionAPIsClient.Client)

	componentWorkItemConfigsAPIsClient, err := componentworkitemconfigsapis.NewComponentWorkItemConfigsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentWorkItemConfigsAPIs client: %+v", err)
	}
	configureFunc(componentWorkItemConfigsAPIsClient.Client)

	componentsAPIsClient, err := componentsapis.NewComponentsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentsAPIs client: %+v", err)
	}
	configureFunc(componentsAPIsClient.Client)

	favoritesAPIsClient, err := favoritesapis.NewFavoritesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FavoritesAPIs client: %+v", err)
	}
	configureFunc(favoritesAPIsClient.Client)

	myworkbooksAPIsClient, err := myworkbooksapis.NewMyworkbooksAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MyworkbooksAPIs client: %+v", err)
	}
	configureFunc(myworkbooksAPIsClient.Client)

	webTestLocationsAPIsClient, err := webtestlocationsapis.NewWebTestLocationsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebTestLocationsAPIs client: %+v", err)
	}
	configureFunc(webTestLocationsAPIsClient.Client)

	webTestsAPIsClient, err := webtestsapis.NewWebTestsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebTestsAPIs client: %+v", err)
	}
	configureFunc(webTestsAPIsClient.Client)

	workbooksAPIsClient, err := workbooksapis.NewWorkbooksAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkbooksAPIs client: %+v", err)
	}
	configureFunc(workbooksAPIsClient.Client)

	return &Client{
		AnalyticsItemsAPIs:              analyticsItemsAPIsClient,
		ComponentAnnotationsAPIs:        componentAnnotationsAPIsClient,
		ComponentApiKeysAPIs:            componentApiKeysAPIsClient,
		ComponentContinuousExportAPIs:   componentContinuousExportAPIsClient,
		ComponentFeaturesAndPricingAPIs: componentFeaturesAndPricingAPIsClient,
		ComponentProactiveDetectionAPIs: componentProactiveDetectionAPIsClient,
		ComponentWorkItemConfigsAPIs:    componentWorkItemConfigsAPIsClient,
		ComponentsAPIs:                  componentsAPIsClient,
		FavoritesAPIs:                   favoritesAPIsClient,
		MyworkbooksAPIs:                 myworkbooksAPIsClient,
		WebTestLocationsAPIs:            webTestLocationsAPIsClient,
		WebTestsAPIs:                    webTestsAPIsClient,
		WorkbooksAPIs:                   workbooksAPIsClient,
	}, nil
}
