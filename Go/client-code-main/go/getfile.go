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
	//"encoding/json"
	// "strings"
	// "github.com/gorilla/mux"
	//"fmt"
)

func ScsAsIdSubscriptionsGet(w http.ResponseWriter, r *http.Request) {
	// urlPart := strings.Split(r.URL.Path, "/")
	// id:=urlPart[len(urlPart)-2]

	// var getmonitoringevent []MonitoringEventSubscription =[]MonitoringEventSubscription{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// for _, monitoringeventid := range monitoringevent {
	// 	if monitoringeventid.MtcProviderId == id {
	// 		getmonitoringevent = append(getmonitoringevent,monitoringeventid)

	// 	}
	// }
	// json.NewEncoder(w).Encode(getmonitoringevent)
	w.WriteHeader(http.StatusOK)
}
