package entity

type Role uint

const (
	TypicalUserRole = iota + 1
	SuperAdminRole
	StaffRole
	CopyWriterRole
	AccountRole
)
