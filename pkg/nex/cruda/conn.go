package cruda

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"gorm.io/gorm"
)

type CrudConn struct {
	n *nex.Conn

	Db *gorm.DB
}

func NewCrudaConn(conn *nex.Conn) *CrudConn {
	return &CrudConn{
		n: conn,
	}
}
