package swagger

type SubscriptionEventReportResponse struct {
	Subscription      *Subscription `json:"subscription,omitempty"`
	SupportedFeatures string        `json:"supportedFeatures,omitempty"`
	SubscriptionId    string        `json:"subscriptionId,omitempty"`
	ReportList        []ReportList  `json:"reportList,omitempty"`
}

type ReportList struct {
	Type      string   `json:"type,omitempty"`
	State     State    `json:"state,omitempty"`
	TimeStamp string   `json:"timeStamp,omitempty"`
	Supi      string   `json:"supi,omitempty"`
	Location  Location `json:"location,omitempty"`
}

type State struct {
	Active bool `json:"active,omitempty"`
}

type Location struct {
	NrLocation NrLocation `json:"nrLocation,omitempty"`
}

type NrLocation struct {
	Tai                      Tai    `json:"tai,omitempty"`
	Ncgi                     Ncgi   `json:"ncgi,omitempty"`
	AgeOfLocationInformation int32  `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string `json:"ueLocationTimestamp,omitempty"`
}

type Tai struct {
	PlmnId PlmnId `json:"plmnId,omitempty"`
	Tac    string `json:"tac,omitempty"`
}
type Ncgi struct {
	PlmnId   PlmnId `json:"plmnId,omitempty"`
	NrCellId string `json:"nrCellId,omitempty"`
}
type PlmnId struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
}

type Subscription struct {
	Supi                          string      `json:"supi,omitempty"`
	AnyUE                         bool        `json:"anyUE,omitempty"`
	GroupId                       string      `json:"groupId,omitempty"`
	NotifyCorrelationId           string      `json:"notifyCorrelationId,omitempty"`
	EventNotifyUri                string      `json:"eventNotifyUri,omitempty"`
	NfId                          string      `json:"NfId,omitempty"`
	Options                       Options     `json:"options,omitempty"`
	EventList                     []EventList `json:"eventList,omitempty"`
	SubsChangeNotifyUri           string      `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string      `json:"subsChangeNotifyCorrelationId,omitempty"`
}

type Options struct {
	Trigger    string `json:"trigger,omitempty"`
	MaxReports int32  `json:"maxReports,omitempty"`
	Expiry     string `json:"expiry,omitempty"`
	RepPeriod  string `json:"repPeriod,omitempty"`
}

type EventList struct {
	Type          string `json:"type,omitempty"`
	ImmediateFlag bool   `json:"immediateFlag,omitempty"`
	RefId         int32  `json:"refId,omitempty"`
}
