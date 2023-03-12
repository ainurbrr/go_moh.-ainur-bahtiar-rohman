package main

type user struct {
	id       int
	username int
	password int
}

//nama struct tidak menggunakan camel case
type userservice struct {
	//nama variabel struct tidak jelas menggambarkan apa
	t []user
}

//nama fungsi tidak menggunakan camel case
func (u userservice) getallusers() []user {
	//returnnya juga tidak jelas menggambarkan apa
	return u.t
}

//nama fungsi tidak menggunakan camel case
func (u userservice) getuserbyid(id int) user {
	//returnnya juga tidak jelas menggambarkan apa
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

//Berapa banyak kekurangan dalam penulisan kode tersebut? 6
//Bagian mana saja terjadi kekurangan tersebut? 'karena tidak menggunakan penulisan camel case dan nama variabel, object serta return yang tidak jelas