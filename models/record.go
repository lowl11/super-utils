package models

const (
	IndividualsAttr         = "individuals"
	IndividualDocumentsAttr = "individual_documents"
	IndividualContactsAttr  = "individual_contacts"
	LegalEntity             = "legal_entity"
)

type Record struct {
	GoldenId            string           `json:"goldenID,omitempty"`
	CreatedTS           string           `json:"createdTS"`
	KafkaField          string           `json:"kafkaField"`
	Sid                 string           `json:"sid"`
	System              string           `json:"system"`
	Source              string           `json:"source"`
	BusinessTime        string           `json:"businessTS"`
	Entity              string           `json:"entity"`
	Attribute           string           `json:"attribute"`
	ValueToBe           string           `json:"valueToBe"`
	ValuePrev           string           `json:"valuePrev"`
	InvalidKeyAttribute bool             `json:"invalidAttribute"`
	Customer            Customer         `json:"customer"`
	ManagerId           string           `json:"managerId"`
	KafkaID             KafkaMessageInfo `json:"kafkaID"`
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
