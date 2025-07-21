package scheduler

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func runJob(runID int, command string) {
	ctx := context.Background()
	startTime := time.Now()
	_, _ = DB().Exec(ctx, `
		UPDATE job_runs SET status = 'running', started_at = $1 WHERE id = $2
	`, startTime, runID)

	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	out, err := cmd.CombinedOutput()

	status := "success"
	if err != nil {
		status = "failed"
	}
	endTime := time.Now()

	fmt.Println(string(out))

	_, _ = DB().Exec(ctx, `
		UPDATE job_runs SET status = $1, finished_at = $2, output = $3 WHERE id = $4
	`, status, endTime, string(out), runID)

}

func RunScheduledJobs() error {
	nowTime := time.Now()
	rows, err := DB().Query(context.Background(), `
		SELECT jr.id, j.Command
		FROM job_runs jr
		JOIN jobs j ON jr.job_id = j.id
		WHERE jr.status = 'scheduled' AND jr.run_at <= $1
	`, nowTime)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var runID int
		var command string
		if err := rows.Scan(&runID, &command); err != nil {
			return err
		}
		go runJob(runID, command)
	}
	return nil
}

func ManualRerun(jobRunID int) error {
	var command string
	err := DB().QueryRow(context.Background(), `
		SELECT j.command
		FROM job_runs jr
		JOIN jobs j ON jr.job_id = j.id
		WHERE jr.id = $1
	`, jobRunID).Scan(&command)
	if err != nil {
		return fmt.Errorf("job run not found: %w", err)
	}
	go runJob(jobRunID, command)
	return nil
}
