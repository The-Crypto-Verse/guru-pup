package lib

var (
	shrimpy = "1023961419550560387"
	crabby  = "1023961659921928232"
	octo    = "1023961928193802352"
	fish    = "1023962077733343333"
	dolphin = "1023962224982773791"
	shark   = "1023962422937137192"
	whale   = "1023960917333000283"
)

func GetRoles(nfts int) []string {
	roles := []string{}

	if nfts <= 10 {
		roles = append(roles, shrimpy)
	}
	if nfts > 10 && nfts <= 25 {
		roles = append(roles, shrimpy, crabby)
	}
	if nfts > 25 && nfts <= 100 {
		roles = append(roles, shrimpy, crabby, octo)
	}
	if nfts > 100 && nfts <= 250 {
		roles = append(roles, shrimpy, crabby, octo, fish)
	}
	if nfts > 250 && nfts <= 500 {
		roles = append(roles, shrimpy, crabby, octo, fish, dolphin)
	}
	if nfts > 500 && nfts <= 1000 {
		roles = append(roles, shrimpy, crabby, octo, fish, dolphin, shark)
	}
	if nfts > 1000 {
		roles = append(roles, shrimpy, crabby, octo, fish, dolphin, shark, whale)
	}

	return roles
}
