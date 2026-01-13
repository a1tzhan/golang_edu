INSERT INTO faculty (name) VALUES ('Engineering'), ('Humanities');

INSERT INTO
    uni_group (name, faculty_id)
VALUES ('MATH161', 1),
    ('CSCI152', 1),
    ('BUS101', 2),
    ('HST100', 2);

INSERT INTO
    student (
        name,
        gender,
        birth_date,
        faculty_id,
        group_id
    )
VALUES (
        'Winston Churchill',
        'm',
        '2000-05-15',
        1,
        1
    ),
    (
        'Marie Curie',
        'f',
        '1999-08-22',
        1,
        2
    ),
    (
        'Michael Kalashnikov',
        'm',
        '2001-12-03',
        2,
        3
    ),
    (
        'Ada Lovelace',
        'f',
        '2000-03-30',
        2,
        4
    );

INSERT INTO
    schedule (
        subject,
        group_id,
        faculty_id,
        time_slot
    )
VALUES ('Calculus', 1, 1, '09:00:00'),
    (
        'Data Structures',
        2,
        1,
        '11:00:00'
    ),
    (
        'Business Ethics',
        3,
        2,
        '10:00:00'
    ),
    (
        'World History',
        4,
        2,
        '13:00:00'
    );