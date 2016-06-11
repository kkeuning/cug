//************************************************************************//
// API "cellar": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma-cellar
// --design=github.com/goadesign/gorma-cellar/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma-cellar/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)
//START OMIT
// Cellar Account
type Account struct {
	ID        int      `gorm:"primary_key"` // primary key
	BottleID  int      // has one Account
	Bottles   []Bottle // has many Bottles
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string
	UpdatedAt time.Time
}
//END OMIT

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Account) TableName() string {
	return "accounts"

}

// AccountDB is the implementation of the storage interface for
// Account.
type AccountDB struct {
	Db gorm.DB
}

// NewAccountDB creates a new storage type.
func NewAccountDB(db gorm.DB) *AccountDB {
	return &AccountDB{Db: db}
}

// DB returns the underlying database.
func (m *AccountDB) DB() interface{} {
	return &m.Db
}

// AccountStorage represents the storage interface.
type AccountStorage interface {
	DB() interface{}
	List(ctx context.Context) []Account
	Get(ctx context.Context, id int) (Account, error)
	Add(ctx context.Context, account *Account) (*Account, error)
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id int) error

	ListAccount(ctx context.Context) []*app.Account
	OneAccount(ctx context.Context, id int) (*app.Account, error)

	ListAccountLink(ctx context.Context) []*app.AccountLink
	OneAccountLink(ctx context.Context, id int) (*app.AccountLink, error)

	ListAccountTiny(ctx context.Context) []*app.AccountTiny
	OneAccountTiny(ctx context.Context, id int) (*app.AccountTiny, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

}

// CRUD Functions

// Get returns a single Account as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountDB) Get(ctx context.Context, id int) (Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "get"}, time.Now())

	var native Account
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Account{}, nil
	}

	return native, err
}

// List returns an array of Account
func (m *AccountDB) List(ctx context.Context) []Account {
	defer goa.MeasureSince([]string{"goa", "db", "account", "list"}, time.Now())

	var objs []Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Account", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *AccountDB) Add(ctx context.Context, model *Account) (*Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Account", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx context.Context, model *Account) error {
	defer goa.MeasureSince([]string{"goa", "db", "account", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "account", "delete"}, time.Now())

	var obj Account

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Account", "error", err.Error())
		return err
	}

	return nil
}
