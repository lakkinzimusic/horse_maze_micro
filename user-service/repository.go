package main

import (
	"context"

	"github.com/jinzhu/gorm"
	pb "github.com/lakkinzimusic/horse_maze_micro/user-service/proto/user"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*pb.User, error)
	Get(ctx context.Context, id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(ctx context.Context, id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
