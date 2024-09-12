package driverootlistitemdocumentsetversion

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

type ListDriveRootListItemDocumentSetVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DocumentSetVersion
}

type ListDriveRootListItemDocumentSetVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DocumentSetVersion
}

type ListDriveRootListItemDocumentSetVersionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDriveRootListItemDocumentSetVersionsOperationOptions() ListDriveRootListItemDocumentSetVersionsOperationOptions {
	return ListDriveRootListItemDocumentSetVersionsOperationOptions{}
}

func (o ListDriveRootListItemDocumentSetVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootListItemDocumentSetVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootListItemDocumentSetVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootListItemDocumentSetVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootListItemDocumentSetVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootListItemDocumentSetVersions - Get documentSetVersions from users. Version information for a document set
// version created by a user.
func (c DriveRootListItemDocumentSetVersionClient) ListDriveRootListItemDocumentSetVersions(ctx context.Context, id stable.UserIdDriveId, options ListDriveRootListItemDocumentSetVersionsOperationOptions) (result ListDriveRootListItemDocumentSetVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootListItemDocumentSetVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/root/listItem/documentSetVersions", id.ID()),
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
		Values *[]stable.DocumentSetVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveRootListItemDocumentSetVersionsComplete retrieves all the results into a single object
func (c DriveRootListItemDocumentSetVersionClient) ListDriveRootListItemDocumentSetVersionsComplete(ctx context.Context, id stable.UserIdDriveId, options ListDriveRootListItemDocumentSetVersionsOperationOptions) (ListDriveRootListItemDocumentSetVersionsCompleteResult, error) {
	return c.ListDriveRootListItemDocumentSetVersionsCompleteMatchingPredicate(ctx, id, options, DocumentSetVersionOperationPredicate{})
}

// ListDriveRootListItemDocumentSetVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootListItemDocumentSetVersionClient) ListDriveRootListItemDocumentSetVersionsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdDriveId, options ListDriveRootListItemDocumentSetVersionsOperationOptions, predicate DocumentSetVersionOperationPredicate) (result ListDriveRootListItemDocumentSetVersionsCompleteResult, err error) {
	items := make([]stable.DocumentSetVersion, 0)

	resp, err := c.ListDriveRootListItemDocumentSetVersions(ctx, id, options)
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

	result = ListDriveRootListItemDocumentSetVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
