package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/youngconnorr/go-redis-jobqueue/internal/queue"
	"github.com/youngconnorr/go-redis-jobqueue/internal/worker"
)

func main() {

	ctx := context.Background()

	// Intialize Redis from container
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
	})


	// Ping redis to check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	log.Println("Connected to Redis Successfully")

	// Beegin processing jobs...
	numWorkers := 3
	waitGroup := sync.WaitGroup{} 

	for i := 1; i <= numWorkers; i++ {
		waitGroup.Add(1)
		go worker.Start(ctx, rdb, i, &waitGroup)
	}

	time.Sleep(1 * time.Second)

	numJobs := 10

	for i := 1; i <= numJobs; i++ {
		job := queue.Job{
			ID:	  i,
			Name:  fmt.Sprintf("Job-%d", i),
			Payload: fmt.Sprintf("Payload for job %d", i),
		}
		err := queue.EnqueueJob(ctx, rdb, &job)
		if err != nil {
			log.Printf("Failed to enqueue job %d: %v", i, err)
		} else {
			log.Printf("Enqueued job %d: %s", i, job.Name)
		}
	}

	waitGroup.Wait()

	// 
}
