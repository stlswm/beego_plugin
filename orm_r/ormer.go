package orm_r

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
)

type Ormer struct {
	orm.Ormer
	transactionNo uint
}

// begin transaction
func (o *Ormer) Begin() error {
	o.transactionNo++
	if o.transactionNo == 1 {
		return o.Ormer.Begin()
	}
	return nil
}

// commit transaction
func (o *Ormer) Commit() error {
	o.transactionNo--
	if o.transactionNo == 0 {
		return o.Ormer.Commit()
	}
	return nil
}

// rollback transaction
func (o *Ormer) Rollback() error {
	o.transactionNo--
	if o.transactionNo == 0 {
		return o.Ormer.Rollback()
	}
	return nil
}

// NewOrm create new orm_r
func NewOrm() *Ormer {
	o := &Ormer{}
	o.Ormer = orm.NewOrm()
	return o
}

// NewOrmWithDB create a new ormer object with specify *sql.DB for query
func NewOrmWithDB(driverName, aliasName string, db *sql.DB) (*Ormer, error) {
	o := &Ormer{}
	beeOrm, err := orm.NewOrmWithDB(driverName, aliasName, db)
	if err != nil {
		return nil, err
	}
	o.Ormer = beeOrm
	return o, nil
}
