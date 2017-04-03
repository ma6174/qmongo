package mgo

import (
	"time"

	"gopkg.in/mgo.v2"

	qlang "qlang.io/spec"
)

// ===================================================

func Dial(url string) (*Session, error) {
	s, err := mgo.Dial(url)
	return &Session{s}, err
}

func DialWithInfo(info *mgo.DialInfo) (*Session, error) {
	s, err := mgo.DialWithInfo(info)
	return &Session{s}, err
}

func DialWithTimeout(url string, timeout time.Duration) (*Session, error) {
	s, err := mgo.DialWithTimeout(url, timeout)
	return &Session{s}, err
}

type Session struct {
	*mgo.Session
}

func (s *Session) Run(cmd interface{}) (result interface{}, err error) {
	err = s.Session.Run(cmd, &result)
	return
}

func (s *Session) DB(db string) *Database {
	return &Database{s.Session.DB(db)}
}

func (s *Session) BuildInfo() (info *mgo.BuildInfo, err error) {
	info2, err := s.Session.BuildInfo()
	return &info2, err
}

// ===================================================

type Database struct {
	*mgo.Database
}

func (d *Database) Run(cmd interface{}) (result interface{}, err error) {
	err = d.Database.Run(cmd, &result)
	return
}

func (d *Database) C(name string) *Collection {
	return &Collection{d.Database.C(name)}
}

// ===================================================

type Collection struct {
	*mgo.Collection
}

func (c *Collection) Find(query interface{}) *Query {
	return &Query{c.Collection.Find(query)}
}

func (c *Collection) FindId(id interface{}) *Query {
	return &Query{c.Collection.FindId(id)}
}

func (c *Collection) Pipe(pipeline interface{}) *Pipe {
	return &Pipe{c.Collection.Pipe(pipeline)}
}

// ===================================================

type Query struct {
	*mgo.Query
}

func (q *Query) All() (results []interface{}, err error) {
	err = q.Query.All(&results)
	return
}

func (q *Query) Apply(change mgo.Change) (result interface{}, info *mgo.ChangeInfo, err error) {
	info, err = q.Query.Apply(change, &result)
	return
}

func (q *Query) Distinct(key string) (result interface{}, err error) {
	err = q.Query.Distinct(key, &result)
	return
}

func (q *Query) Explain() (result interface{}, err error) {
	err = q.Query.Explain(&result)
	return
}

func (q *Query) For(f func() error) (result interface{}, err error) {
	err = q.Query.For(&result, f)
	return
}

func (q *Query) MapReduce(job *mgo.MapReduce) (result interface{}, info *mgo.MapReduceInfo, err error) {
	info, err = q.Query.MapReduce(job, &result)
	return
}

func (q *Query) One() (result interface{}, err error) {
	err = q.Query.One(&result)
	return
}

func (q *Query) Comment(comment string) *Query {
	return &Query{q.Query.Comment(comment)}
}

func (q *Query) Hint(indexKey ...string) *Query {
	return &Query{q.Query.Hint(indexKey...)}
}

func (q *Query) Limit(n int) *Query {
	return &Query{q.Query.Limit(n)}
}

func (q *Query) LogReplay() *Query {
	return &Query{q.Query.LogReplay()}
}

func (q *Query) Prefetch(p float64) *Query {
	return &Query{q.Query.Prefetch(p)}
}

func (q *Query) Select(selector interface{}) *Query {
	return &Query{q.Query.Select(selector)}
}

func (q *Query) SetMaxScan(n int) *Query {
	return &Query{q.Query.SetMaxScan(n)}
}

func (q *Query) SetMaxTime(d time.Duration) *Query {
	return &Query{q.Query.SetMaxTime(d)}
}

func (q *Query) Skip(n int) *Query {
	return &Query{q.Query.Skip(n)}
}

func (q *Query) Snapshot() *Query {
	return &Query{q.Query.Snapshot()}
}

func (q *Query) Sort(fields ...string) *Query {
	return &Query{q.Query.Sort(fields...)}
}

func (q *Query) Iter() *Iter {
	return &Iter{q.Query.Iter()}
}

func (q *Query) Tail(timeout time.Duration) *Iter {
	return &Iter{q.Query.Tail(timeout)}
}

// ===================================================

type Iter struct {
	*mgo.Iter
}

func (iter *Iter) All() (results []interface{}, err error) {
	err = iter.Iter.All(&results)
	return
}

func (iter *Iter) For(f func() error) (result interface{}, err error) {
	err = iter.Iter.For(&result, f)
	return
}

func (iter *Iter) Next() (result interface{}, b bool) {
	b = iter.Iter.Next(&result)
	return
}

// ===================================================

type Pipe struct {
	*mgo.Pipe
}

func (p *Pipe) All() (results []interface{}, err error) {
	err = p.Pipe.All(&results)
	return
}

// ===================================================

var Exports = map[string]interface{}{
	"_name": "gopkg.in/mgo.v2",

	// var
	"ErrCursor":   mgo.ErrCursor,
	"ErrNotFound": mgo.ErrNotFound,

	// mode
	"Eventual":  mgo.Eventual,
	"Monotonic": mgo.Monotonic,
	"Strong":    mgo.Strong,

	// role
	"RoleRoot":         mgo.RoleRoot,
	"RoleRead":         mgo.RoleRead,
	"RoleReadAny":      mgo.RoleReadAny,
	"RoleReadWrite":    mgo.RoleReadWrite,
	"RoleReadWriteAny": mgo.RoleReadWriteAny,
	"RoleDBAdmin":      mgo.RoleDBAdmin,
	"RoleDBAdminAny":   mgo.RoleDBAdminAny,
	"RoleUserAdmin":    mgo.RoleUserAdmin,
	"RoleUserAdminAny": mgo.RoleUserAdminAny,
	"RoleClusterAdmin": mgo.RoleClusterAdmin,

	// func
	"dial":            Dial,
	"dialWithInfo":    DialWithInfo,
	"dialWithTimeout": DialWithTimeout,
	"getStats":        mgo.GetStats,
	"isDup":           mgo.IsDup,
	"parseURL":        mgo.ParseURL,
	"resetStatus":     mgo.ResetStats,
	"setDebug":        mgo.SetDebug,
	"setLogger":       mgo.SetLogger,
	"setStatus":       mgo.SetStats,

	// exported structs
	"Change":     qlang.StructOf((*mgo.Change)(nil)),
	"Credential": qlang.StructOf((*mgo.Credential)(nil)),
	"DBRef":      qlang.StructOf((*mgo.DBRef)(nil)),
	"DialInfo":   qlang.StructOf((*mgo.DialInfo)(nil)),
	"Index":      qlang.StructOf((*mgo.Index)(nil)),
	"MapReduce":  qlang.StructOf((*mgo.MapReduce)(nil)),
	"Role":       qlang.StructOf((*mgo.Role)(nil)),
	"Safe":       qlang.StructOf((*mgo.Safe)(nil)),
	"ServerAddr": qlang.StructOf((*mgo.ServerAddr)(nil)),
	"User":       qlang.StructOf((*mgo.User)(nil)),
}
