package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MuhammedAshifVnr/user_service/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// UserRepository defines methods for interacting with the user database
type UserRepository interface {
	GetUserByID(id uint) (*models.User, error)
	GetUsersByIDs(ids []uint) ([]models.User, error)
	SearchUsers(city, phone, query string, married bool, limit, offset int) ([]models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{db: db, cache: redis}
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUsersByIDs(ids []uint) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) SearchUsers(city, phone, query string, married bool, limit, offset int) ([]models.User, error) {
	cacheKey := fmt.Sprintf("search:city:%s:phone:%s:query:%s:married:%v:limit:%d:offset:%d", city, phone, query, married, limit, offset)
	ctx := context.Background()

	// Check cache
	if cachedData, err := r.cache.Get(ctx, cacheKey).Result(); err == nil {
		var users []models.User
		if err := json.Unmarshal([]byte(cachedData), &users); err == nil {
			return users, nil
		}
		// Log cache unmarshalling error for debugging, but do not return
		fmt.Printf("Cache unmarshal error: %v\n", err)
	}

	// Build query with full-text search
	queryBuilder := r.db.Model(&models.User{})
	if city != "" {
		queryBuilder = queryBuilder.Where("city = ?", city)
	}
	if phone != "" {
		queryBuilder = queryBuilder.Where("phone = ?", phone)
	}
	if query != "" {
		// Use plainto_tsquery for natural language input
		queryBuilder = queryBuilder.Where("search_vector @@ plainto_tsquery(?)", query)
	}
	queryBuilder = queryBuilder.Where("married = ?", married).Limit(limit).Offset(offset)

	// Execute query
	var users []models.User
	if err := queryBuilder.Find(&users).Error; err != nil {
		// Log the database error for debugging
		fmt.Printf("Database query error: %v\n", err)
		return nil, err
	}

	// Cache results if query is successful
	if len(users) > 0 {
		data, err := json.Marshal(users)
		if err == nil {
			// Cache the data for 5 minutes
			if err := r.cache.Set(ctx, cacheKey, data, 5*time.Minute).Err(); err != nil {
				fmt.Printf("Cache set error: %v\n", err)
			}
		} else {
			fmt.Printf("Cache marshal error: %v\n", err)
		}
	}

	return users, nil
}


func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	fmt.Println(&user)
	return nil
}
