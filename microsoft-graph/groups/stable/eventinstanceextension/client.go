package eventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewEventInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceExtensionClient: %+v", err)
	}

	return &EventInstanceExtensionClient{
		Client: client,
	}, nil
}
