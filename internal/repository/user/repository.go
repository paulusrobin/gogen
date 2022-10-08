package user

type Repository interface {
	Creator
	GetterByID
}
