package domain

type Computer struct {
	InventoryNumber int
	HDDVolume      int
	RAMVolume      int
	Vendor         string
	Core           Core
	Owner          Owner
}

//NewComputer return empty Computer's instance
func NewComputer() Computer {
	return Computer{}
}
