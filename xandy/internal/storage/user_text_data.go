package storage

import (
	"context"
	"net/http"

	"github.com/eac0de/xandy/internal/models"
	"github.com/eac0de/xandy/shared/pkg/httperror"
	"github.com/google/uuid"
)

func (s *xandyStorage) InsertUserTextData(ctx context.Context, userTextData *models.UserTextData) error {
	query := `INSERT INTO user_text_data (id, user_id, name, created_at, updated_at, data, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.Exec(ctx, query, userTextData.ID, userTextData.UserID, userTextData.Name, userTextData.CreatedAt, userTextData.UpdatedAt, userTextData.Data, userTextData.Metadata)
	return err
}

func (s *xandyStorage) UpdateUserTextData(ctx context.Context, userTextData *models.UserTextData) error {
	query := `UPDATE user_text_data SET name=$3, updated_at=$4, data=$5, metadata=$6 WHERE id=$1 AND user_id=$2`
	_, err := s.Exec(ctx, query, userTextData.ID, userTextData.UserID, userTextData.Name, userTextData.UpdatedAt, userTextData.Data, userTextData.Metadata)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return httperror.New(err, "UserTextData not found", http.StatusNotFound)
		}
		return err
	}
	return nil
}

func (s *xandyStorage) GetUserTextData(ctx context.Context, dataID uuid.UUID, userID uuid.UUID) (*models.UserTextData, error) {
	query := `SELECT name, created_at, updated_at, data, metadata FROM user_text_data WHERE id=$1 AND user_id=$2`
	row := s.QueryRow(ctx, query, dataID, userID)
	userTextData := models.UserTextData{BaseUserData: models.BaseUserData{ID: dataID, UserID: userID}}
	err := row.Scan(
		&userTextData.Name,
		&userTextData.CreatedAt,
		&userTextData.UpdatedAt,
		&userTextData.Data,
		&userTextData.Metadata,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, httperror.New(err, "UserTextData not found", http.StatusNotFound)
		}
		return nil, err
	}
	return &userTextData, nil
}

func (s *xandyStorage) DeleteUserTextData(ctx context.Context, dataID uuid.UUID, userID uuid.UUID) error {
	query := `DELETE FROM user_text_data WHERE id=$1 AND user_id=$2`
	_, err := s.Exec(ctx, query, dataID, userID)
	return err
}

func (s *xandyStorage) GetUserTextDataList(ctx context.Context, userID uuid.UUID, offset int) ([]models.UserTextData, error) {
	query := `SELECT id, name, created_at, updated_at, data, metadata FROM user_text_data WHERE user_id=$1 ORDER BY created_at DESC LIMIT 20 OFFSET $2`

	rows, err := s.Query(ctx, query, userID, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTextDataList []models.UserTextData
	for rows.Next() {
		var userTextData models.UserTextData
		err := rows.Scan(
			&userTextData.ID,
			&userTextData.Name,
			&userTextData.CreatedAt,
			&userTextData.UpdatedAt,
			&userTextData.Data,
			&userTextData.Metadata,
		)
		if err != nil {
			return nil, err
		}
		userTextData.UserID = userID
		userTextDataList = append(userTextDataList, userTextData)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userTextDataList, nil
}
