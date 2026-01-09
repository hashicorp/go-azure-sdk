package accessreviewdefinitioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceContactedReviewerClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
