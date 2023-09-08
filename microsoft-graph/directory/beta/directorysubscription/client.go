package directorysubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectorySubscriptionClient struct {
	Client *msgraph.Client
}

func NewDirectorySubscriptionClientWithBaseURI(api sdkEnv.Api) (*DirectorySubscriptionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorysubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectorySubscriptionClient: %+v", err)
	}

	return &DirectorySubscriptionClient{
		Client: client,
	}, nil
}
