package provisioning

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProvisioningsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ProvisioningObjectSummary
}

type ListProvisioningsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ProvisioningObjectSummary
}

type ListProvisioningsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListProvisioningsOperationOptions() ListProvisioningsOperationOptions {
	return ListProvisioningsOperationOptions{}
}

func (o ListProvisioningsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProvisioningsOperationOptions) ToOData() *odata.Query {
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

func (o ListProvisioningsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProvisioningsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProvisioningsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProvisionings - List provisioningObjectSummary. Get all provisioning events that occurred in your tenant, such as
// the deletion of a group in a target application or the creation of a user when provisioning user accounts from your
// HR system.
func (c ProvisioningClient) ListProvisionings(ctx context.Context, options ListProvisioningsOperationOptions) (result ListProvisioningsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProvisioningsCustomPager{},
		Path:          "/auditLogs/provisioning",
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
		Values *[]stable.ProvisioningObjectSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProvisioningsComplete retrieves all the results into a single object
func (c ProvisioningClient) ListProvisioningsComplete(ctx context.Context, options ListProvisioningsOperationOptions) (ListProvisioningsCompleteResult, error) {
	return c.ListProvisioningsCompleteMatchingPredicate(ctx, options, ProvisioningObjectSummaryOperationPredicate{})
}

// ListProvisioningsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProvisioningClient) ListProvisioningsCompleteMatchingPredicate(ctx context.Context, options ListProvisioningsOperationOptions, predicate ProvisioningObjectSummaryOperationPredicate) (result ListProvisioningsCompleteResult, err error) {
	items := make([]stable.ProvisioningObjectSummary, 0)

	resp, err := c.ListProvisionings(ctx, options)
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

	result = ListProvisioningsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
