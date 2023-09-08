package useronenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &UserOnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
