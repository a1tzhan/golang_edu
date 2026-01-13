CREATE TABLE faculty (
    faculty_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE uni_group (
    group_id SERIAL PRIMARY KEY,
    name VARCHAR(10) NOT NULL,
    faculty_id INT REFERENCES faculty (faculty_id)
);

CREATE TABLE student (
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    gender CHAR(1) CHECK (gender IN ('m', 'f')),
    birth_date DATE NOT NULL,
    faculty_id INT REFERENCES faculty (faculty_id),
    group_id INT REFERENCES uni_group (group_id)
);

CREATE TABLE schedule (
    schedule_id SERIAL PRIMARY KEY,
    subject VARCHAR(50) NOT NULL,
    group_id INT REFERENCES uni_group (group_id),
    faculty_id INT REFERENCES faculty (faculty_id),
    time_slot TIME NOT NULL
);