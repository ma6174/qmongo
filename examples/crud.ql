// connect
s, err = mgo.dial("localhost")
if err != nil {
	panic(err)
}
defer s.Close()

info, err = s.buildInfo()
if err != nil {
	panic(err)
}
printf("%+v\n", info)

coll = s.DB("test").C("test_qmongo")
coll.DropCollection()

// insert
for i=0; i<10; i++ {
	err = coll.Insert({"id":i})
	if err != nil {
		panic(err)
	}
}

// count
n, err = coll.Count()
if err != nil {
	panic(err)
}
if n != 10 {
	panic(n)
}

// find all
all, err = coll.Find(nil).Sort("id").All()
if err != nil {
	panic(err)
}

if len(all) != 10 {
	panic(len(all))
}

// check and update
for id, v = range all {
	println(v._id.Time(), v._id, v.id)
	if v.id != id {
		panic(v.id, id)
	}
	err = coll.Update(map[string]var{"_id":v._id}, {"$set": {"id2":id*id}})
	if err != nil{
		panic(err)
	}
}

// use iter and remove one by one
iter = coll.find(nil).Sort("-id").Iter()
for {
	result, ok = iter.Next()
	if !ok {
		break
	}
	if result.id * result.id != result.id2 {
		panic(result)
	}
	err = coll.RemoveId(result._id)
	if err != nil{
		panic(err)
	}
}
if iter.Err() != nil {
	panic(iter.Err())
}

// final count
count, err = coll.find(nil).count()
if err != nil {
	panic(err)
}
if count != 0{
	panic(count)
}

println("success!")
