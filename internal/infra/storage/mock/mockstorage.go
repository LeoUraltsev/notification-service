package mock

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
