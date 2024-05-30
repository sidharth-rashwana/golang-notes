package services

import (
	"errors"
	"fmt"

	internal "github.com/sidharth-rashwana/golang-notes/internal/model"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService(status *bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	query := n.db
	if status != nil {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&notes).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *NotesService) UpdateNotesService(title string, status bool, id int) (*internal.Notes, error) {
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	if err := n.db.Save(&note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return note, nil
}

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}
	if note.Title == "" {
		return nil, errors.New("Title is required")
	}
	if err := n.db.Create(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return note, nil
}

func (n *NotesService) DeleteNotesService(id int64) error {
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func (n *NotesService) GetNoteService(id int64) (*internal.Notes, error) {
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	return note, nil
}
