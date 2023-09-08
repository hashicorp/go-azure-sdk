package meonenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &MeOnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
