package meonenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionParentNotebookClient: %+v", err)
	}

	return &MeOnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
