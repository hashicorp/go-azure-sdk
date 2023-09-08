package meonenotesection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionClient: %+v", err)
	}

	return &MeOnenoteSectionClient{
		Client: client,
	}, nil
}
