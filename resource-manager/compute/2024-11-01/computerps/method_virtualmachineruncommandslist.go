package computerps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRunCommandsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RunCommandDocumentBase
}

type VirtualMachineRunCommandsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RunCommandDocumentBase
}

type VirtualMachineRunCommandsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VirtualMachineRunCommandsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VirtualMachineRunCommandsList ...
func (c ComputeRPSClient) VirtualMachineRunCommandsList(ctx context.Context, id LocationId) (result VirtualMachineRunCommandsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VirtualMachineRunCommandsListCustomPager{},
		Path:       fmt.Sprintf("%s/runCommands", id.ID()),
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
		Values *[]RunCommandDocumentBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VirtualMachineRunCommandsListComplete retrieves all the results into a single object
func (c ComputeRPSClient) VirtualMachineRunCommandsListComplete(ctx context.Context, id LocationId) (VirtualMachineRunCommandsListCompleteResult, error) {
	return c.VirtualMachineRunCommandsListCompleteMatchingPredicate(ctx, id, RunCommandDocumentBaseOperationPredicate{})
}

// VirtualMachineRunCommandsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeRPSClient) VirtualMachineRunCommandsListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate RunCommandDocumentBaseOperationPredicate) (result VirtualMachineRunCommandsListCompleteResult, err error) {
	items := make([]RunCommandDocumentBase, 0)

	resp, err := c.VirtualMachineRunCommandsList(ctx, id)
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

	result = VirtualMachineRunCommandsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
