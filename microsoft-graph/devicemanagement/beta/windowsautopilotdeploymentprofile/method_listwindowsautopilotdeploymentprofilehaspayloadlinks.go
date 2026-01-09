package windowsautopilotdeploymentprofile

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

type ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HasPayloadLinkResultItem
}

type ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HasPayloadLinkResultItem
}

type ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions() ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions {
	return ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions{}
}

func (o ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsAutopilotDeploymentProfileHasPayloadLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsAutopilotDeploymentProfileHasPayloadLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsAutopilotDeploymentProfileHasPayloadLinks - Invoke action hasPayloadLinks
func (c WindowsAutopilotDeploymentProfileClient) ListWindowsAutopilotDeploymentProfileHasPayloadLinks(ctx context.Context, input ListWindowsAutopilotDeploymentProfileHasPayloadLinksRequest, options ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions) (result ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListWindowsAutopilotDeploymentProfileHasPayloadLinksCustomPager{},
		Path:          "/deviceManagement/windowsAutopilotDeploymentProfiles/hasPayloadLinks",
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
		Values *[]beta.HasPayloadLinkResultItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsAutopilotDeploymentProfileHasPayloadLinksComplete retrieves all the results into a single object
func (c WindowsAutopilotDeploymentProfileClient) ListWindowsAutopilotDeploymentProfileHasPayloadLinksComplete(ctx context.Context, input ListWindowsAutopilotDeploymentProfileHasPayloadLinksRequest, options ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions) (ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteResult, error) {
	return c.ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteMatchingPredicate(ctx, input, options, HasPayloadLinkResultItemOperationPredicate{})
}

// ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsAutopilotDeploymentProfileClient) ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteMatchingPredicate(ctx context.Context, input ListWindowsAutopilotDeploymentProfileHasPayloadLinksRequest, options ListWindowsAutopilotDeploymentProfileHasPayloadLinksOperationOptions, predicate HasPayloadLinkResultItemOperationPredicate) (result ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteResult, err error) {
	items := make([]beta.HasPayloadLinkResultItem, 0)

	resp, err := c.ListWindowsAutopilotDeploymentProfileHasPayloadLinks(ctx, input, options)
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

	result = ListWindowsAutopilotDeploymentProfileHasPayloadLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
