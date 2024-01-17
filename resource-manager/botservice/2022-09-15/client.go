package v2022_09_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bot"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/botconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/bothostsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/channel"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listqnamakerendpointkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listserviceproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Bot                        *bot.BotClient
	BotConnection              *botconnection.BotConnectionClient
	BotHostSettings            *bothostsettings.BotHostSettingsClient
	Channel                    *channel.ChannelClient
	ListQnAMakerEndpointKeys   *listqnamakerendpointkeys.ListQnAMakerEndpointKeysClient
	ListServiceProviders       *listserviceproviders.ListServiceProvidersClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	botClient, err := bot.NewBotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Bot client: %+v", err)
	}
	configureFunc(botClient.Client)

	botConnectionClient, err := botconnection.NewBotConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BotConnection client: %+v", err)
	}
	configureFunc(botConnectionClient.Client)

	botHostSettingsClient, err := bothostsettings.NewBotHostSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BotHostSettings client: %+v", err)
	}
	configureFunc(botHostSettingsClient.Client)

	channelClient, err := channel.NewChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Channel client: %+v", err)
	}
	configureFunc(channelClient.Client)

	listQnAMakerEndpointKeysClient, err := listqnamakerendpointkeys.NewListQnAMakerEndpointKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListQnAMakerEndpointKeys client: %+v", err)
	}
	configureFunc(listQnAMakerEndpointKeysClient.Client)

	listServiceProvidersClient, err := listserviceproviders.NewListServiceProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListServiceProviders client: %+v", err)
	}
	configureFunc(listServiceProvidersClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	return &Client{
		Bot:                        botClient,
		BotConnection:              botConnectionClient,
		BotHostSettings:            botHostSettingsClient,
		Channel:                    channelClient,
		ListQnAMakerEndpointKeys:   listQnAMakerEndpointKeysClient,
		ListServiceProviders:       listServiceProvidersClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
	}, nil
}
