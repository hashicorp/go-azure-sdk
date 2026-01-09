package driveitemversion

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

type ListDriveItemVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DriveItemVersion
}

type ListDriveItemVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DriveItemVersion
}

type ListDriveItemVersionsOperationOptions struct {
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

func DefaultListDriveItemVersionsOperationOptions() ListDriveItemVersionsOperationOptions {
	return ListDriveItemVersionsOperationOptions{}
}

func (o ListDriveItemVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemVersions - Get versions from me. The list of previous versions of the item. For more info, see getting
// previous versions. Read-only. Nullable.
func (c DriveItemVersionClient) ListDriveItemVersions(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemVersionsOperationOptions) (result ListDriveItemVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveItemVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]stable.DriveItemVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveItemVersionsComplete retrieves all the results into a single object
func (c DriveItemVersionClient) ListDriveItemVersionsComplete(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemVersionsOperationOptions) (ListDriveItemVersionsCompleteResult, error) {
	return c.ListDriveItemVersionsCompleteMatchingPredicate(ctx, id, options, DriveItemVersionOperationPredicate{})
}

// ListDriveItemVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemVersionClient) ListDriveItemVersionsCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemVersionsOperationOptions, predicate DriveItemVersionOperationPredicate) (result ListDriveItemVersionsCompleteResult, err error) {
	items := make([]stable.DriveItemVersion, 0)

	resp, err := c.ListDriveItemVersions(ctx, id, options)
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

	result = ListDriveItemVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
