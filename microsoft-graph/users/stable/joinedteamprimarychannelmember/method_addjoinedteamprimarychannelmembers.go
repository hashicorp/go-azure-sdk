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

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddJoinedTeamPrimaryChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ActionResultPart
}

type AddJoinedTeamPrimaryChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ActionResultPart
}

type AddJoinedTeamPrimaryChannelMembersOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAddJoinedTeamPrimaryChannelMembersOperationOptions() AddJoinedTeamPrimaryChannelMembersOperationOptions {
	return AddJoinedTeamPrimaryChannelMembersOperationOptions{}
}

func (o AddJoinedTeamPrimaryChannelMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddJoinedTeamPrimaryChannelMembersOperationOptions) ToOData() *odata.Query {
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

func (o AddJoinedTeamPrimaryChannelMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddJoinedTeamPrimaryChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddJoinedTeamPrimaryChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddJoinedTeamPrimaryChannelMembers - Invoke action add. Add multiple members in a single request to a team. The
// response provides details about which memberships could and couldn't be created.
func (c JoinedTeamPrimaryChannelMemberClient) AddJoinedTeamPrimaryChannelMembers(ctx context.Context, id stable.UserIdJoinedTeamId, input AddJoinedTeamPrimaryChannelMembersRequest, options AddJoinedTeamPrimaryChannelMembersOperationOptions) (result AddJoinedTeamPrimaryChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddJoinedTeamPrimaryChannelMembersCustomPager{},
		Path:          fmt.Sprintf("%s/primaryChannel/members/add", id.ID()),
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

// AddJoinedTeamPrimaryChannelMembersComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMemberClient) AddJoinedTeamPrimaryChannelMembersComplete(ctx context.Context, id stable.UserIdJoinedTeamId, input AddJoinedTeamPrimaryChannelMembersRequest, options AddJoinedTeamPrimaryChannelMembersOperationOptions) (AddJoinedTeamPrimaryChannelMembersCompleteResult, error) {
	return c.AddJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// AddJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMemberClient) AddJoinedTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamId, input AddJoinedTeamPrimaryChannelMembersRequest, options AddJoinedTeamPrimaryChannelMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result AddJoinedTeamPrimaryChannelMembersCompleteResult, err error) {
	items := make([]stable.ActionResultPart, 0)

	resp, err := c.AddJoinedTeamPrimaryChannelMembers(ctx, id, input, options)
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

	result = AddJoinedTeamPrimaryChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
