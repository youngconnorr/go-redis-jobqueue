package queue

// for redis context
// converting Go structs -> JSON
// printing



type Job struct {
	ID  	int 	`json:"id"` // unique Job id
	Name 	string  `json:"name"` // name of the Job
	Payload string  `json:"payload"` // job payload
}

