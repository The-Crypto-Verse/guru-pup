package lib

var (
	Shrimpy = "1023961419550560387"
	Crabby  = "1023961659921928232"
	Octo    = "1023961928193802352"
	Fish    = "1023962077733343333"
	Dolphin = "1023962224982773791"
	Shark   = "1023962422937137192"
	Whale   = "1023960917333000283"
)

func GetRoles(nfts int) []string {
	roles := []string{}

	if nfts <= 10 {
		roles = append(roles, Shrimpy)
	}
	if nfts > 10 && nfts <= 25 {
		roles = append(roles, Shrimpy, Crabby)
	}
	if nfts > 25 && nfts <= 100 {
		roles = append(roles, Shrimpy, Crabby, Octo)
	}
	if nfts > 100 && nfts <= 250 {
		roles = append(roles, Shrimpy, Crabby, Octo, Fish)
	}
	if nfts > 250 && nfts <= 500 {
		roles = append(roles, Shrimpy, Crabby, Octo, Fish, Dolphin)
	}
	if nfts > 500 && nfts <= 1000 {
		roles = append(roles, Shrimpy, Crabby, Octo, Fish, Dolphin, Shark)
	}
	if nfts > 1000 {
		roles = append(roles, Shrimpy, Crabby, Octo, Fish, Dolphin, Shark, Whale)
	}

	return roles
}
