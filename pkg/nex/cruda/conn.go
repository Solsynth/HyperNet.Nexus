package cruda

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"gorm.io/gorm"
)

type CrudConn struct {
	Conn *nex.Conn

	db *gorm.DB
}

func NewCrudaConn(conn *nex.Conn) *CrudConn {
	return &CrudConn{
		Conn: conn,
	}
}
