package macossoftwareupdateaccountsummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMacOSSoftwareUpdateAccountSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MacOSSoftwareUpdateAccountSummary
}

type CreateMacOSSoftwareUpdateAccountSummaryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMacOSSoftwareUpdateAccountSummaryOperationOptions() CreateMacOSSoftwareUpdateAccountSummaryOperationOptions {
	return CreateMacOSSoftwareUpdateAccountSummaryOperationOptions{}
}

func (o CreateMacOSSoftwareUpdateAccountSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMacOSSoftwareUpdateAccountSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMacOSSoftwareUpdateAccountSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMacOSSoftwareUpdateAccountSummary - Create new navigation property to macOSSoftwareUpdateAccountSummaries for
// deviceManagement
func (c MacOSSoftwareUpdateAccountSummaryClient) CreateMacOSSoftwareUpdateAccountSummary(ctx context.Context, input beta.MacOSSoftwareUpdateAccountSummary, options CreateMacOSSoftwareUpdateAccountSummaryOperationOptions) (result CreateMacOSSoftwareUpdateAccountSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/macOSSoftwareUpdateAccountSummaries",
		RetryFunc:     options.RetryFunc,
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

	var model beta.MacOSSoftwareUpdateAccountSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
