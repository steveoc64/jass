package shared

type Product struct {
	ID               int    `db:"id"`
	CatID            int    `db:"cat_id"`
	Name             string `db:"name"`
	Descr            string `db:"descr"`
	Image            string `db:"image"`
	VolumeML         int    `db:"volume_ml"`
	WeightG          int    `db:"weight_g"`
	ShippingVolumeML int    `db:"shipping_volume_ml"`
	ShippingWeightG  int    `db:"shipping_weight_g"`
	ShippingCode     int    `db:"shipping_code"`
	ImageURL         string
}

func (t *Product) GetImageURL() string {
	return "/img/product/" + t.Image
}

func (t *Product) GetNameDesc() string {
	return t.Name + " - " + t.Descr
}
