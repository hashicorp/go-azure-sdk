package directoryaudit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAuditClient struct {
	Client *msgraph.Client
}

func NewDirectoryAuditClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryAuditClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryaudit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryAuditClient: %+v", err)
	}

	return &DirectoryAuditClient{
		Client: client,
	}, nil
}
