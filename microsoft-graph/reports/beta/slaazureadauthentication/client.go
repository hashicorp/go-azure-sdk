package slaazureadauthentication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlaAzureADAuthenticationClient struct {
	Client *msgraph.Client
}

func NewSlaAzureADAuthenticationClientWithBaseURI(sdkApi sdkEnv.Api) (*SlaAzureADAuthenticationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "slaazureadauthentication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SlaAzureADAuthenticationClient: %+v", err)
	}

	return &SlaAzureADAuthenticationClient{
		Client: client,
	}, nil
}
