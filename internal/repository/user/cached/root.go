package cached

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/paulusrobin/gogen/internal/repository/model"
	"github.com/paulusrobin/gogen/internal/repository/user"
	"github.com/paulusrobin/gogen/internal/repository/user/dto"
	"github.com/rs/zerolog/log"
	"time"
)

type implementation struct {
	redis *redis.Client
	repo  user.Repository
}

var defaultUserTTL = 3 * time.Minute

func cacheKey(userID uint) string {
	return fmt.Sprintf(`gogen#user#id#%d`, userID)
}

func (i implementation) Create(ctx context.Context, request dto.CreateRequest) (dto.CreateResponse, error) {
	return i.repo.Create(ctx, request)
}

func (i implementation) GetByID(ctx context.Context, request dto.GetByIDRequest) (dto.GetByIDResponse, error) {
	var (
		key      = cacheKey(request.UserID)
		userData model.User
	)

	if err := i.redis.Get(ctx, key).Scan(&userData); err == nil {
		log.Ctx(ctx).Info().
			Fields(map[string]interface{}{
				"key":  key,
				"user": userData,
				"ttl":  defaultUserTTL,
			}).
			Msg("[user-cached-repository] get user data from redis")
		return dto.GetByIDResponse{User: &userData}, err
	}

	response, err := i.repo.GetByID(ctx, request)
	if err == nil {
		if redisErr := i.redis.Set(ctx, key, *response.User, defaultUserTTL).Err(); redisErr != nil {
			log.Ctx(ctx).Error().Err(redisErr).
				Fields(map[string]interface{}{
					"key":  key,
					"user": response.User,
					"ttl":  defaultUserTTL,
				}).
				Msg("[user-cached-repository] error on set user data to redis")
		}
	}

	return response, err
}

// UserRepository function to return user repository from database
func UserRepository(redis *redis.Client, repo user.Repository) user.Repository {
	return implementation{redis: redis, repo: repo}
}
