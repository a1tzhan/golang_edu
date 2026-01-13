CREATE TABLE attendance (
    ID SERIAL PRIMARY KEY,
    student_id INT REFERENCES student (student_id),
    schedule_id INT REFERENCES schedule (schedule_id),
    status BOOLEAN NOT NULL,
    date DATE NOT NULL;
);