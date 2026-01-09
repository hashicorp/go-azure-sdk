package elevationrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetElevationRequestAllElevationRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrivilegeManagementElevationRequest
}

type GetElevationRequestAllElevationRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrivilegeManagementElevationRequest
}

type GetElevationRequestAllElevationRequestsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetElevationRequestAllElevationRequestsOperationOptions() GetElevationRequestAllElevationRequestsOperationOptions {
	return GetElevationRequestAllElevationRequestsOperationOptions{}
}

func (o GetElevationRequestAllElevationRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetElevationRequestAllElevationRequestsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetElevationRequestAllElevationRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetElevationRequestAllElevationRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetElevationRequestAllElevationRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetElevationRequestAllElevationRequests - Invoke action getAllElevationRequests
func (c ElevationRequestClient) GetElevationRequestAllElevationRequests(ctx context.Context, id beta.DeviceManagementElevationRequestId, options GetElevationRequestAllElevationRequestsOperationOptions) (result GetElevationRequestAllElevationRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetElevationRequestAllElevationRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/getAllElevationRequests", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.PrivilegeManagementElevationRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetElevationRequestAllElevationRequestsComplete retrieves all the results into a single object
func (c ElevationRequestClient) GetElevationRequestAllElevationRequestsComplete(ctx context.Context, id beta.DeviceManagementElevationRequestId, options GetElevationRequestAllElevationRequestsOperationOptions) (GetElevationRequestAllElevationRequestsCompleteResult, error) {
	return c.GetElevationRequestAllElevationRequestsCompleteMatchingPredicate(ctx, id, options, PrivilegeManagementElevationRequestOperationPredicate{})
}

// GetElevationRequestAllElevationRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ElevationRequestClient) GetElevationRequestAllElevationRequestsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementElevationRequestId, options GetElevationRequestAllElevationRequestsOperationOptions, predicate PrivilegeManagementElevationRequestOperationPredicate) (result GetElevationRequestAllElevationRequestsCompleteResult, err error) {
	items := make([]beta.PrivilegeManagementElevationRequest, 0)

	resp, err := c.GetElevationRequestAllElevationRequests(ctx, id, options)
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

	result = GetElevationRequestAllElevationRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
