package macossoftwareupdateaccountsummarycategorysummary

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions() DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions {
	return DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions{}
}

func (o DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteMacOSSoftwareUpdateAccountSummaryCategorySummary - Delete navigation property categorySummaries for
// deviceManagement
func (c MacOSSoftwareUpdateAccountSummaryCategorySummaryClient) DeleteMacOSSoftwareUpdateAccountSummaryCategorySummary(ctx context.Context, id beta.DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId, options DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationOptions) (result DeleteMacOSSoftwareUpdateAccountSummaryCategorySummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	return
}
