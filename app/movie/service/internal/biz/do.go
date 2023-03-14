package biz

type MovieDetail struct {
	Rating   float64
	Metadata *Metadata
}

type Movie struct {
	RecordID   string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID     string `json:"userId"`
	Value      uint32 `json:"value"`
}

type Metadata struct {
	Id          uint64
	Title       string
	Description string
	Director    string
}
