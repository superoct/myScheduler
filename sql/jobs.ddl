**************************************************
*
*    Name: jobs.ddl
*
*    Author    Date        Comment
*    octavio   2025-07-20  Initial Creation
*
**************************************************

\c scheduler

CREATE TABLE jobs (
    id                     SERIAL PRIMARY KEY,
    name                   TEXT   UNIQUE NOT NULL,
    command                TEXT          NOT NULL,
    dependencies           TEXT[] DEFAULT '{}',
    schedule_type          TEXT CHECK (schedule_type IN ('hourly', 'daily', 'weekly', 'yearly')),
    schedule_start_time    TIME,
    repeat_every           INTEGER,
    repeat_amount          INTEGER,
    connection_id          INTEGER,
    parent_id              INTEGER,
    variable_ids           INTEGER[]    DEFAULT '{}'
);
