package users



const (
	table = "users"
	settings = []string {"users" :{"champ":"lastname,firstname,email", "left":{"user_roles":{"on":"users.roleid"}}}}
)


type User struct {
	Id string `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname string `db:"lastname" json:"lastname"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserSettings struct {
	Pool map[string]User
}



