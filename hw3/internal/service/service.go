package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"hw3/internal/model"
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