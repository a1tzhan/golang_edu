package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"hw4/internal/model"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{pool: db}
}

func (s *Service) GetStudentByID(id string) (*model.Student, error) {
	row := s.pool.QueryRow(context.Background(), "SELECT * FROM student WHERE student_id=$1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Gender, &student.BirthDate, &student.FacultyID, &student.GroupID)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *Service) GetAllSchedule() ([]model.Schedule, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM schedule")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.Subject, &schedule.GroupID, &schedule.FacultyID, &schedule.TimeSlot)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *Service) GetGroupScheduleByID(id string) ([]model.Schedule, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM schedule WHERE group_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupSchedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.Subject, &schedule.GroupID, &schedule.FacultyID, &schedule.TimeSlot)
		if err != nil {
			return nil, err
		}
		groupSchedules = append(groupSchedules, schedule)
	}

	if len(groupSchedules) == 0 {
		return nil, errors.New("No schedule found for the specified group")
	}
	return groupSchedules, nil
}

func (s *Service) GetSubjectAttendance(id string) ([]model.Attendance, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM attendance WHERE schedule_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjectAttendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.ScheduleID, &attendance.Status, &attendance.Date)
		if err != nil {
			return nil, err
		}
		subjectAttendances = append(subjectAttendances, attendance)
	}

	return subjectAttendances, nil
}

func (s *Service) GetStudentAttendance(id string) ([]model.Attendance, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM attendance WHERE student_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentAttendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.ScheduleID, &attendance.Status, &attendance.Date)
		if err != nil {
			return nil, err
		}
		studentAttendances = append(studentAttendances, attendance)
	}

	return studentAttendances, nil
}

func (s *Service) PostSubjectAttendance(attendance *model.Attendance) (*model.Attendance, error) {
	// Validate required fields
	if attendance.StudentID == 0 {
		return nil, errors.New("student_id is required")
	}
	if attendance.ScheduleID == 0 {
		return nil, errors.New("schedule_id is required")
	}
	if attendance.Date.IsZero() {
		return nil, errors.New("date is required")
	}

	// Check if student exists
	var studentExists bool
	err := s.pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM student WHERE student_id=$1)",
		attendance.StudentID).Scan(&studentExists)
	if err != nil {
		return nil, err
	}
	if !studentExists {
		return nil, errors.New("student not found")
	}

	// Check if schedule exists
	var scheduleExists bool
	err = s.pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM schedule WHERE schedule_id=$1)",
		attendance.ScheduleID).Scan(&scheduleExists)
	if err != nil {
		return nil, err
	}
	if !scheduleExists {
		return nil, errors.New("schedule not found")
	}

	// Check for duplicate attendance record
	var duplicateExists bool
	err = s.pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM attendance WHERE student_id=$1 AND schedule_id=$2 AND date=$3)",
		attendance.StudentID, attendance.ScheduleID, attendance.Date).Scan(&duplicateExists)
	if err != nil {
		return nil, err
	}
	if duplicateExists {
		return nil, errors.New("attendance record already exists for this student, schedule, and date")
	}

	// Insert attendance record and return the created record with ID
	query := `
		INSERT INTO attendance (student_id, schedule_id, status, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, student_id, schedule_id, status, date`

	var created model.Attendance
	err = s.pool.QueryRow(context.Background(), query,
		attendance.StudentID,
		attendance.ScheduleID,
		attendance.Status,
		attendance.Date).Scan(
		&created.ID,
		&created.StudentID,
		&created.ScheduleID,
		&created.Status,
		&created.Date)

	if err != nil {
		return nil, err
	}

	return &created, nil
}