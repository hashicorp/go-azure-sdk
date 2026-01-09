package onenotesectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupClient: %+v", err)
	}

	return &OnenoteSectionGroupClient{
		Client: client,
	}, nil
}
