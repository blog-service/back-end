package constants

const (
	UserRoleAdmin = iota + 1
	UserRoleWriter
	UserRoleReviewer
	UserRoleReader
)

var MapUserRoles = map[int]string{
	UserRoleAdmin:    "admin",
	UserRoleWriter:   "writer",
	UserRoleReviewer: "reviewer",
	UserRoleReader:   "reader",
}

const (
	UserStatusNew = iota + 1
	UserStatusRegistered
	UserStatusBlocked
	UserStatusDeleted
)
