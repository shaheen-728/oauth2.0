# {{classname}}

All URIs are relative to *{apiRoot}/3gpp-monitoring-event/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsSubscriptionIdPut**](MonitoringEventAPISubscriptionLevelPUTOperationApi.md#ScsAsIdSubscriptionsSubscriptionIdPut) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource

# **ScsAsIdSubscriptionsSubscriptionIdPut**
> MonitoringEventSubscription ScsAsIdSubscriptionsSubscriptionIdPut(ctx, body, scsAsId, subscriptionId)
Updates/replaces an existing subscription resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitoringEventSubscription**](MonitoringEventSubscription.md)| Parameters to update/replace the existing subscription | 
  **scsAsId** | **string**| Identifier of the SCS/AS | 
  **subscriptionId** | **string**| Identifier of the subscription resource | 

### Return type

[**MonitoringEventSubscription**](MonitoringEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

