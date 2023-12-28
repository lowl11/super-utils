package models

import "time"

const (
	IndividualsAttr         = "individuals"
	IndividualDocumentsAttr = "individual_documents"
	IndividualContactsAttr  = "individual_contacts"
	LegalEntity             = "legal_entity"
)

type Record struct {
	GoldenId            string           `json:"goldenID"`
	Sid                 string           `json:"sid"`
	System              string           `json:"system"`
	Source              string           `json:"source"`
	BusinessTime        time.Time        `json:"businessTS"`
	Entity              string           `json:"entity"`
	Attribute           string           `json:"attribute"`
	ValueToBe           string           `json:"valueToBe"`
	ValuePrev           string           `json:"valuePrev"`
	ManagerId           string           `json:"managerId"`
	InvalidKeyAttribute bool             `json:"invalidAttribute"`
	KafkaID             KafkaMessageInfo `json:"kafkaID"`
	KafkaField          string           `json:"kafkaField"`
	CreatedTS           time.Time        `json:"CreatedTS"`
	SidEntity           string           `json:"sidEntity"`
	CustomerId          string           `json:"customerId"`
	LegalEntityId       string           `json:"legalEntityId"`
}

type KafkaMessageInfo struct {
	Topic     string `json:"topic"`
	Partition string `json:"partition"`
	Offset    string `json:"offset"`
}

type Customer struct {
	Sid    string `json:"sid"`
	Entity string `json:"entity"`
	System string `json:"system"`
}
