package ledgerdigestuploads

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LedgerDigestUploadsClient struct {
	Client *resourcemanager.Client
}

func NewLedgerDigestUploadsClientWithBaseURI(sdkApi sdkEnv.Api) (*LedgerDigestUploadsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "ledgerdigestuploads", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LedgerDigestUploadsClient: %+v", err)
	}

	return &LedgerDigestUploadsClient{
		Client: client,
	}, nil
}
