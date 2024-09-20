package manageddevicedetectedapp

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

type ListManagedDeviceDetectedAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DetectedApp
}

type ListManagedDeviceDetectedAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DetectedApp
}

type ListManagedDeviceDetectedAppsOperationOptions struct {
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

func DefaultListManagedDeviceDetectedAppsOperationOptions() ListManagedDeviceDetectedAppsOperationOptions {
	return ListManagedDeviceDetectedAppsOperationOptions{}
}

func (o ListManagedDeviceDetectedAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceDetectedAppsOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceDetectedAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceDetectedAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceDetectedAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceDetectedApps - Get detectedApps from deviceManagement. All applications currently installed on the
// device
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedApps(ctx context.Context, id beta.DeviceManagementManagedDeviceId, options ListManagedDeviceDetectedAppsOperationOptions) (result ListManagedDeviceDetectedAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceDetectedAppsCustomPager{},
		Path:          fmt.Sprintf("%s/detectedApps", id.ID()),
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

// ListManagedDeviceDetectedAppsComplete retrieves all the results into a single object
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedAppsComplete(ctx context.Context, id beta.DeviceManagementManagedDeviceId, options ListManagedDeviceDetectedAppsOperationOptions) (ListManagedDeviceDetectedAppsCompleteResult, error) {
	return c.ListManagedDeviceDetectedAppsCompleteMatchingPredicate(ctx, id, options, DetectedAppOperationPredicate{})
}

// ListManagedDeviceDetectedAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDetectedAppClient) ListManagedDeviceDetectedAppsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementManagedDeviceId, options ListManagedDeviceDetectedAppsOperationOptions, predicate DetectedAppOperationPredicate) (result ListManagedDeviceDetectedAppsCompleteResult, err error) {
	items := make([]beta.DetectedApp, 0)

	resp, err := c.ListManagedDeviceDetectedApps(ctx, id, options)
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

	result = ListManagedDeviceDetectedAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
