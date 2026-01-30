-- Populate users
INSERT INTO users (email, password_hash) VALUES
('john.doe@university.edu', 'hash123'),
('jane.smith@student.edu', 'hash456'),
('alice.jones@student.edu', 'hash789'),
('bob.brown@student.edu', 'hash101'),
('charlie.davis@student.edu', 'hash112');

-- Populate profiles (1 Instructor, 4 Students)
INSERT INTO profiles (user_id, first_name, last_name, role) VALUES
(1, 'John', 'Doe', 'instructor'),
(2, 'Jane', 'Smith', 'student'),
(3, 'Alice', 'Jones', 'student'),
(4, 'Bob', 'Brown', 'student'),
(5, 'Charlie', 'Davis', 'student');

-- Populate courses (All taught by the instructor)
INSERT INTO courses (name, description, instructor_id, time_slot) VALUES
('CSCI151', 'Basics of CS', 1, 'Mon 10:00-12:00'),
('CSCI241', 'SQL and Relational DBs', 1, 'Tue 14:00-16:00'),
('CSCI261', 'Data Structures and Algos', 1, 'Wed 10:00-12:00'),
('CSCI361', 'Processes and Threads', 1, 'Thu 14:00-16:00'),
('CSCI315', 'Web Development', 1, 'Fri 10:00-12:00');

-- Populate enrollments
INSERT INTO enrollments (student_id, course_id) VALUES
(2, 1), -- Jane in CS
(3, 1), -- Alice in CS
(4, 2), -- Bob in DBs
(5, 3), -- Charlie in Algos
(2, 5); -- Jane in Web Dev

-- Populate grades
INSERT INTO grades (enrollment_id, grade) VALUES
(1, 85),
(2, 90),
(3, 78),
(4, 92),
(5, 88);

-- Populate attendances
INSERT INTO attendances (enrollment_id, date, status) VALUES
(1, '2023-09-01', 'present'),
(2, '2023-09-01', 'present'),
(3, '2023-09-01', 'absent'),
(4, '2023-09-01', 'late'),
(5, '2023-09-01', 'present');

-- Populate assignments
INSERT INTO assignments (course_id, title, weight, description, deadline) VALUES
(1, 'CS Homework 1', 10, 'Introductory exercises', '2023-09-10'),
(2, 'DB Project', 20, 'Design a schema', '2023-10-15'),
(3, 'Algo Quiz', 15, 'Sorting algorithms', '2023-09-20'),
(4, 'OS Lab', 25, 'Implement a scheduler', '2023-11-01'),
(5, 'Web App', 30, 'Build a react app', '2023-12-01');

-- Populate submissions
INSERT INTO submissions (assignment_id, student_id, file_path, grade, feedback) VALUES
(1, 2, '/submissions/jane_cs_hw1.pdf', 90, 'Good job'),
(1, 3, '/submissions/alice_cs_hw1.pdf', 85, 'Nice work'),
(2, 4, '/submissions/bob_db_project.zip', 80, 'Well designed'),
(3, 5, '/submissions/charlie_algo_quiz.txt', 95, 'Perfect'),
(5, 2, '/submissions/jane_web_app.zip', 92, 'Great UI');
