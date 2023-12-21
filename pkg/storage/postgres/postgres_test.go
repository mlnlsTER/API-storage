package postgres

import (
	"reflect"
	"testing"
)

func TestStorage_Tasks(t *testing.T) {
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Tasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetTaskByAuthor(t *testing.T) {
	type args struct {
		authorID int
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTaskByAuthor(tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetTaskByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetTaskByAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetTaskByID(t *testing.T) {
	type args struct {
		taskID int
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTaskByID(tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	type args struct {
		task Task
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type args struct {
		taskID int
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteTask(tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
