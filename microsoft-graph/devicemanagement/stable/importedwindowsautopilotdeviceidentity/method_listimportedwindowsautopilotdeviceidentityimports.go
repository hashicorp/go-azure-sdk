package importedwindowsautopilotdeviceidentity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListImportedWindowsAutopilotDeviceIdentityImportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions() ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions {
	return ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions{}
}

func (o ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedWindowsAutopilotDeviceIdentityImports - Invoke action import. Not yet documented
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImports(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest, options ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions) (result ListImportedWindowsAutopilotDeviceIdentityImportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager{},
		Path:          "/deviceManagement/importedWindowsAutopilotDeviceIdentities/import",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.ImportedWindowsAutopilotDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImportedWindowsAutopilotDeviceIdentityImportsComplete retrieves all the results into a single object
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImportsComplete(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest, options ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions) (ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult, error) {
	return c.ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate(ctx, input, options, ImportedWindowsAutopilotDeviceIdentityOperationPredicate{})
}

// ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest, options ListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions, predicate ImportedWindowsAutopilotDeviceIdentityOperationPredicate) (result ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult, err error) {
	items := make([]stable.ImportedWindowsAutopilotDeviceIdentity, 0)

	resp, err := c.ListImportedWindowsAutopilotDeviceIdentityImports(ctx, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
