//************************************************************************//
// API "cellar": Model Helpers
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

// MediaType Retrieval Functions

// ListBottle returns an array of view: default.
func (m *BottleDB) ListBottle(ctx context.Context) []*app.Bottle {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottle"}, time.Now())

	var native []*Bottle
	var objs []*app.Bottle
	err := m.Db.Scopes().Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottle())
	}

	return objs
}
// BottleToBottle returns the Bottle representation of Bottle.
func (m *Bottle) BottleToBottle() *app.Bottle {
	bottle := &app.Bottle{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	bottle.Account = m.Account.AccountToAccount()
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating
	bottle.Varietal = m.Varietal
	bottle.Vineyard = m.Vineyard
	bottle.Vintage = m.Vintage

	return bottle
}

// OneBottle returns an array of view: default.
//START OMIT
func (m *BottleDB) OneBottle(ctx context.Context, id int) (*app.Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottle"}, time.Now())

	var native Bottle
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottle()
	return &view, err
}
//END OMIT

// MediaType Retrieval Functions

// ListBottleFull returns an array of view: full.
func (m *BottleDB) ListBottleFull(ctx context.Context) []*app.BottleFull {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottlefull"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleFull
	err := m.Db.Scopes().Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleFull())
	}

	return objs
}

// BottleToBottleFull returns the Bottle representation of Bottle.
func (m *Bottle) BottleToBottleFull() *app.BottleFull {
	bottle := &app.BottleFull{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	bottle.Account = m.Account.AccountToAccount()
	bottle.Color = m.Color
	bottle.Country = m.Country
	bottle.CreatedAt = &m.CreatedAt
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating
	bottle.Region = m.Region
	bottle.Review = m.Review
	bottle.Sweetness = m.Sweetness
	bottle.UpdatedAt = &m.UpdatedAt
	bottle.Varietal = m.Varietal
	bottle.Vineyard = m.Vineyard
	bottle.Vintage = m.Vintage

	return bottle
}

// OneBottleFull returns an array of view: full.
func (m *BottleDB) OneBottleFull(ctx context.Context, id int) (*app.BottleFull, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottlefull"}, time.Now())

	var native Bottle
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottleFull()
	return &view, err
}

// MediaType Retrieval Functions

// ListBottleTiny returns an array of view: tiny.
func (m *BottleDB) ListBottleTiny(ctx context.Context) []*app.BottleTiny {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottletiny"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleTiny
	err := m.Db.Scopes().Table(m.TableName()).Preload("Account").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleTiny())
	}

	return objs
}

// BottleToBottleTiny returns the Bottle representation of Bottle.
func (m *Bottle) BottleToBottleTiny() *app.BottleTiny {
	bottle := &app.BottleTiny{}
	tmp1 := m.Account.AccountToAccountLink()
	bottle.Links = &app.BottleLinks{Account: tmp1}
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Rating = &m.Rating

	return bottle
}

// OneBottleTiny returns an array of view: tiny.
func (m *BottleDB) OneBottleTiny(ctx context.Context, id int) (*app.BottleTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottletiny"}, time.Now())

	var native Bottle
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottleTiny()
	return &view, err
}
