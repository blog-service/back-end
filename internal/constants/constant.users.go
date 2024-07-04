package constants

const (
	UserRoleAdmin = iota + 1
	UserRoleWriter
	UserRoleReviewer
	UserRoleReader
)

const (
	UserStatusNew = iota + 1
	UserStatusRegistered
	UserStatusBlocked
	UserStatusDeleted
)
