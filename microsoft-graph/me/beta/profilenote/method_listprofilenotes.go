package profilenote

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

type ListProfileNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonAnnotation
}

type ListProfileNotesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonAnnotation
}

type ListProfileNotesOperationOptions struct {
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

func DefaultListProfileNotesOperationOptions() ListProfileNotesOperationOptions {
	return ListProfileNotesOperationOptions{}
}

func (o ListProfileNotesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileNotesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileNotesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileNotesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileNotesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileNotes - List notes. Retrieve a list of personAnnotation objects from a user's profile.
func (c ProfileNoteClient) ListProfileNotes(ctx context.Context, options ListProfileNotesOperationOptions) (result ListProfileNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileNotesCustomPager{},
		Path:          "/me/profile/notes",
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
		Values *[]beta.PersonAnnotation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileNotesComplete retrieves all the results into a single object
func (c ProfileNoteClient) ListProfileNotesComplete(ctx context.Context, options ListProfileNotesOperationOptions) (ListProfileNotesCompleteResult, error) {
	return c.ListProfileNotesCompleteMatchingPredicate(ctx, options, PersonAnnotationOperationPredicate{})
}

// ListProfileNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileNoteClient) ListProfileNotesCompleteMatchingPredicate(ctx context.Context, options ListProfileNotesOperationOptions, predicate PersonAnnotationOperationPredicate) (result ListProfileNotesCompleteResult, err error) {
	items := make([]beta.PersonAnnotation, 0)

	resp, err := c.ListProfileNotes(ctx, options)
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

	result = ListProfileNotesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
