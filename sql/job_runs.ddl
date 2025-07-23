**************************************************
*
*    Name: job_runs.ddl
*
*    Author    Date        Comment
*    octavio   2025-07-20  Initial Creation
*
**************************************************

\c scheduler

CREATE TABLE job_runs (
    id             SERIAL    PRIMARY KEY,
    job_id         INTEGER   REFERENCE jobs(id),
    run_at         TIMESTAMP        NOT NULL,
    status         TEXT      DEFAULT 'schedule',
    started_at     TIMESTAMP,
    finished_at    TIMESTAMP,
    output         TEXT
);
