package data

import (
	"fmt"
	"hollow/internal/conf"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	snowflake "github.com/bwmarrin/snowflake"

	shortmsg "hollow/internal/pkg/aliyun/shortmsg"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewSnowflake, NewOSS, NewRedis, NewShortMsg, NewUserRepo, NewForestRepo)

// Data .
type Data struct {
	// TODO wrapped database client

	db       *gorm.DB                 // 阿里云RDS
	node     *snowflake.Node          // 雪花
	oss      *oss.Bucket              // 阿里云OSS
	shortmsg *shortmsg.AliyunShortMsg // 阿里云短信服务封装
	dbRedis  redis.Conn               // Redis
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, oss *oss.Bucket, dbRedis redis.Conn, shortmsg *shortmsg.AliyunShortMsg, node *snowflake.Node) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		dbRedis.Close()
		fmt.Println("Redis 连接关闭")
	}

	return &Data{
		db:       db,
		node:     node,
		oss:      oss,
		shortmsg: shortmsg,
		dbRedis:  dbRedis,
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
	fmt.Println("DATABASE CONNECTED")
	return db
}

func NewRedis(c *conf.Data) redis.Conn {
	addr := c.Redis.Address + ":" + strconv.FormatInt(c.Redis.Port, 10)
	account := c.Redis.Username + ":" + c.Redis.Password
	redisdb, err := redis.Dial("tcp", addr, redis.DialDatabase(0), redis.DialPassword(account))

	if err != nil {
		panic(err)
	}

	fmt.Println("Redis 连接成功")
	return redisdb
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

func NewOSS(c *conf.Data) *oss.Bucket {
	Client, err := oss.New(c.OSS.Endpoint, c.OSS.AccessKeyID, c.OSS.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		panic(err)
	}
	Bucket, err := Client.Bucket(c.OSS.Bucket)
	if err != nil {
		panic(err)
	}
	fmt.Println("OSS CONNECTED, SDK VERSION:", oss.Version)
	return Bucket
}

func NewShortMsg(c *conf.Data) *shortmsg.AliyunShortMsg {
	asm := shortmsg.NewAliyunShortMsg()
	err := asm.Init(c.Shortmsg.AccessKeyID, c.Shortmsg.AccessKeySecret, c.Shortmsg.RegionId, c.Shortmsg.Scheme, c.Shortmsg.Signname, c.Shortmsg.Templatecode)

	if err != nil {
		panic(err)
	}

	return asm
}
