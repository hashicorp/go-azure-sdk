package virtualmachineruncommands

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVirtualMachineOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineRunCommand

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByVirtualMachineOperationResponse, error)
}

type ListByVirtualMachineCompleteResult struct {
	Items []VirtualMachineRunCommand
}

func (r ListByVirtualMachineOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByVirtualMachineOperationResponse) LoadMore(ctx context.Context) (resp ListByVirtualMachineOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByVirtualMachineOperationOptions struct {
	Expand *string
}

func DefaultListByVirtualMachineOperationOptions() ListByVirtualMachineOperationOptions {
	return ListByVirtualMachineOperationOptions{}
}

func (o ListByVirtualMachineOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByVirtualMachineOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// ListByVirtualMachine ...
func (c VirtualMachineRunCommandsClient) ListByVirtualMachine(ctx context.Context, id VirtualMachineId, options ListByVirtualMachineOperationOptions) (resp ListByVirtualMachineOperationResponse, err error) {
	req, err := c.preparerForListByVirtualMachine(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByVirtualMachine(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListByVirtualMachineComplete retrieves all of the results into a single object
func (c VirtualMachineRunCommandsClient) ListByVirtualMachineComplete(ctx context.Context, id VirtualMachineId, options ListByVirtualMachineOperationOptions) (ListByVirtualMachineCompleteResult, error) {
	return c.ListByVirtualMachineCompleteMatchingPredicate(ctx, id, options, VirtualMachineRunCommandOperationPredicate{})
}

// ListByVirtualMachineCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VirtualMachineRunCommandsClient) ListByVirtualMachineCompleteMatchingPredicate(ctx context.Context, id VirtualMachineId, options ListByVirtualMachineOperationOptions, predicate VirtualMachineRunCommandOperationPredicate) (resp ListByVirtualMachineCompleteResult, err error) {
	items := make([]VirtualMachineRunCommand, 0)

	page, err := c.ListByVirtualMachine(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByVirtualMachineCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListByVirtualMachine prepares the ListByVirtualMachine request.
func (c VirtualMachineRunCommandsClient) preparerForListByVirtualMachine(ctx context.Context, id VirtualMachineId, options ListByVirtualMachineOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/runCommands", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByVirtualMachineWithNextLink prepares the ListByVirtualMachine request with the given nextLink token.
func (c VirtualMachineRunCommandsClient) preparerForListByVirtualMachineWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByVirtualMachine handles the response to the ListByVirtualMachine request. The method always
// closes the http.Response Body.
func (c VirtualMachineRunCommandsClient) responderForListByVirtualMachine(resp *http.Response) (result ListByVirtualMachineOperationResponse, err error) {
	type page struct {
		Values   []VirtualMachineRunCommand `json:"value"`
		NextLink *string                    `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByVirtualMachineOperationResponse, err error) {
			req, err := c.preparerForListByVirtualMachineWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByVirtualMachine(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "ListByVirtualMachine", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
