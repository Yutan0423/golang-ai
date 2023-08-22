-- migrate:up
CREATE TABLE answers (
    answer_id INT PRIMARY KEY AUTO_INCREMENT,
    question_id INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    ctime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES questions(question_id)
);

-- migrate:down
DROP TABLE answers;
