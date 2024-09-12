package eventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceClient struct {
	Client *msgraph.Client
}

func NewEventInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceClient: %+v", err)
	}

	return &EventInstanceClient{
		Client: client,
	}, nil
}
