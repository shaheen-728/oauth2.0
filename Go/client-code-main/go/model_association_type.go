/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event. Â© 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Possible values are - IMEI: The value shall be used when the change of IMSI-IMEI association shall be detected - IMEISV: The value shall be used when the change of IMSI-IMEISV association shall be detected
type AssociationType struct {
	AssociationType string `json:"associationType"`
}
