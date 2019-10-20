package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)

type (
	Store interface {
		GetDB() *gorm.DB
	}

	Database struct {
		*gorm.DB
	}

	options struct {
		maxIdleConns    int
		maxOpenConns    int
		connMaxLifetime time.Duration
	}

	Option interface {
		apply(*options)
	}

	optionFunc func(*options)
)

func (f optionFunc) apply(o *options) {
	f(o)
}

func MaxOpenConns(count int) Option {
	return optionFunc(func(o *options) {
		o.maxOpenConns = count
	})
}

func MaxIdleConns(count int) Option {
	return optionFunc(func(o *options) {
		o.maxIdleConns = count
	})
}

func ConnMaxLifetime(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.connMaxLifetime = t
	})
}

func NewDatabase(ops ...Option) Store {
	id := viper.Get("database.id")
	pw := viper.Get("database.password")
	url := viper.Get("database.url")
	port := viper.Get("database.port")
	schema := viper.Get("database.schema")
	options := options{
		maxIdleConns:    10,
		maxOpenConns:    100,
		connMaxLifetime: time.Hour,
	}
	_url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", id, pw, url, port, schema)

	db, err := gorm.Open("mysql", _url)
	if err != nil {
		panic(err)
	}

	for _, o := range ops {
		o.apply(&options)
	}

	db.DB().SetMaxIdleConns(options.maxIdleConns)
	db.DB().SetConnMaxLifetime(options.connMaxLifetime)
	db.DB().SetMaxOpenConns(options.maxOpenConns)

	return &Database{db}
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
