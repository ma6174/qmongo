r = os.stdin
if len(os.args)>=2{
	r, err = os.open(os.args[1])
	if err != nil{
		panic(err)
	}
}
for {
	m, err = bson.unmarshal(r)
	if err != nil{
		if err == io.EOF{
			break
		}
		panic(err)
	}
	// println(m._id)
	b, err = json.marshal(m)
	if err != nil{
		panic(err)
	}
	println(string(b))
}
