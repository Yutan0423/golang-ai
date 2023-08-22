-- migrate:up
CREATE TABLE questions (
    question_id INT NOT NULL AUTO_INCREMENT,
    content VARCHAR(255) NOT NULL,
    ctime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`question_id`)
);

-- migrate:down
DROP TABLE questions;