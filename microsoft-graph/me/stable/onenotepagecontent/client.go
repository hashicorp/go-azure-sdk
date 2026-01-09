package onenotepagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePageContentClient struct {
	Client *msgraph.Client
}

func NewOnenotePageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenotePageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotepagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenotePageContentClient: %+v", err)
	}

	return &OnenotePageContentClient{
		Client: client,
	}, nil
}
