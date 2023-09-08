package meonenotesectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupClient: %+v", err)
	}

	return &MeOnenoteSectionGroupClient{
		Client: client,
	}, nil
}
