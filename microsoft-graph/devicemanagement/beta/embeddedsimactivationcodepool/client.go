package embeddedsimactivationcodepool

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodePoolClient struct {
	Client *msgraph.Client
}

func NewEmbeddedSIMActivationCodePoolClientWithBaseURI(sdkApi sdkEnv.Api) (*EmbeddedSIMActivationCodePoolClient, error) {
	client, err := msgraph.NewClient(sdkApi, "embeddedsimactivationcodepool", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmbeddedSIMActivationCodePoolClient: %+v", err)
	}

	return &EmbeddedSIMActivationCodePoolClient{
		Client: client,
	}, nil
}
