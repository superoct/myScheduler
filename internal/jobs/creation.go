package jobs

import (
	"context"
	"fmt"
	"time"

	"scheduler/internal/db"
)

func CreateJob(
	name string,
	command string,
	dependencies []string,
	scheduleType string,
	scheduleStartTimeStr string,
	repeatEvery int,
	repeatAmount int,
	connectionID int,
	parentID int,
	variableIDs []int,
) error {
	scheduleStartTime, err := time.Parse("15:04:05", scheduleStartTimeStr)
	if err != nil {
		return fmt.Errorf("invalid schedule_start_time format: %w", err)
	}

	_, err = db.DB().Exec(context.Background(), `
		INSERT INTO jobs (
			name, command, dependencies, schedule_type, schedule_start_time, repeat_every,
			repeat_amount, connection_id, parent_id, variable_ids
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, name, command, dependencies, scheduleType, scheduleStartTime, repeatEvery,
		repeatAmount, connectionID, parentID, variableIDs)
	if err != nil {
		return fmt.Errorf("job couldn't be created: %w", err)
	}
	return nil
}
