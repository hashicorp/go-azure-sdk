package cloudpc

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

type ValidateCloudPCsBulkResizesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCResizeValidationResult
}

type ValidateCloudPCsBulkResizesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCResizeValidationResult
}

type ValidateCloudPCsBulkResizesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultValidateCloudPCsBulkResizesOperationOptions() ValidateCloudPCsBulkResizesOperationOptions {
	return ValidateCloudPCsBulkResizesOperationOptions{}
}

func (o ValidateCloudPCsBulkResizesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateCloudPCsBulkResizesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ValidateCloudPCsBulkResizesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ValidateCloudPCsBulkResizesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ValidateCloudPCsBulkResizesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ValidateCloudPCsBulkResizes - Invoke action validateBulkResize. Validate that a set of cloudPC devices meet the
// requirements to be bulk resized.
func (c CloudPCClient) ValidateCloudPCsBulkResizes(ctx context.Context, input ValidateCloudPCsBulkResizesRequest, options ValidateCloudPCsBulkResizesOperationOptions) (result ValidateCloudPCsBulkResizesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ValidateCloudPCsBulkResizesCustomPager{},
		Path:          "/me/cloudPCs/validateBulkResize",
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
		Values *[]beta.CloudPCResizeValidationResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ValidateCloudPCsBulkResizesComplete retrieves all the results into a single object
func (c CloudPCClient) ValidateCloudPCsBulkResizesComplete(ctx context.Context, input ValidateCloudPCsBulkResizesRequest, options ValidateCloudPCsBulkResizesOperationOptions) (ValidateCloudPCsBulkResizesCompleteResult, error) {
	return c.ValidateCloudPCsBulkResizesCompleteMatchingPredicate(ctx, input, options, CloudPCResizeValidationResultOperationPredicate{})
}

// ValidateCloudPCsBulkResizesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCClient) ValidateCloudPCsBulkResizesCompleteMatchingPredicate(ctx context.Context, input ValidateCloudPCsBulkResizesRequest, options ValidateCloudPCsBulkResizesOperationOptions, predicate CloudPCResizeValidationResultOperationPredicate) (result ValidateCloudPCsBulkResizesCompleteResult, err error) {
	items := make([]beta.CloudPCResizeValidationResult, 0)

	resp, err := c.ValidateCloudPCsBulkResizes(ctx, input, options)
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

	result = ValidateCloudPCsBulkResizesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
