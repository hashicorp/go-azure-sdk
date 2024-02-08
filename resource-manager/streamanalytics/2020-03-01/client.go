package v2020_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/functions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/inputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/outputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/streamingjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/transformations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Clusters         *clusters.ClustersClient
	Functions        *functions.FunctionsClient
	Inputs           *inputs.InputsClient
	Outputs          *outputs.OutputsClient
	PrivateEndpoints *privateendpoints.PrivateEndpointsClient
	StreamingJobs    *streamingjobs.StreamingJobsClient
	Subscriptions    *subscriptions.SubscriptionsClient
	Transformations  *transformations.TransformationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	functionsClient, err := functions.NewFunctionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Functions client: %+v", err)
	}
	configureFunc(functionsClient.Client)

	inputsClient, err := inputs.NewInputsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Inputs client: %+v", err)
	}
	configureFunc(inputsClient.Client)

	outputsClient, err := outputs.NewOutputsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Outputs client: %+v", err)
	}
	configureFunc(outputsClient.Client)

	privateEndpointsClient, err := privateendpoints.NewPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpoints client: %+v", err)
	}
	configureFunc(privateEndpointsClient.Client)

	streamingJobsClient, err := streamingjobs.NewStreamingJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StreamingJobs client: %+v", err)
	}
	configureFunc(streamingJobsClient.Client)

	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	transformationsClient, err := transformations.NewTransformationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transformations client: %+v", err)
	}
	configureFunc(transformationsClient.Client)

	return &Client{
		Clusters:         clustersClient,
		Functions:        functionsClient,
		Inputs:           inputsClient,
		Outputs:          outputsClient,
		PrivateEndpoints: privateEndpointsClient,
		StreamingJobs:    streamingJobsClient,
		Subscriptions:    subscriptionsClient,
		Transformations:  transformationsClient,
	}, nil
}
