package character

type HitBox struct {
	offset, size [2]int
}

func NewHitBox(offsetX, offsetY, sizeX, sizeY int) HitBox {
	return HitBox{[2]int{offsetX, offsetY}, [2]int{sizeX, sizeY}}
}

func (hitBox *HitBox) Add(x, y int) HitBox {
	return HitBox{[2]int{hitBox.offset[0] + x, hitBox.offset[1] + y}, hitBox.size}
}

func (hitBox *HitBox) Intersects(hitBox2 *HitBox) bool {
	for i := 0; i < 2; i++ {
		if hitBox.offset[i] < hitBox2.offset[i] {
			if hitBox.offset[i]+hitBox.size[i] < hitBox2.offset[i] {
				return false
			}
		} else {
			if hitBox2.offset[i]+hitBox2.size[i] < hitBox.offset[i] {
				return false
			}
		}
	}
	return true
}
