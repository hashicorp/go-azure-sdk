package onenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &OnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
