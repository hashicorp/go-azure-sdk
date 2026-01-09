package joinedteamprimarychannelallmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListJoinedTeamPrimaryChannelAllMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelAllMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelAllMembersOperationOptions struct {
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

func DefaultListJoinedTeamPrimaryChannelAllMembersOperationOptions() ListJoinedTeamPrimaryChannelAllMembersOperationOptions {
	return ListJoinedTeamPrimaryChannelAllMembersOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelAllMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelAllMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelAllMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelAllMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelAllMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelAllMembers - Get allMembers from users. A collection of membership records associated
// with the channel, including both direct and indirect members of shared channels.
func (c JoinedTeamPrimaryChannelAllMemberClient) ListJoinedTeamPrimaryChannelAllMembers(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelAllMembersOperationOptions) (result ListJoinedTeamPrimaryChannelAllMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelAllMembersCustomPager{},
		Path:          fmt.Sprintf("%s/primaryChannel/allMembers", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.ConversationMember, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalConversationMemberImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ConversationMember (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListJoinedTeamPrimaryChannelAllMembersComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelAllMemberClient) ListJoinedTeamPrimaryChannelAllMembersComplete(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelAllMembersOperationOptions) (ListJoinedTeamPrimaryChannelAllMembersCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelAllMembersCompleteMatchingPredicate(ctx, id, options, ConversationMemberOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelAllMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelAllMemberClient) ListJoinedTeamPrimaryChannelAllMembersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelAllMembersOperationOptions, predicate ConversationMemberOperationPredicate) (result ListJoinedTeamPrimaryChannelAllMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelAllMembers(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelAllMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
