package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

// Получать список задач по автору
func (s *Storage) GetTaskByAuthor(authorID int) (Task, error) {
	var t Task
	err := s.db.QueryRow(context.Background(), `
        SELECT 
            id,
            opened,
            closed,
            author_id,
            assigned_id,
            title,
            content
        FROM tasks
        WHERE author_id = $1;
    `, authorID).Scan(
		&t.ID,
		&t.Opened,
		&t.Closed,
		&t.AuthorID,
		&t.AssignedID,
		&t.Title,
		&t.Content,
	)
	return t, err
}

// Получать список задач по метке
func (s *Storage) GetTaskByID(taskID int) (Task, error) {
	var t Task
	err := s.db.QueryRow(context.Background(), `
        SELECT 
            id,
            opened,
            closed,
            author_id,
            assigned_id,
            title,
            content
        FROM tasks
        WHERE id = $1;
    `, taskID).Scan(
		&t.ID,
		&t.Opened,
		&t.Closed,
		&t.AuthorID,
		&t.AssignedID,
		&t.Title,
		&t.Content,
	)
	return t, err
}

// Обновлять задачу по id
func (s *Storage) UpdateTask(task Task) error {
	_, err := s.db.Exec(context.Background(), `
        UPDATE tasks
        SET 
            title = $1,
            content = $2,
            closed = $3,
            assigned_id = $4
        WHERE id = $5;
    `, task.Title, task.Content, task.Closed, task.AssignedID, task.ID)
	return err
}

// Удалять задачу по id
func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Exec(context.Background(), `
        DELETE FROM tasks
        WHERE id = $1;
    `, taskID)
	return err
}
