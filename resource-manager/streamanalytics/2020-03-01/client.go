package v2020_03_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/functions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/inputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/outputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/streamingjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/transformations"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	functionsClient := functions.NewFunctionsClientWithBaseURI(endpoint)
	configureAuthFunc(&functionsClient.Client)

	inputsClient := inputs.NewInputsClientWithBaseURI(endpoint)
	configureAuthFunc(&inputsClient.Client)

	outputsClient := outputs.NewOutputsClientWithBaseURI(endpoint)
	configureAuthFunc(&outputsClient.Client)

	privateEndpointsClient := privateendpoints.NewPrivateEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointsClient.Client)

	streamingJobsClient := streamingjobs.NewStreamingJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&streamingJobsClient.Client)

	subscriptionsClient := subscriptions.NewSubscriptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionsClient.Client)

	transformationsClient := transformations.NewTransformationsClientWithBaseURI(endpoint)
	configureAuthFunc(&transformationsClient.Client)

	return Client{
		Clusters:         &clustersClient,
		Functions:        &functionsClient,
		Inputs:           &inputsClient,
		Outputs:          &outputsClient,
		PrivateEndpoints: &privateEndpointsClient,
		StreamingJobs:    &streamingJobsClient,
		Subscriptions:    &subscriptionsClient,
		Transformations:  &transformationsClient,
	}
}
