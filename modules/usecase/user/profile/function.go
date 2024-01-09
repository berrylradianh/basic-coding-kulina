package profile

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	p "basic-coding-kulina/helper/password"
	ut "basic-coding-kulina/modules/entity/user"
)

func (pc *profileUsecase) GetAllUserProfile(user *[]ut.User) error {
	return pc.profileRepo.GetAllUserProfile(user)
}

func (pc *profileUsecase) GetUserProfile(user *ut.User, id string) error {
	return pc.profileRepo.GetUserProfile(user, id)
}

func (pc *profileUsecase) GetUserDetailProfile(userDetail *ut.UserDetail, id string) error {
	return pc.profileRepo.GetUserDetailProfile(userDetail, id)
}

func (pc *profileUsecase) UpdateUserProfile(userRequest *ut.UserRequest, id string) error {
	return pc.profileRepo.UpdateUserProfile(userRequest, id)
}

func (pc *profileUsecase) UpdateUserDetailProfile(userDetailRequest *ut.UserDetailRequest, id string) error {
	return pc.profileRepo.UpdateUserDetailProfile(userDetailRequest, id)
}

func (pc *profileUsecase) CreateAddressProfile(address *ut.UserAddress) error {
	return pc.profileRepo.CreateAddressProfile(address)
}

func (pc *profileUsecase) GetAllAddressProfileNoPagination(address *[]ut.UserAddress, idUser string) error {
	return pc.profileRepo.GetAllAddressProfileNoPagination(address, idUser)
}

func (pc *profileUsecase) GetAllAddressProfile(address *[]ut.UserAddress, idUser string, offset, pageSize int) (*[]ut.UserAddress, int64, error) {
	return pc.profileRepo.GetAllAddressProfile(address, idUser, offset, pageSize)
}

func (pc *profileUsecase) GetAddressByIdProfile(address *ut.UserAddress, idUser string, idAddress int) error {
	return pc.profileRepo.GetAddressByIdProfile(address, idUser, idAddress)
}

func (pc *profileUsecase) UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser string) error {
	return pc.profileRepo.UpdateAddressPrimaryProfile(address, idUser)
}

func (pc *profileUsecase) UpdateAddressByIdProfile(addressRequest *ut.UserAddressRequest, idUser string, idAddress int) error {
	return pc.profileRepo.UpdateAddressByIdProfile(addressRequest, idUser, idAddress)
}

func (pc *profileUsecase) UpdatePasswordProfile(user *ut.User, oldPassword string, newPassword string, id string) (string, error) {
	if err := p.VerifyPassword(user.Password, oldPassword); err != nil {
		return "Password salah", err
	}

	hashNewPassword, err := p.HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	return "", pc.profileRepo.UpdatePasswordProfile(string(hashNewPassword), id)
}

func (pc *profileUsecase) GetAllProvince() ([]ut.Province, error) {
	apiKey := "8bb5248063ed493d90aac0311f8a3edb"
	url := "https://api.rajaongkir.com/starter/province"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", apiKey)

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	var provinceResponse ut.ProvinceResponse
	if err := json.Unmarshal(body, &provinceResponse); err != nil {
		return nil, err
	}

	var provinces []ut.Province
	for _, prov := range provinceResponse.RajaOngkir.Results {
		province := ut.Province{
			ProvinceId:   prov.ProvinceId,
			ProvinceName: prov.Province,
		}

		provinces = append(provinces, province)
	}

	return provinces, nil
}

func (pc *profileUsecase) GetAllCityByProvince(provinceId string) ([]ut.City, error) {
	apiKey := "8bb5248063ed493d90aac0311f8a3edb"
	url := "https://api.rajaongkir.com/starter/city?province=" + provinceId

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", apiKey)

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	var cityResponse ut.CityResponse
	if err := json.Unmarshal(body, &cityResponse); err != nil {
		return nil, err
	}

	var cities []ut.City
	for _, city := range cityResponse.RajaOngkir.Results {
		province := ut.City{
			CityId:   city.CityId,
			CityName: city.CityName,
		}

		cities = append(cities, province)
	}

	return cities, nil
}
