package seed

import (
	ut "basic-coding-kulina/modules/entity/user"
	"time"
)

func CreateUserAddress() []*ut.UserAddress {
	userAddress := []*ut.UserAddress{
		{
			ID:           "47750cc7-9868-46e5-8545-7ed87550920c",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Recipient:    "Supplier",
			Phone:        "08917283129283",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "255",
			CityName:     "Malang",
			Address:      "Sukun",
			Note:         "Rumah Kuning",
			Mark:         "Rumah",
			IsPrimary:    true,
			UserId:       "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			ID:           "414354cc-45a5-47f6-879e-bfb4476eb0d6",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Recipient:    "Ibu user1",
			Phone:        "085123456789",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "247",
			CityName:     "Madiun",
			Address:      "Balerejo",
			Note:         "Rumah cat krem",
			Mark:         "Rumah",
			IsPrimary:    true,
			UserId:       "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			ID:           "c8b974da-ccca-475f-8d2f-c8c505e152fa",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Recipient:    "Satpam user1",
			Phone:        "085123456789",
			ProvinceId:   "5",
			ProvinceName: "DI Yogyakarta",
			CityId:       "39",
			CityName:     "Sleman",
			Address:      "Gentan",
			Note:         "Titip ke mas kos",
			Mark:         "Rumah",
			IsPrimary:    false,
			UserId:       "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
	}

	return userAddress
}
