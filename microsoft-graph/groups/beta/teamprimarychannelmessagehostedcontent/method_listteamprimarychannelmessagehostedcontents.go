package teamprimarychannelmessagehostedcontent

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

type ListTeamPrimaryChannelMessageHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageHostedContentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTeamPrimaryChannelMessageHostedContentsOperationOptions() ListTeamPrimaryChannelMessageHostedContentsOperationOptions {
	return ListTeamPrimaryChannelMessageHostedContentsOperationOptions{}
}

func (o ListTeamPrimaryChannelMessageHostedContentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPrimaryChannelMessageHostedContentsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPrimaryChannelMessageHostedContentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPrimaryChannelMessageHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelMessageHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelMessageHostedContents - Get hostedContents from groups. Content in a message hosted by
// Microsoft Teams - for example, images or code snippets.
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContents(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageId, options ListTeamPrimaryChannelMessageHostedContentsOperationOptions) (result ListTeamPrimaryChannelMessageHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPrimaryChannelMessageHostedContentsCustomPager{},
		Path:          fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]beta.ChatMessageHostedContent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelMessageHostedContentsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContentsComplete(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageId, options ListTeamPrimaryChannelMessageHostedContentsOperationOptions) (ListTeamPrimaryChannelMessageHostedContentsCompleteResult, error) {
	return c.ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx, id, options, ChatMessageHostedContentOperationPredicate{})
}

// ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelMessageHostedContentClient) ListTeamPrimaryChannelMessageHostedContentsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageId, options ListTeamPrimaryChannelMessageHostedContentsOperationOptions, predicate ChatMessageHostedContentOperationPredicate) (result ListTeamPrimaryChannelMessageHostedContentsCompleteResult, err error) {
	items := make([]beta.ChatMessageHostedContent, 0)

	resp, err := c.ListTeamPrimaryChannelMessageHostedContents(ctx, id, options)
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

	result = ListTeamPrimaryChannelMessageHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
