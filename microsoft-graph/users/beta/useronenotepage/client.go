package useronenotepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenotePageClient struct {
	Client *msgraph.Client
}

func NewUserOnenotePageClientWithBaseURI(api sdkEnv.Api) (*UserOnenotePageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenotePageClient: %+v", err)
	}

	return &UserOnenotePageClient{
		Client: client,
	}, nil
}
