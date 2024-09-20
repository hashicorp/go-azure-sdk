package settingstoragequotaservice

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

type ListSettingStorageQuotaServicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceStorageQuotaBreakdown
}

type ListSettingStorageQuotaServicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceStorageQuotaBreakdown
}

type ListSettingStorageQuotaServicesOperationOptions struct {
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

func DefaultListSettingStorageQuotaServicesOperationOptions() ListSettingStorageQuotaServicesOperationOptions {
	return ListSettingStorageQuotaServicesOperationOptions{}
}

func (o ListSettingStorageQuotaServicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSettingStorageQuotaServicesOperationOptions) ToOData() *odata.Query {
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

func (o ListSettingStorageQuotaServicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSettingStorageQuotaServicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSettingStorageQuotaServicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSettingStorageQuotaServices - Get services from users
func (c SettingStorageQuotaServiceClient) ListSettingStorageQuotaServices(ctx context.Context, id stable.UserId, options ListSettingStorageQuotaServicesOperationOptions) (result ListSettingStorageQuotaServicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSettingStorageQuotaServicesCustomPager{},
		Path:          fmt.Sprintf("%s/settings/storage/quota/services", id.ID()),
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
		Values *[]stable.ServiceStorageQuotaBreakdown `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSettingStorageQuotaServicesComplete retrieves all the results into a single object
func (c SettingStorageQuotaServiceClient) ListSettingStorageQuotaServicesComplete(ctx context.Context, id stable.UserId, options ListSettingStorageQuotaServicesOperationOptions) (ListSettingStorageQuotaServicesCompleteResult, error) {
	return c.ListSettingStorageQuotaServicesCompleteMatchingPredicate(ctx, id, options, ServiceStorageQuotaBreakdownOperationPredicate{})
}

// ListSettingStorageQuotaServicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SettingStorageQuotaServiceClient) ListSettingStorageQuotaServicesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListSettingStorageQuotaServicesOperationOptions, predicate ServiceStorageQuotaBreakdownOperationPredicate) (result ListSettingStorageQuotaServicesCompleteResult, err error) {
	items := make([]stable.ServiceStorageQuotaBreakdown, 0)

	resp, err := c.ListSettingStorageQuotaServices(ctx, id, options)
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

	result = ListSettingStorageQuotaServicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
