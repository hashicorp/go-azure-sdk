package meonenotesectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
