package onenotepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePageClient struct {
	Client *msgraph.Client
}

func NewOnenotePageClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenotePageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenotePageClient: %+v", err)
	}

	return &OnenotePageClient{
		Client: client,
	}, nil
}
