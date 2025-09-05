package mock

// Реализовать базу в которой будут храниться chatid для рассылки, id пользователя в сервисе и статусах
// 1. id админов, по которым будут рассылки о новых регистрациях (возможно отправка сообщений о крит ошибках в сервисах)
// 2. id простых пользователей, кто прошел авторизацию, с рассылкой о статусах, новых входах, и других уведомлений (только уведомления касающиеся пользователя)
type Repo struct {
	users map[int64]int64
}

func New() *Repo {
	return &Repo{
		users: map[int64]int64{
			0: 1597142092,
		},
	}
}

func (r *Repo) AdminUsers() ([]int64, error) {
	res := make([]int64, 0, len(r.users))
	for _, u := range r.users {
		res = append(res, u)
	}
	return res, nil
}
