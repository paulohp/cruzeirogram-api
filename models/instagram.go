package instagram

var Pic struct {
	Task      string    `bson:"t"`
	Created   time.Time `bson:"c"`
	Updated   time.Time `bson:"u,omitempty"`
	Completed time.Time `bson:"cp,omitempty"`
}
