/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event. Â© 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MonitoringNotification struct {
	Subscription string `json:"subscription"`
	// Each element identifies a notification of grouping configuration result.
	//ConfigResults []Ts29122CommonDataYamlcomponentsschemasConfigResult `json:"configResults,omitempty"`
	// Monitoring event reports.
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports,omitempty"`
	// Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise.
	CancelInd bool `json:"cancelInd,omitempty"`
	AppliedParam *AppliedParameterConfiguration `json:"appliedParam,omitempty"`
}
