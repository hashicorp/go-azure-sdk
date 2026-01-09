package inboundshareduserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundSharedUserProfileClient struct {
	Client *msgraph.Client
}

func NewInboundSharedUserProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*InboundSharedUserProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "inboundshareduserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InboundSharedUserProfileClient: %+v", err)
	}

	return &InboundSharedUserProfileClient{
		Client: client,
	}, nil
}
