package joinedteamprimarychannelmember

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

type RemoveJoinedTeamPrimaryChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ActionResultPart
}

type RemoveJoinedTeamPrimaryChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ActionResultPart
}

type RemoveJoinedTeamPrimaryChannelMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveJoinedTeamPrimaryChannelMembersOperationOptions() RemoveJoinedTeamPrimaryChannelMembersOperationOptions {
	return RemoveJoinedTeamPrimaryChannelMembersOperationOptions{}
}

func (o RemoveJoinedTeamPrimaryChannelMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveJoinedTeamPrimaryChannelMembersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o RemoveJoinedTeamPrimaryChannelMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveJoinedTeamPrimaryChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveJoinedTeamPrimaryChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveJoinedTeamPrimaryChannelMembers - Invoke action remove. Remove multiple members from a team in a single
// request. The response provides details about which memberships could and couldn't be removed.
func (c JoinedTeamPrimaryChannelMemberClient) RemoveJoinedTeamPrimaryChannelMembers(ctx context.Context, id stable.MeJoinedTeamId, input RemoveJoinedTeamPrimaryChannelMembersRequest, options RemoveJoinedTeamPrimaryChannelMembersOperationOptions) (result RemoveJoinedTeamPrimaryChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveJoinedTeamPrimaryChannelMembersCustomPager{},
		Path:          fmt.Sprintf("%s/primaryChannel/members/remove", id.ID()),
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

	temp := make([]stable.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// RemoveJoinedTeamPrimaryChannelMembersComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMemberClient) RemoveJoinedTeamPrimaryChannelMembersComplete(ctx context.Context, id stable.MeJoinedTeamId, input RemoveJoinedTeamPrimaryChannelMembersRequest, options RemoveJoinedTeamPrimaryChannelMembersOperationOptions) (RemoveJoinedTeamPrimaryChannelMembersCompleteResult, error) {
	return c.RemoveJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// RemoveJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMemberClient) RemoveJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, input RemoveJoinedTeamPrimaryChannelMembersRequest, options RemoveJoinedTeamPrimaryChannelMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result RemoveJoinedTeamPrimaryChannelMembersCompleteResult, err error) {
	items := make([]stable.ActionResultPart, 0)

	resp, err := c.RemoveJoinedTeamPrimaryChannelMembers(ctx, id, input, options)
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

	result = RemoveJoinedTeamPrimaryChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
