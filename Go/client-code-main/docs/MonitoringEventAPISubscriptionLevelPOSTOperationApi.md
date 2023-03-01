# {{classname}}

All URIs are relative to *{apiRoot}/3gpp-monitoring-event/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsPost**](MonitoringEventAPISubscriptionLevelPOSTOperationApi.md#ScsAsIdSubscriptionsPost) | **Post** /{scsAsId}/subscriptions | Creates a new subscription resource for monitoring event notification

# **ScsAsIdSubscriptionsPost**
> MonitoringEventReport ScsAsIdSubscriptionsPost(ctx, body, scsAsId)
Creates a new subscription resource for monitoring event notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitoringEventSubscription**](MonitoringEventSubscription.md)| Subscription for notification about monitoring event | 
  **scsAsId** | **string**| Identifier of the SCS/AS | 

### Return type

[**MonitoringEventReport**](MonitoringEventReport.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

