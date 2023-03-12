package main

type user struct {
	id       int
	username int //  nama variabel tidak menggunakan tipe data int, melainkan lebih deskriptif seperti string untuk username dan password.
	password int // tipe data variabel []byte
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user { //Fungsi getallusers diubah menjadi GetAllUsers dengan aturan CamelCase.
	return u.t
}

func (u userservice) getuserbyid(id int) user { // Fungsi getuserbyid diubah menjadi GetUserByID.
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
