package manageddevicewindowsosimage

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

type ListManagedDeviceWindowsOSImagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceWindowsOperatingSystemImage
}

type ListManagedDeviceWindowsOSImagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceWindowsOperatingSystemImage
}

type ListManagedDeviceWindowsOSImagesOperationOptions struct {
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

func DefaultListManagedDeviceWindowsOSImagesOperationOptions() ListManagedDeviceWindowsOSImagesOperationOptions {
	return ListManagedDeviceWindowsOSImagesOperationOptions{}
}

func (o ListManagedDeviceWindowsOSImagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceWindowsOSImagesOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceWindowsOSImagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceWindowsOSImagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceWindowsOSImagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceWindowsOSImages - Get managedDeviceWindowsOSImages from deviceManagement. A list of
// ManagedDeviceWindowsOperatingSystemImages
func (c ManagedDeviceWindowsOSImageClient) ListManagedDeviceWindowsOSImages(ctx context.Context, options ListManagedDeviceWindowsOSImagesOperationOptions) (result ListManagedDeviceWindowsOSImagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceWindowsOSImagesCustomPager{},
		Path:          "/deviceManagement/managedDeviceWindowsOSImages",
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
		Values *[]beta.ManagedDeviceWindowsOperatingSystemImage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceWindowsOSImagesComplete retrieves all the results into a single object
func (c ManagedDeviceWindowsOSImageClient) ListManagedDeviceWindowsOSImagesComplete(ctx context.Context, options ListManagedDeviceWindowsOSImagesOperationOptions) (ListManagedDeviceWindowsOSImagesCompleteResult, error) {
	return c.ListManagedDeviceWindowsOSImagesCompleteMatchingPredicate(ctx, options, ManagedDeviceWindowsOperatingSystemImageOperationPredicate{})
}

// ListManagedDeviceWindowsOSImagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceWindowsOSImageClient) ListManagedDeviceWindowsOSImagesCompleteMatchingPredicate(ctx context.Context, options ListManagedDeviceWindowsOSImagesOperationOptions, predicate ManagedDeviceWindowsOperatingSystemImageOperationPredicate) (result ListManagedDeviceWindowsOSImagesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceWindowsOperatingSystemImage, 0)

	resp, err := c.ListManagedDeviceWindowsOSImages(ctx, options)
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

	result = ListManagedDeviceWindowsOSImagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
