package grouponenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
