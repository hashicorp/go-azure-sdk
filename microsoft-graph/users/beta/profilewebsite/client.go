package profilewebsite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileWebsiteClient struct {
	Client *msgraph.Client
}

func NewProfileWebsiteClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileWebsiteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilewebsite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileWebsiteClient: %+v", err)
	}

	return &ProfileWebsiteClient{
		Client: client,
	}, nil
}
