package accessreviewdefinitioninstancedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
