package database

import (
	"context"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"github.com/paulusrobin/gogen/internal/repository/user"
	"github.com/paulusrobin/gogen/internal/repository/user/dto"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type implementation struct {
	db *gorm.DB
}

func (i implementation) Create(ctx context.Context, request dto.CreateRequest) (dto.CreateResponse, error) {
	tx := i.db.WithContext(ctx).Model(model.User{}).Create(&request.User)
	if err := tx.Error; err != nil {
		log.Ctx(ctx).Error().Err(err).
			Fields(map[string]interface{}{"user": request.User}).
			Msg("[user-repository] error on create user data")
		return dto.CreateResponse{}, err
	}
	log.Ctx(ctx).Info().
		Fields(map[string]interface{}{"user": request.User}).
		Msg("[user-repository] success create user")
	return dto.CreateResponse{User: request.User}, nil
}

func (i implementation) GetByID(ctx context.Context, request dto.GetByIDRequest) (dto.GetByIDResponse, error) {
	var userData model.User
	tx := i.db.WithContext(ctx).Model(model.User{}).Where("id = ?", request.UserID).Take(&userData)
	if err := tx.Error; err != nil {
		log.Ctx(ctx).Error().Err(err).
			Fields(map[string]interface{}{"user-id": request.UserID}).
			Msg("[user-repository] error on get user data")
		return dto.GetByIDResponse{}, err
	}
	log.Ctx(ctx).Info().
		Fields(map[string]interface{}{"user-id": request.UserID}).
		Msg("[user-repository] success get user by id")
	return dto.GetByIDResponse{User: &userData}, nil
}

// UserRepository function to return user repository from database
func UserRepository(db *gorm.DB) user.Repository {
	return implementation{db: db}
}
