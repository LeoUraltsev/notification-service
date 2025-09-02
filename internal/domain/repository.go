package domain

type Repository interface {
	AdminUsers() ([]int64, error)
}
