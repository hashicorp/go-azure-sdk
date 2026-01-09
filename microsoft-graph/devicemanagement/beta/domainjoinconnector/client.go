package domainjoinconnector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainJoinConnectorClient struct {
	Client *msgraph.Client
}

func NewDomainJoinConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*DomainJoinConnectorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "domainjoinconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainJoinConnectorClient: %+v", err)
	}

	return &DomainJoinConnectorClient{
		Client: client,
	}, nil
}
