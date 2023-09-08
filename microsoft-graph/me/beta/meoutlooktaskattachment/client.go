package meoutlooktaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskAttachmentClient: %+v", err)
	}

	return &MeOutlookTaskAttachmentClient{
		Client: client,
	}, nil
}
