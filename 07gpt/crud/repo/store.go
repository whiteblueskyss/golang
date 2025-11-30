package repo

import (
	"crud/models"
	"errors"
	"sync"
)

type Store struct {
	mu       sync.Mutex
	students map[int]models.Student
	nextID   int
}

func NewStore() *Store {
	return &Store{
		students: make(map[int]models.Student),
		nextID:   1,
	}
}

func (s *Store) Create(student models.Student) (models.Student, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	student.ID = s.nextID
	s.nextID++
	s.students[student.ID] = student
	return student, nil
}

func (s *Store) GetAll() ([]models.Student, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]models.Student, 0, len(s.students))
	for _, v := range s.students {
		out = append(out, v)
	}
	return out, nil
}

func (s *Store) GetByID(id int) (models.Student, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	st, ok := s.students[id]
	if !ok {
		return models.Student{}, errors.New("not found")
	}
	return st, nil
}

func (s *Store) Update(id int, student models.Student) (models.Student, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.students[id]
	if !ok {
		return models.Student{}, errors.New("not found")
	}
	student.ID = id
	s.students[id] = student
	return student, nil
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.students[id]
	if !ok {
		return errors.New("not found")
	}
	delete(s.students, id)
	return nil
}
