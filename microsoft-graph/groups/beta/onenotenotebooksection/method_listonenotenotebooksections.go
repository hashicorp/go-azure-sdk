package onenotenotebooksection

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

type ListOnenoteNotebookSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenoteSection
}

type ListOnenoteNotebookSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenoteSection
}

type ListOnenoteNotebookSectionsOperationOptions struct {
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

func DefaultListOnenoteNotebookSectionsOperationOptions() ListOnenoteNotebookSectionsOperationOptions {
	return ListOnenoteNotebookSectionsOperationOptions{}
}

func (o ListOnenoteNotebookSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteNotebookSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteNotebookSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteNotebookSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteNotebookSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteNotebookSections - Get sections from groups. The sections in the notebook. Read-only. Nullable.
func (c OnenoteNotebookSectionClient) ListOnenoteNotebookSections(ctx context.Context, id beta.GroupIdOnenoteNotebookId, options ListOnenoteNotebookSectionsOperationOptions) (result ListOnenoteNotebookSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteNotebookSectionsCustomPager{},
		Path:          fmt.Sprintf("%s/sections", id.ID()),
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
		Values *[]beta.OnenoteSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteNotebookSectionsComplete retrieves all the results into a single object
func (c OnenoteNotebookSectionClient) ListOnenoteNotebookSectionsComplete(ctx context.Context, id beta.GroupIdOnenoteNotebookId, options ListOnenoteNotebookSectionsOperationOptions) (ListOnenoteNotebookSectionsCompleteResult, error) {
	return c.ListOnenoteNotebookSectionsCompleteMatchingPredicate(ctx, id, options, OnenoteSectionOperationPredicate{})
}

// ListOnenoteNotebookSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteNotebookSectionClient) ListOnenoteNotebookSectionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdOnenoteNotebookId, options ListOnenoteNotebookSectionsOperationOptions, predicate OnenoteSectionOperationPredicate) (result ListOnenoteNotebookSectionsCompleteResult, err error) {
	items := make([]beta.OnenoteSection, 0)

	resp, err := c.ListOnenoteNotebookSections(ctx, id, options)
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

	result = ListOnenoteNotebookSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
