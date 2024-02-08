package v2021_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/functions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/inputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/outputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/streamingjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/transformations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Functions       *functions.FunctionsClient
	Inputs          *inputs.InputsClient
	Outputs         *outputs.OutputsClient
	StreamingJobs   *streamingjobs.StreamingJobsClient
	Subscriptions   *subscriptions.SubscriptionsClient
	Transformations *transformations.TransformationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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
		Functions:       functionsClient,
		Inputs:          inputsClient,
		Outputs:         outputsClient,
		StreamingJobs:   streamingJobsClient,
		Subscriptions:   subscriptionsClient,
		Transformations: transformationsClient,
	}, nil
}
