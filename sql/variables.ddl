**************************************************
*
*    Name: variables.ddl
*
*    Author    Date        Comment
*    octavio   2025-07-21  Initial Creation
*
**************************************************

\c scheduler

CREATE TABLE variables (
    id       SERIAL    PRIMARY KEY,
    key      TEXT      NOT NULL,
    value    TEXT      NOT NULL
);
