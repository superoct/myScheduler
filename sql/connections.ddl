**************************************************
*
*    Name: connections.ddl
*
*    Author    Date        Comment
*    octavio   2025-07-21  Initial Creation
*
**************************************************

\c scheduler

CREATE TABLE connections (
    id             SERIAL    PRIMARY KEY,
    name           TEXT      NOT NULL,
    conn_string    TEXT      NOT NULL
);
