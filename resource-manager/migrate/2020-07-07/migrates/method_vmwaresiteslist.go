package migrates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareSitesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMwareSite
}

type VMwareSitesListCompleteResult struct {
	Items []VMwareSite
}

// VMwareSitesList ...
func (c MigratesClient) VMwareSitesList(ctx context.Context, id commonids.ResourceGroupId) (result VMwareSitesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.OffAzure/vmwareSites", id.ID()),
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
		Values *[]VMwareSite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VMwareSitesListComplete retrieves all the results into a single object
func (c MigratesClient) VMwareSitesListComplete(ctx context.Context, id commonids.ResourceGroupId) (VMwareSitesListCompleteResult, error) {
	return c.VMwareSitesListCompleteMatchingPredicate(ctx, id, VMwareSiteOperationPredicate{})
}

// VMwareSitesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigratesClient) VMwareSitesListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate VMwareSiteOperationPredicate) (result VMwareSitesListCompleteResult, err error) {
	items := make([]VMwareSite, 0)

	resp, err := c.VMwareSitesList(ctx, id)
	if err != nil {
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

	result = VMwareSitesListCompleteResult{
		Items: items,
	}
	return
}
