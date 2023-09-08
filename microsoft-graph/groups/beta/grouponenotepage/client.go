package grouponenotepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenotePageClient struct {
	Client *msgraph.Client
}

func NewGroupOnenotePageClientWithBaseURI(api sdkEnv.Api) (*GroupOnenotePageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenotePageClient: %+v", err)
	}

	return &GroupOnenotePageClient{
		Client: client,
	}, nil
}
