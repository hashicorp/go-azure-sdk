package windowsautopilotdeploymentprofileassignment

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

type ListWindowsAutopilotDeploymentProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsAutopilotDeploymentProfileAssignment
}

type ListWindowsAutopilotDeploymentProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsAutopilotDeploymentProfileAssignment
}

type ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions struct {
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

func DefaultListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions() ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions {
	return ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions{}
}

func (o ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsAutopilotDeploymentProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsAutopilotDeploymentProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsAutopilotDeploymentProfileAssignments - Get assignments from deviceManagement. The list of group
// assignments for the profile.
func (c WindowsAutopilotDeploymentProfileAssignmentClient) ListWindowsAutopilotDeploymentProfileAssignments(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileId, options ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions) (result ListWindowsAutopilotDeploymentProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListWindowsAutopilotDeploymentProfileAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.WindowsAutopilotDeploymentProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsAutopilotDeploymentProfileAssignmentsComplete retrieves all the results into a single object
func (c WindowsAutopilotDeploymentProfileAssignmentClient) ListWindowsAutopilotDeploymentProfileAssignmentsComplete(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileId, options ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions) (ListWindowsAutopilotDeploymentProfileAssignmentsCompleteResult, error) {
	return c.ListWindowsAutopilotDeploymentProfileAssignmentsCompleteMatchingPredicate(ctx, id, options, WindowsAutopilotDeploymentProfileAssignmentOperationPredicate{})
}

// ListWindowsAutopilotDeploymentProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsAutopilotDeploymentProfileAssignmentClient) ListWindowsAutopilotDeploymentProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileId, options ListWindowsAutopilotDeploymentProfileAssignmentsOperationOptions, predicate WindowsAutopilotDeploymentProfileAssignmentOperationPredicate) (result ListWindowsAutopilotDeploymentProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.WindowsAutopilotDeploymentProfileAssignment, 0)

	resp, err := c.ListWindowsAutopilotDeploymentProfileAssignments(ctx, id, options)
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

	result = ListWindowsAutopilotDeploymentProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
