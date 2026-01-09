package outboundshareduserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundSharedUserProfileClient struct {
	Client *msgraph.Client
}

func NewOutboundSharedUserProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*OutboundSharedUserProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outboundshareduserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutboundSharedUserProfileClient: %+v", err)
	}

	return &OutboundSharedUserProfileClient{
		Client: client,
	}, nil
}
