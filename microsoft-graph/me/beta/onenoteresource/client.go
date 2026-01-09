package onenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewOnenoteResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteResourceClient: %+v", err)
	}

	return &OnenoteResourceClient{
		Client: client,
	}, nil
}
