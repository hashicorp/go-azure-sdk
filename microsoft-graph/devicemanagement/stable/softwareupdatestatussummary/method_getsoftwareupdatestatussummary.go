package softwareupdatestatussummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSoftwareUpdateStatusSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SoftwareUpdateStatusSummary
}

type GetSoftwareUpdateStatusSummaryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetSoftwareUpdateStatusSummaryOperationOptions() GetSoftwareUpdateStatusSummaryOperationOptions {
	return GetSoftwareUpdateStatusSummaryOperationOptions{}
}

func (o GetSoftwareUpdateStatusSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSoftwareUpdateStatusSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSoftwareUpdateStatusSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSoftwareUpdateStatusSummary - Get softwareUpdateStatusSummary. Read properties and relationships of the
// softwareUpdateStatusSummary object.
func (c SoftwareUpdateStatusSummaryClient) GetSoftwareUpdateStatusSummary(ctx context.Context, options GetSoftwareUpdateStatusSummaryOperationOptions) (result GetSoftwareUpdateStatusSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/softwareUpdateStatusSummary",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model stable.SoftwareUpdateStatusSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
