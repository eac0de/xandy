package storage

import (
	"context"
	"net/http"

	"github.com/eac0de/xandy/internal/models"
	"github.com/eac0de/xandy/shared/pkg/httperror"
	"github.com/google/uuid"
)

func (s *xandyStorage) InsertUserAuthInfo(ctx context.Context, userAuthInfo *models.UserAuthInfo) error {
	query := `INSERT INTO user_auth_info (id, user_id, name, created_at, updated_at, login, password, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.Exec(
		ctx,
		query,
		userAuthInfo.ID,
		userAuthInfo.UserID,
		userAuthInfo.Name,
		userAuthInfo.CreatedAt,
		userAuthInfo.UpdatedAt,
		userAuthInfo.Login,
		userAuthInfo.Password,
		userAuthInfo.Metadata,
	)
	return err
}

func (s *xandyStorage) UpdateUserAuthInfo(ctx context.Context, userAuthInfo *models.UserAuthInfo) error {
	query := `UPDATE user_auth_info SET name=$3, updated_at=$4, login=$5, password=$6, metadata=$7 WHERE id=$1 AND user_id=$2`
	_, err := s.Exec(
		ctx,
		query,
		userAuthInfo.ID,
		userAuthInfo.UserID,
		userAuthInfo.Name,
		userAuthInfo.UpdatedAt,
		userAuthInfo.Login,
		userAuthInfo.Password,
		userAuthInfo.Metadata,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return httperror.New(err, "UserAuthInfo not found", http.StatusNotFound)
		}
		return err
	}
	return nil
}

func (s *xandyStorage) GetUserAuthInfo(ctx context.Context, dataID uuid.UUID, userID uuid.UUID) (*models.UserAuthInfo, error) {
	query := `SELECT name, created_at, updated_at, login, password, metadata FROM user_auth_info WHERE id=$1 AND user_id=$2`
	userAuthInfo := models.UserAuthInfo{BaseUserData: models.BaseUserData{ID: dataID, UserID: userID}}
	row := s.QueryRow(ctx, query, dataID, userID)
	err := row.Scan(
		&userAuthInfo.Name,
		&userAuthInfo.CreatedAt,
		&userAuthInfo.UpdatedAt,
		&userAuthInfo.Login,
		&userAuthInfo.Password,
		&userAuthInfo.Metadata,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, httperror.New(err, "UserAuthInfo not found", http.StatusNotFound)
		}
		return nil, err
	}
	return &userAuthInfo, nil
}

func (s *xandyStorage) GetUserAuthInfoList(ctx context.Context, userID uuid.UUID, offset int) ([]models.UserAuthInfo, error) {
	query := `SELECT id, name, created_at, updated_at, login, password, metadata FROM user_auth_info WHERE user_id=$1 ORDER BY created_at DESC LIMIT 20 OFFSET $2`

	rows, err := s.Query(ctx, query, userID, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userAuthInfoList []models.UserAuthInfo
	for rows.Next() {
		var userAuthInfo models.UserAuthInfo
		err := rows.Scan(
			&userAuthInfo.ID,
			&userAuthInfo.Name,
			&userAuthInfo.CreatedAt,
			&userAuthInfo.UpdatedAt,
			&userAuthInfo.Login,
			&userAuthInfo.Password,
			&userAuthInfo.Metadata,
		)
		if err != nil {
			return nil, err
		}
		userAuthInfo.UserID = userID
		userAuthInfoList = append(userAuthInfoList, userAuthInfo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userAuthInfoList, nil
}

func (s *xandyStorage) DeleteUserAuthInfo(ctx context.Context, dataID uuid.UUID, userID uuid.UUID) error {
	query := `DELETE FROM user_auth_info WHERE id=$1 AND user_id=$1`
	_, err := s.Exec(ctx, query, dataID, userID)
	return err
}
