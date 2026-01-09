package onenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteClient struct {
	Client *msgraph.Client
}

func NewOnenoteClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteClient: %+v", err)
	}

	return &OnenoteClient{
		Client: client,
	}, nil
}
