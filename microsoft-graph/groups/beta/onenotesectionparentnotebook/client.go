package onenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionParentNotebookClient: %+v", err)
	}

	return &OnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
