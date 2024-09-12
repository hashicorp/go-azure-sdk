package importeddeviceidentity

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

type ListImportedDeviceIdentityImportListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedDeviceIdentityResult
}

type ListImportedDeviceIdentityImportListsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedDeviceIdentityResult
}

type ListImportedDeviceIdentityImportListsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListImportedDeviceIdentityImportListsOperationOptions() ListImportedDeviceIdentityImportListsOperationOptions {
	return ListImportedDeviceIdentityImportListsOperationOptions{}
}

func (o ListImportedDeviceIdentityImportListsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListImportedDeviceIdentityImportListsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListImportedDeviceIdentityImportListsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListImportedDeviceIdentityImportListsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedDeviceIdentityImportListsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedDeviceIdentityImportLists - Invoke action importDeviceIdentityList
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentityImportLists(ctx context.Context, input ListImportedDeviceIdentityImportListsRequest, options ListImportedDeviceIdentityImportListsOperationOptions) (result ListImportedDeviceIdentityImportListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListImportedDeviceIdentityImportListsCustomPager{},
		Path:          "/deviceManagement/importedDeviceIdentities/importDeviceIdentityList",
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
		Values *[]beta.ImportedDeviceIdentityResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImportedDeviceIdentityImportListsComplete retrieves all the results into a single object
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentityImportListsComplete(ctx context.Context, input ListImportedDeviceIdentityImportListsRequest, options ListImportedDeviceIdentityImportListsOperationOptions) (ListImportedDeviceIdentityImportListsCompleteResult, error) {
	return c.ListImportedDeviceIdentityImportListsCompleteMatchingPredicate(ctx, input, options, ImportedDeviceIdentityResultOperationPredicate{})
}

// ListImportedDeviceIdentityImportListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentityImportListsCompleteMatchingPredicate(ctx context.Context, input ListImportedDeviceIdentityImportListsRequest, options ListImportedDeviceIdentityImportListsOperationOptions, predicate ImportedDeviceIdentityResultOperationPredicate) (result ListImportedDeviceIdentityImportListsCompleteResult, err error) {
	items := make([]beta.ImportedDeviceIdentityResult, 0)

	resp, err := c.ListImportedDeviceIdentityImportLists(ctx, input, options)
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

	result = ListImportedDeviceIdentityImportListsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
