package outlooktaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskAttachmentClient: %+v", err)
	}

	return &OutlookTaskAttachmentClient{
		Client: client,
	}, nil
}
