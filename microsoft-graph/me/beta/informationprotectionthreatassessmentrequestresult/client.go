package informationprotectionthreatassessmentrequestresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionThreatAssessmentRequestResultClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionThreatAssessmentRequestResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionthreatassessmentrequestresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionThreatAssessmentRequestResultClient: %+v", err)
	}

	return &InformationProtectionThreatAssessmentRequestResultClient{
		Client: client,
	}, nil
}
