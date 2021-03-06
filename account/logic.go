package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger log.Logger
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger,"method","CreateUser")
	user := User{
		Email:    email,
		Password: password,
	}
	if err := s.repostory.CreateUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create user", user.Email)
	return "success", nil
}


func (s service) GetUser(ctx context.Context, id string) (User, error) {
	logger := log.With(s.logger, "method", "GetUser")

	user, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return User{}, err
	}

	logger.Log("Get user", id)

	return user, nil
}

func (s service) Login(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger,"method","Login")
	token, err := s.repostory.Login(ctx, email, password)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("Login",token)
	return token, nil
}


func (s service) CreateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string)(string, error) {
	logger := log.With(s.logger,"method","CreateConfOwnProfile")
	conf_profile := ConfOwnProfile{
		Reviews:   reviews,
		Reads:     reads,
		Useful:    useful,
		Attend:    attend,
		Favourite: favourite,
		Picture:   picture,
	}
	if _,err := s.repostory.CreateConfOwnProfile(ctx,conf_profile); err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create user", conf_profile.Picture)
	return "success", nil
}

func (s service) GetConfOwnProfile(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetConfOwnProfile")
	name, err := s.repostory.GetConfOwnProfile(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateConfOwnProfile(ctx context.Context, reviews int16,reads int16,useful int16,attend int16,favourite string,picture string) (string, error) {
	logger := log.With(s.logger,"method","UpdateConfOwnProfile")
	conf_profile := ConfOwnProfile{
		Reviews:   reviews,
		Reads:     reads,
		Useful:    useful,
		Attend:    attend,
		Favourite: favourite,
		Picture:   picture,
	}
	name, err := s.repostory.UpdateConfOwnProfile(ctx, conf_profile)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteConfOwnProfile(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteConfOwnProfile")
	name, err := s.repostory.DeleteConfOwnProfile(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating) (string, error) {
	logger := log.With(s.logger,"method","CreateConference")
	conference := Conference{
		Name:        name,
		Website:     website,
		About:       about,
		Phone:       phone,
		Email:       email,
		Address:     address,
		City:        city,
		ZipCode:     zipcode,
		Speakers:    speakers,
		Facebook:    facebook,
		Twitter:     twitter,
		Instagram:   instagram,
		OrganizerID: organizer_id,
		Locations:   locations,
		Rating:      rating,
	}
	conf, err := s.repostory.CreateConference(ctx, conference)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return conf, nil
}

func (s service) GetConference(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetConference")
	name, err := s.repostory.GetConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateConference(ctx context.Context, name string,website string, about string,	phone string,email string,address string,city string,zipcode string,speakers []Speaker,	facebook string,twitter string,	instagram string,organizer_id int16, locations []Location, rating Rating) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateConference")
	conference := Conference{
		Name:        name,
		Website:     website,
		About:       about,
		Phone:       phone,
		Email:       email,
		Address:     address,
		City:        city,
		ZipCode:     zipcode,
		Speakers:    speakers,
		Facebook:    facebook,
		Twitter:     twitter,
		Instagram:   instagram,
		OrganizerID: organizer_id,
		Locations:   locations,
		Rating:      rating,
	}
	name, err := s.repostory.UpdateCreateConference(ctx, conference)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateConference(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateConference")
	name, err := s.repostory.DeleteCreateConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateLocation(ctx context.Context, name string, date string, time string) (string, error) {
	logger := log.With(s.logger,"method","CreateLocation")
	location := Location{
		Name:  name,
		Date:  date,
		Time:  time,
	}
	_, err := s.repostory.CreateLocation(ctx, location)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create location")
	return location.Name, nil
}

func (s service) GetLocation(ctx context.Context, id string) (Location, error) {
	logger := log.With(s.logger,"method","GetLocation")
	location, err := s.repostory.GetLocation(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return Location{}, err
	}
	logger.Log("get location", location)
	return location, nil
}

func (s service) UpdateCreateLocation(ctx context.Context, name string, date string, time string) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateLocation")
	location := Location{
		Name:  name,
		Date:  date,
		Time:  time,
	}
	name, err := s.repostory.UpdateCreateLocation(ctx, location)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateLocation(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateLocation")
	name, err := s.repostory.DeleteCreateConference(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool) (string, error) {
	logger := log.With(s.logger,"method","CreateRating")
	rating := Rating{
		Rate:         rate,
		Comment:      comment,
		Caption:      caption,
		AttendQ:      attend_q,
		EnjoyQ:       enjoy_q,
		LocationQ:    location_q,
		ConnectPeerQ: connectPeer_q,
		Awesome:      awesome,
		Great:        great,
		Average:      average,
		Poor:         poor,
		Terrible:     terrible,
		Favorite:     favorite,
		Like:         like,
	}
	_, err := s.repostory.CreateRating(ctx, rating)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "name", nil
}

func (s service) GetRating(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetRating")
	name, err := s.repostory.GetRating(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateRating(ctx context.Context, rate int16, comment string, caption string, attend_q bool, enjoy_q bool, location_q bool, connectPeer_q bool, awesome int16, great int16, average int16, poor int16, terrible int16, favorite bool, like bool) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateRating")
	rating := Rating{
		Rate:         rate,
		Comment:      comment,
		Caption:      caption,
		AttendQ:      attend_q,
		EnjoyQ:       enjoy_q,
		LocationQ:    location_q,
		ConnectPeerQ: connectPeer_q,
		Awesome:      awesome,
		Great:        great,
		Average:      average,
		Poor:         poor,
		Terrible:     terrible,
		Favorite:     favorite,
		Like:         like,
	}
	name, err := s.repostory.UpdateCreateRating(ctx, rating)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateRating(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateRating")
	name, err := s.repostory.DeleteCreateRating(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool) (string, error) {
	logger := log.With(s.logger,"method","CreateReport")
	report := Report{

		Offensive:     offensive,
		Violence:      violence,
		Spam:          spam,
		InAppropriate: in_appropriate,
	}
	_, err := s.repostory.CreateReport(ctx, report)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category")
	return "name", nil
}

func (s service) GetReport(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger,"method","GetReport")
	name, err := s.repostory.GetReport(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCreateReport(ctx context.Context, offensive bool,violence bool,spam bool,in_appropriate bool) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateReport")
	report := Report{

		Offensive:     offensive,
		Violence:      violence,
		Spam:          spam,
		InAppropriate: in_appropriate,
	}
	name, err := s.repostory.UpdateCreateReport(ctx, report)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateReport(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","GetCategory")
	name, err := s.repostory.DeleteCreateReport(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) CreateSpeaker(ctx context.Context, name string, position string) (string, error) {
	logger := log.With(s.logger,"method","CreateSpeaker")
	speaker := Speaker{
		Name:     name,
		Position: position,
	}
	_, err := s.repostory.CreateSpeaker(ctx, speaker)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create speaker")
	return "name", nil
}

func (s service) GetSpeaker(ctx context.Context, id string) (Speaker, error) {
	logger := log.With(s.logger,"method","GetSpeaker")
	name, err := s.repostory.GetSpeaker(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return Speaker{}, err
	}
	logger.Log("get speaker", name)
	return name, nil
}

func (s service) GetAllSpeaker(ctx context.Context)([]Speaker, error) {
	logger := log.With(s.logger,"method","get speakers")
	speaker, err := s.repostory.GetAllSpeaker(ctx)
	if err != nil{
		level.Error(logger).Log("err", err)
		return []Speaker{}, err
	}
	return speaker, nil
}

func (s service) UpdateCreateSpeaker(ctx context.Context, name string, position string) (string, error) {
	logger := log.With(s.logger,"method","UpdateCreateSpeaker")
	speaker := Speaker{
		Name:     name,
		Position: position,
	}
	name, err := s.repostory.UpdateCreateSpeaker(ctx, speaker)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) DeleteCreateSpeaker(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","DeleteCreateSpeaker")
	name, err := s.repostory.DeleteCreateSpeaker(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}
func (s service) GetCategory(ctx context.Context, id string) ([]Category, error) {
	logger := log.With(s.logger,"method","GetCategory")
	cat, err := s.repostory.GetCategory(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return []Category{}, err
	}
	logger.Log("get category", cat)
	return cat, nil
}
func (s service) GetCategories(ctx context.Context) ([]Category, error) {
	logger := log.With(s.logger,"method","GetCategory")
	cat, err := s.repostory.GetCategories(ctx)
	if err != nil{
		level.Error(logger).Log("err",err)
		return []Category{}, err
	}
	logger.Log("get category", cat)
	return cat, nil
}
func (s service) UpdateCategory(ctx context.Context, name string) (string, error) {
	logger := log.With(s.logger, "method","UpdateCategory")
	name, err := s.repostory.UpdateCategory(ctx, name)
	if err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("update category", name)
	return name, nil
}

func (s service) CreateCategory(ctx context.Context, name string)(string, error){
	logger := log.With(s.logger,"method","CreateCategory")
	category := Category{
		Name:  name,
	}
	if err := s.repostory.CreateCategory(ctx, category); err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create user", category.Name)
	return "success", nil
}
func NewService(rep Repository, logger log.Logger) Service{
	return &service{
		repostory: rep,
		logger:    logger,
	}
}