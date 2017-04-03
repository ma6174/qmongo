mongoHost = "127.0.0.1:27017"
if len(os.args) > 1 {
	mongoHost = os.args[1]
}
session, err = mgo.dial(mongoHost)
if err != nil {
	log.Fatal(err)
}
last, err = session.DB("local").C("oplog.rs").Find(nil).Sort("-$natural").One()
if err != nil {
	panic(err)
}
iter = session.DB("local").C("oplog.rs").Find({"ts":map[string]bson.MongoTimestamp{"$gt":last.ts}}).LogReplay().Tail(time.Second*5)
for {
	for {
		result, ok = iter.Next()
		if !ok {
			break
		}
		b, err = json.marshal(result)
		if err != nil {
			panic(err)
		}
		println(string(b))
	}
	if iter.Err() != nil {
		panic(iter.Close())
	}
	if iter.Timeout() {
		continue
	}
	panic("unexpect error")
}
