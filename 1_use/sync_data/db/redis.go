package db

import (
    "errors"
    "github.com/FZambia/sentinel"
    "github.com/gomodule/redigo/redis"
    "goPush/K12User/config"
    "time"
)

var pool *redis.Pool

func NewSentinelPool(rc config.RedisConf) {
    sntnl := &sentinel.Sentinel{
        Addrs:      rc.IP,         // []string{"10.202.80.117:17400", "10.202.80.118:17400"},
        MasterName: rc.MasterName, // "master7400",
        Dial: func(addr string) (redis.Conn, error) {
            timeout := 500 * time.Millisecond
            c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
            if err != nil {
                return nil, err
            }
            return c, nil
        },
    }
    pool = &redis.Pool{
        MaxIdle:     3,
        MaxActive:   64,
        Wait:        true,
        IdleTimeout: 240 * time.Second,
        Dial: func() (redis.Conn, error) {
            masterAddr, err := sntnl.MasterAddr()
            if err != nil {
                return nil, err
            }

            c, err := redis.Dial("tcp", masterAddr)
            if err != nil {
                return nil, err
            }

            password := rc.Pwd // "yb@ops"
            db := rc.DbIndex

            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }
            if _, err := c.Do("SELECT", db); err != nil {
                c.Close()
                return nil, err
            }
            return c, nil
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if !sentinel.TestRole(c, "master") {
                return errors.New("Role check failed")
            } else {
                return nil
            }
        },
    }
}

func GetConn() redis.Conn {
    return pool.Get()
}
