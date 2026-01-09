package remoteassistancepartner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistancePartnerClient struct {
	Client *msgraph.Client
}

func NewRemoteAssistancePartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*RemoteAssistancePartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "remoteassistancepartner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RemoteAssistancePartnerClient: %+v", err)
	}

	return &RemoteAssistancePartnerClient{
		Client: client,
	}, nil
}
