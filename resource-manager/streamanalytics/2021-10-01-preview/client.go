// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_10_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/functions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/inputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/outputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/streamingjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/transformations"
)

type Client struct {
	Functions       *functions.FunctionsClient
	Inputs          *inputs.InputsClient
	Outputs         *outputs.OutputsClient
	StreamingJobs   *streamingjobs.StreamingJobsClient
	Subscriptions   *subscriptions.SubscriptionsClient
	Transformations *transformations.TransformationsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	functionsClient := functions.NewFunctionsClientWithBaseURI(endpoint)
	configureAuthFunc(&functionsClient.Client)

	inputsClient := inputs.NewInputsClientWithBaseURI(endpoint)
	configureAuthFunc(&inputsClient.Client)

	outputsClient := outputs.NewOutputsClientWithBaseURI(endpoint)
	configureAuthFunc(&outputsClient.Client)

	streamingJobsClient := streamingjobs.NewStreamingJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&streamingJobsClient.Client)

	subscriptionsClient := subscriptions.NewSubscriptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionsClient.Client)

	transformationsClient := transformations.NewTransformationsClientWithBaseURI(endpoint)
	configureAuthFunc(&transformationsClient.Client)

	return Client{
		Functions:       &functionsClient,
		Inputs:          &inputsClient,
		Outputs:         &outputsClient,
		StreamingJobs:   &streamingJobsClient,
		Subscriptions:   &subscriptionsClient,
		Transformations: &transformationsClient,
	}
}
