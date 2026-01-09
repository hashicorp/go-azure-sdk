package deleteditem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedItemClient struct {
	Client *msgraph.Client
}

func NewDeletedItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deleteditem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedItemClient: %+v", err)
	}

	return &DeletedItemClient{
		Client: client,
	}, nil
}
