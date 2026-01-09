package profilepublication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfilePublicationClient struct {
	Client *msgraph.Client
}

func NewProfilePublicationClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfilePublicationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilepublication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfilePublicationClient: %+v", err)
	}

	return &ProfilePublicationClient{
		Client: client,
	}, nil
}
