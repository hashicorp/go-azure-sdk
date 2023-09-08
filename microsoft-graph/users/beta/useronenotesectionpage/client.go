package useronenotesectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionPageClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionPageClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionPageClient: %+v", err)
	}

	return &UserOnenoteSectionPageClient{
		Client: client,
	}, nil
}
