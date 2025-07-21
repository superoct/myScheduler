package scheduler

import (
	"context"
	"fmt"
	"time"
)

func ScheduleJob(jobName string, runAtStr string) error {
	runAt, err := time.Parse(time.RFC3339, runAtStr)
	if err != nil {
		return fmt.Errorf("invalid run_at format: %w", err)
	}

	var jobID int
	err = DB().QueryRow(context.Background(), "SELECT id FROM jobs WHERE name = $1", jobName).Scan(&jobID)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	_, err = DB().Exec(context.Background(), `
		INSERT INTO job_runs (job_id, run_at, status)
		VALUES ($1, $2, 'scheduled')
	`, jobID, runAt)
	return err
}
