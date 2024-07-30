package microsofttunnelserverlogcollectionrespons

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerLogCollectionResponsClient struct {
	Client *msgraph.Client
}

func NewMicrosoftTunnelServerLogCollectionResponsClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftTunnelServerLogCollectionResponsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "microsofttunnelserverlogcollectionrespons", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftTunnelServerLogCollectionResponsClient: %+v", err)
	}

	return &MicrosoftTunnelServerLogCollectionResponsClient{
		Client: client,
	}, nil
}
