package teamprimarychannelmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddTeamPrimaryChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type AddTeamPrimaryChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type AddTeamPrimaryChannelMembersOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAddTeamPrimaryChannelMembersOperationOptions() AddTeamPrimaryChannelMembersOperationOptions {
	return AddTeamPrimaryChannelMembersOperationOptions{}
}

func (o AddTeamPrimaryChannelMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddTeamPrimaryChannelMembersOperationOptions) ToOData() *odata.Query {
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

func (o AddTeamPrimaryChannelMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddTeamPrimaryChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddTeamPrimaryChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddTeamPrimaryChannelMembers - Invoke action add. Add multiple members in a single request to a team. The response
// provides details about which memberships could and couldn't be created.
func (c TeamPrimaryChannelMemberClient) AddTeamPrimaryChannelMembers(ctx context.Context, id beta.GroupId, input AddTeamPrimaryChannelMembersRequest, options AddTeamPrimaryChannelMembersOperationOptions) (result AddTeamPrimaryChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddTeamPrimaryChannelMembersCustomPager{},
		Path:          fmt.Sprintf("%s/team/primaryChannel/members/add", id.ID()),
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

	temp := make([]beta.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// AddTeamPrimaryChannelMembersComplete retrieves all the results into a single object
func (c TeamPrimaryChannelMemberClient) AddTeamPrimaryChannelMembersComplete(ctx context.Context, id beta.GroupId, input AddTeamPrimaryChannelMembersRequest, options AddTeamPrimaryChannelMembersOperationOptions) (AddTeamPrimaryChannelMembersCompleteResult, error) {
	return c.AddTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// AddTeamPrimaryChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelMemberClient) AddTeamPrimaryChannelMembersCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, input AddTeamPrimaryChannelMembersRequest, options AddTeamPrimaryChannelMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result AddTeamPrimaryChannelMembersCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.AddTeamPrimaryChannelMembers(ctx, id, input, options)
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

	result = AddTeamPrimaryChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
