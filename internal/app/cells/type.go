package cells

type Cell struct {
	Drawable  uint32
	X         int
	Y         int
	Alive     bool
	AliveNext bool
}
