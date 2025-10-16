package v2025_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnector"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnectortemplate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/broker"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/brokerauthentication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/brokerauthorization"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/brokerlistener"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflow"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflowendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflowgraph"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflowprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/instance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/registryendpoint"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AkriConnector         *akriconnector.AkriConnectorClient
	AkriConnectorTemplate *akriconnectortemplate.AkriConnectorTemplateClient
	Broker                *broker.BrokerClient
	BrokerAuthentication  *brokerauthentication.BrokerAuthenticationClient
	BrokerAuthorization   *brokerauthorization.BrokerAuthorizationClient
	BrokerListener        *brokerlistener.BrokerListenerClient
	Dataflow              *dataflow.DataflowClient
	DataflowEndpoint      *dataflowendpoint.DataflowEndpointClient
	DataflowGraph         *dataflowgraph.DataflowGraphClient
	DataflowProfile       *dataflowprofile.DataflowProfileClient
	Instance              *instance.InstanceClient
	RegistryEndpoint      *registryendpoint.RegistryEndpointClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	akriConnectorClient, err := akriconnector.NewAkriConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AkriConnector client: %+v", err)
	}
	configureFunc(akriConnectorClient.Client)

	akriConnectorTemplateClient, err := akriconnectortemplate.NewAkriConnectorTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AkriConnectorTemplate client: %+v", err)
	}
	configureFunc(akriConnectorTemplateClient.Client)

	brokerAuthenticationClient, err := brokerauthentication.NewBrokerAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BrokerAuthentication client: %+v", err)
	}
	configureFunc(brokerAuthenticationClient.Client)

	brokerAuthorizationClient, err := brokerauthorization.NewBrokerAuthorizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BrokerAuthorization client: %+v", err)
	}
	configureFunc(brokerAuthorizationClient.Client)

	brokerClient, err := broker.NewBrokerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Broker client: %+v", err)
	}
	configureFunc(brokerClient.Client)

	brokerListenerClient, err := brokerlistener.NewBrokerListenerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BrokerListener client: %+v", err)
	}
	configureFunc(brokerListenerClient.Client)

	dataflowClient, err := dataflow.NewDataflowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dataflow client: %+v", err)
	}
	configureFunc(dataflowClient.Client)

	dataflowEndpointClient, err := dataflowendpoint.NewDataflowEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataflowEndpoint client: %+v", err)
	}
	configureFunc(dataflowEndpointClient.Client)

	dataflowGraphClient, err := dataflowgraph.NewDataflowGraphClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataflowGraph client: %+v", err)
	}
	configureFunc(dataflowGraphClient.Client)

	dataflowProfileClient, err := dataflowprofile.NewDataflowProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataflowProfile client: %+v", err)
	}
	configureFunc(dataflowProfileClient.Client)

	instanceClient, err := instance.NewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Instance client: %+v", err)
	}
	configureFunc(instanceClient.Client)

	registryEndpointClient, err := registryendpoint.NewRegistryEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegistryEndpoint client: %+v", err)
	}
	configureFunc(registryEndpointClient.Client)

	return &Client{
		AkriConnector:         akriConnectorClient,
		AkriConnectorTemplate: akriConnectorTemplateClient,
		Broker:                brokerClient,
		BrokerAuthentication:  brokerAuthenticationClient,
		BrokerAuthorization:   brokerAuthorizationClient,
		BrokerListener:        brokerListenerClient,
		Dataflow:              dataflowClient,
		DataflowEndpoint:      dataflowEndpointClient,
		DataflowGraph:         dataflowGraphClient,
		DataflowProfile:       dataflowProfileClient,
		Instance:              instanceClient,
		RegistryEndpoint:      registryEndpointClient,
	}, nil
}
