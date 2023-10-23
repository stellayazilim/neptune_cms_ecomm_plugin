package aggregate

import "github.com/stellayazilim/neptune_cms_ecomm_plugin/pkg/entity"

type Category struct {
	root          entity.Category
	subCategories []entity.Category
}
