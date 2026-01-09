package sitedocumentprocessingjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteDocumentProcessingJobClient struct {
	Client *msgraph.Client
}

func NewSiteDocumentProcessingJobClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteDocumentProcessingJobClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitedocumentprocessingjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteDocumentProcessingJobClient: %+v", err)
	}

	return &SiteDocumentProcessingJobClient{
		Client: client,
	}, nil
}
