package onenotepageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenotePageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenotePageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotepageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenotePageParentNotebookClient: %+v", err)
	}

	return &OnenotePageParentNotebookClient{
		Client: client,
	}, nil
}
