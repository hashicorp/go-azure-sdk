package windowsinformationprotectionnetworklearningsummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsInformationProtectionNetworkLearningSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.WindowsInformationProtectionNetworkLearningSummary
}

type CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions() CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions {
	return CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions{}
}

func (o CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsInformationProtectionNetworkLearningSummary - Create windowsInformationProtectionNetworkLearningSummary.
// Create a new windowsInformationProtectionNetworkLearningSummary object.
func (c WindowsInformationProtectionNetworkLearningSummaryClient) CreateWindowsInformationProtectionNetworkLearningSummary(ctx context.Context, input stable.WindowsInformationProtectionNetworkLearningSummary, options CreateWindowsInformationProtectionNetworkLearningSummaryOperationOptions) (result CreateWindowsInformationProtectionNetworkLearningSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.WindowsInformationProtectionNetworkLearningSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
