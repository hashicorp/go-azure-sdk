package grouponenotesectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionPageClient: %+v", err)
	}

	return &GroupOnenoteSectionPageClient{
		Client: client,
	}, nil
}
