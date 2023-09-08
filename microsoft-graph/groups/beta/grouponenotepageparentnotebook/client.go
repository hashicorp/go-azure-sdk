package grouponenotepageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenotePageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenotePageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenotePageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotepageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenotePageParentNotebookClient: %+v", err)
	}

	return &GroupOnenotePageParentNotebookClient{
		Client: client,
	}, nil
}
