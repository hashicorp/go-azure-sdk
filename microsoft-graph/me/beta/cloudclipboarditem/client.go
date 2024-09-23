package cloudclipboarditem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudClipboardItemClient struct {
	Client *msgraph.Client
}

func NewCloudClipboardItemClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudClipboardItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudclipboarditem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudClipboardItemClient: %+v", err)
	}

	return &CloudClipboardItemClient{
		Client: client,
	}, nil
}
