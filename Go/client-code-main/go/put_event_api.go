/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event. Â© 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
	//"strings"
)

func ScsAsIdSubscriptionsSubscriptionIdPut(w http.ResponseWriter, r *http.Request) {
	// urlPart := strings.Split(r.URL.Path, "/")
	// id:=urlPart[len(urlPart)-3]
	// subscriptionId:=urlPart[len(urlPart)-1]

	// for index, updatemonitoringevent := range monitoringevent {
	// 	if updatemonitoringevent.ID == id && updatemonitoringevent.subscriptionId==subscriptionId{
	// 		json.NewDecoder(r.Body).Decode(&updatemonitoringevent)
	// 		monitoringevent[index].ID = updatemonitoringevent.ID
	// 		monitoringevent[index].subscriptionId = updatemonitoringevent.subscriptionId
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Success to update todo"}`))

}