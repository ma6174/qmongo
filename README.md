# qmongo - qlang script support MongoDB and BSON

### install

```
go get -v github.com/ma6174/qmongo
```

### how to use

1. write a [qlang](https://github.com/qiniu/qlang) file like this:

```go
# cat test.ql
s, err = mgo.dial("localhost")
if err != nil {
	panic(err)
}
defer s.Close()
coll = s.DB("test").C("test_qmongo")
err = coll.Insert({"key":"hello qmongo"})
if err != nil {
	panic(err)
}
one, err = coll.Find(nil).One()
if err != nil {
	panic(err)
}
printf("%+v\n", one)
err = coll.RemoveId(one._id)
if err != nil {
	panic(err)
}
```

2. just run it

```
# qmongo test.ql
map[_id:ObjectIdHex("58e1d0cf57ad7a980b77eb6c") key:hello qmongo]
```

more examples please see: [examples](examples)
