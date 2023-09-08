package meonenotepagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenotePageContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenotePageContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenotePageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotepagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenotePageContentClient: %+v", err)
	}

	return &MeOnenotePageContentClient{
		Client: client,
	}, nil
}
