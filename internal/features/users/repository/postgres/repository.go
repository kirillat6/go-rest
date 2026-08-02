package users_postgres_repository

import (
	core_pgx_pool "github.com/kirillat6/go-rest/internal/core/repository/postgres/pool/pgx"
)

type UsersRepository struct {
	pool core_pgx_pool.Pool
}

func NewUsersRepository(
	pool core_pgx_pool.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
