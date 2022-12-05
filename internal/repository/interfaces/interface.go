package interfaces

import "ApiAnd/internal/model"

type IBucket interface {
	addAdvertToBucket(advert model.Advert)
	deleteAdvertToBucket(advertId int)
}
