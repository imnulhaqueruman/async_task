package tasks 

import (
	"time"

    "github.com/hibiken/asynq"
)

const (
	TypeWelcomeEmail = "email:welocome"

	TypeReminderEmail ="email:Reminder"
)

func NewWelcomeEmailTask(id int) *asynq.Task {
    // Specify task payload.
    payload := map[string]interface{}{
        "user_id": id, // set user ID
    }

    // Return a new task with given type and payload.
    return asynq.NewTask(TypeWelcomeEmail, payload)
}

func NewReminderEmailTask(id int, ts time.Time) *asynq.Task {
    // Specify task payload.
    payload := map[string]interface{}{
        "user_id": id,          // set user ID
        "sent_in": ts.String(), // set time to sending
    }

    // Return a new task with given type and payload.
    return asynq.NewTask(TypeReminderEmail, payload)
}

