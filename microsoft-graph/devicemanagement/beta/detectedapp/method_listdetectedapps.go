package detectedapp

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

type ListDetectedAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DetectedApp
}

type ListDetectedAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DetectedApp
}

type ListDetectedAppsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDetectedAppsOperationOptions() ListDetectedAppsOperationOptions {
	return ListDetectedAppsOperationOptions{}
}

func (o ListDetectedAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDetectedAppsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDetectedAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDetectedAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDetectedAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDetectedApps - Get detectedApps from deviceManagement. The list of detected apps associated with a device.
func (c DetectedAppClient) ListDetectedApps(ctx context.Context, options ListDetectedAppsOperationOptions) (result ListDetectedAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDetectedAppsCustomPager{},
		Path:          "/deviceManagement/detectedApps",
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
		Values *[]beta.DetectedApp `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDetectedAppsComplete retrieves all the results into a single object
func (c DetectedAppClient) ListDetectedAppsComplete(ctx context.Context, options ListDetectedAppsOperationOptions) (ListDetectedAppsCompleteResult, error) {
	return c.ListDetectedAppsCompleteMatchingPredicate(ctx, options, DetectedAppOperationPredicate{})
}

// ListDetectedAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectedAppClient) ListDetectedAppsCompleteMatchingPredicate(ctx context.Context, options ListDetectedAppsOperationOptions, predicate DetectedAppOperationPredicate) (result ListDetectedAppsCompleteResult, err error) {
	items := make([]beta.DetectedApp, 0)

	resp, err := c.ListDetectedApps(ctx, options)
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

	result = ListDetectedAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
