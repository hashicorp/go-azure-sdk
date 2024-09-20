package windowsinformationprotectionapplearningsummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsInformationProtectionAppLearningSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsInformationProtectionAppLearningSummary
}

type CreateWindowsInformationProtectionAppLearningSummaryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsInformationProtectionAppLearningSummaryOperationOptions() CreateWindowsInformationProtectionAppLearningSummaryOperationOptions {
	return CreateWindowsInformationProtectionAppLearningSummaryOperationOptions{}
}

func (o CreateWindowsInformationProtectionAppLearningSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsInformationProtectionAppLearningSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsInformationProtectionAppLearningSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsInformationProtectionAppLearningSummary - Create new navigation property to
// windowsInformationProtectionAppLearningSummaries for deviceManagement
func (c WindowsInformationProtectionAppLearningSummaryClient) CreateWindowsInformationProtectionAppLearningSummary(ctx context.Context, input beta.WindowsInformationProtectionAppLearningSummary, options CreateWindowsInformationProtectionAppLearningSummaryOperationOptions) (result CreateWindowsInformationProtectionAppLearningSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsInformationProtectionAppLearningSummaries",
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

	var model beta.WindowsInformationProtectionAppLearningSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
