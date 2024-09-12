package appconsentrequestuserconsentrequestapprovalstage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestUserConsentRequestApprovalStageClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestUserConsentRequestApprovalStageClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestUserConsentRequestApprovalStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestuserconsentrequestapprovalstage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestUserConsentRequestApprovalStageClient: %+v", err)
	}

	return &AppConsentRequestUserConsentRequestApprovalStageClient{
		Client: client,
	}, nil
}
