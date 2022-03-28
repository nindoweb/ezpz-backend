package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Queue struct {
	Name string `bson:"name"`
	Payload []byte `bson:"payload"`
	ReservedAt time.Time `bson:"reserved_at"`
	CreatedAt time.Time `bson:"created_at"`
	AvalibleAt time.Time `bson:"avalible_at"`
}

type FailedQueue struct {
	Queue Queue
	FailedAt time.Time
}

func newQueue(name string, data interface{}) *Queue {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	return &Queue{
		Name: name,
		Payload:payload,
		ReservedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

func newFailedQueue(q Queue) *FailedQueue {
	return &FailedQueue{
		Queue: q,
		FailedAt: time.Now(),
	}
}

func (q Queue) String() string {
	return fmt.Sprintf(q.Name)
}

func (fq FailedQueue) String() string {
	return fmt.Sprintf(fq.Queue.Name)
}