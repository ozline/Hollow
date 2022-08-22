package data

import (
	"hollow/internal/conf"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	snowflake "github.com/bwmarrin/snowflake"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewSnowflake, NewUserRepo, NewForestRepo)

// Data .
type Data struct {
	// TODO wrapped database client

	db   *gorm.DB
	node *snowflake.Node
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, node *snowflake.Node) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		db:   db,
		node: node,
	}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	path := strings.Join([]string{c.Database.Username, ":", c.Database.Password, "@tcp(", c.Database.Address, ":", strconv.FormatInt(c.Database.Port, 10), ")/", c.Database.Dbname, "?charset=utf8&parseTime=True"}, "") //数据库连接文本

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}

	return db
}

func NewSnowflake(c *conf.Data) *snowflake.Node {
	var st time.Time
	st, err := time.Parse("2006-01-02", c.Snowflake.StartTime)
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err := snowflake.NewNode(c.Snowflake.MachineId)
	if err != nil {
		panic(err)
	}
	return node
}
