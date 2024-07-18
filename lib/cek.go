package lib

var Person = struct{
	Name string
	Age int
}{}


func init() {
	Person.Name = "kuluk kuluk"
	Person.Age = 20

	println("pertama exsekusi >>>>>>>")
}