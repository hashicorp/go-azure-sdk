package meonenotepageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenotePageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenotePageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenotePageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotepageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenotePageParentNotebookClient: %+v", err)
	}

	return &MeOnenotePageParentNotebookClient{
		Client: client,
	}, nil
}
