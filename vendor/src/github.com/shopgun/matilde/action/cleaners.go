package action

import (
	"github.com/shopgun/matilde/db"
	"github.com/shopgun/matilde/events/clean"
	"github.com/shopgun/matilde/events/event"
)

/*
  CleanScopeAciton converts the input go struct containing the raw but  with
  surely  correct types into the data we really want to keep Conversion errors
  hare logged inside each conversion action once the event is cleaned is
  checked for fatal missing fields
*/
func CleanOfferView(
	db db.Connector,
	in *event.Event) (out clean.Offer_View, err Errors) {

	out = clean.Offer_View{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Catalog_id:               in.GetCatalogId(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Currency:                 event.CheckCurrency(in.Currency),
		Dealer_id:                in.Dealer.Value,
		Expires:                  in.Expires(db),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Offer_id:                 in.Offer.Value,
		Price:                    in.GetPrice(db),
		Publish:                  in.PublishTimeStamp(db),
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Run_from:                 in.RunFrom(db),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_birth_year:          in.GetByear(db),
		User_gender:              in.GetGender(db),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanOfferClick(
	db db.Connector,
	in *event.Event) (out clean.Offer_Click, err Errors) {

	out = clean.Offer_Click{
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		Catalog_id:               in.GetCatalogId(db),
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Currency:                 event.CheckCurrency(in.Currency),
		Dealer_id:                in.Dealer.Value,
		Expires:                  in.Expires(db),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Offer_id:                 in.Offer.Value,
		Price:                    in.GetPrice(db),
		Publish:                  in.PublishTimeStamp(db),
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Run_from:                 in.RunFrom(db),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_birth_year:          in.GetByear(db),
		User_gender:              in.GetGender(db),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanOfferList(
	db db.Connector,
	in *event.Event) (out clean.Offer_list, err Errors) {

	out = clean.Offer_list{
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		App_groups:               in.AppGroups(db),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value, Query_string: in.Query_string,
		Limit:             in.Limit,
		Offers:            in.Offers.GetArray(),
		Offset:            in.Offset,
		Request_geohash:   in.Geohash(),
		Request_latitude:  in.Latitude(),
		Request_longitude: in.Longitude(),
		Request_radius:    in.Location.ReqRadius,
		Server_addr:       in.Ipv4Ipv6(in.Server_ip),
		Stores:            in.Stores.GetArray(),
		Timestamp:         in.Times.Timestamp(),
		Type:              in.Type(),
		User_id:           in.User.Id.Value,
		Uuid:              in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanOfferReferral(
	db db.Connector,
	in *event.Event) (out clean.Offer_Referral, err Errors) {

	out = clean.Offer_Referral{
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		Client_addr:        in.Ipv4Ipv6(in.IP),
		App_groups:         in.AppGroups(db),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Dealer_id:          in.Dealer.Value,
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.Id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanOfferSearch(
	db db.Connector,
	in *event.Event) (out clean.Offer_Search, err Errors) {

	out = clean.Offer_Search{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		App_groups:               in.AppGroups(db),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Offers:                   in.Offers.GetArray(),
		Offset:                   in.Offset,
		Query:                    in.Query.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Results:                  in.Results.Value,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Stores:                   in.Stores.GetArray(),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_birth_year:          in.GetByear(db),
		User_gender:              in.GetGender(db),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanOfferSuggest(
	db db.Connector,
	in *event.Event) (out clean.Offer_Suggest, err Errors) {

	out = clean.Offer_Suggest{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		App_groups:               in.AppGroups(db),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Offers:                   in.Offers.GetArray(),
		Offset:                   in.Offset,
		Query:                    in.Query.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Results:                  in.Results.Value,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Stores:                   in.Stores.GetArray(),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopItemAdd(
	db db.Connector,
	in *event.Event) (out clean.Shoppingitem_Add, err Errors) {

	out = clean.Shoppingitem_Add{
		Api_build:                in.Api_build,
		Api_env:                  in.ApiEnv(),
		Api_version:              in.Api,
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_owner:                 in.Is_owner,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Offer_id:                 in.Offer.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopItemEdit(
	db db.Connector,
	in *event.Event) (out clean.Shoppingitem_edit, err Errors) {

	out = clean.Shoppingitem_edit{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		App_groups:               in.AppGroups(db),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Owner:                 in.Is_owner,
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Offer_id:                 in.Offer.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopItemTick(
	db db.Connector,
	in *event.Event) (out clean.Shoppingitem_tick, err Errors) {
	out = clean.Shoppingitem_tick{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Description:              in.Description.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Owner:                 in.Is_owner,
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Offer_id:                 in.Offer.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Tick:                     in.Tick,
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}
func CleanShopItemList(
	db db.Connector,
	in *event.Event) (out clean.Shoppingitem_list, err Errors) {
	out = clean.Shoppingitem_list{
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Owner:                 in.Is_owner,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Shopping_list:            in.Shop_list.Value,
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopListAdd(
	db db.Connector,
	in *event.Event) (out clean.Shoppinglist_add, err Errors) {
	out = clean.Shoppinglist_add{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Owner:                 in.Is_owner,
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopShareAccept(
	db db.Connector,
	in *event.Event) (out clean.Shoppingshare_accept, err Errors) {
	out = clean.Shoppingshare_accept{
		Api_Build:          in.Api_build,
		Api_Version:        in.Api,
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Shopping_list:      in.Shop_list.Value,
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.Id.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanShopShareAdd(
	db db.Connector,
	in *event.Event) (out clean.Shoppingshare_add, err Errors) {
	out = clean.Shoppingshare_add{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Shopping_list:            in.Shop_list.Value,
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanStoreClick(
	db db.Connector,
	in *event.Event) (out clean.Store_click, err Errors) {
	out = clean.Store_click{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Store:                    in.Store.Value,
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanStoreList(
	db db.Connector,
	in *event.Event) (out clean.Store_list, err Errors) {
	out = clean.Store_list{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Limit:                    in.Limit,
		Query_string:             in.Query_string,
		Offer_id:                 in.Offer.Value,
		Offset:                   in.Offset,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Stores:                   in.Stores.GetArray(),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}
func CleanStoreView(
	db db.Connector,
	in *event.Event) (out clean.Store_view, err Errors) {

	out = clean.Store_view{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealer_id:                in.Dealer.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Store:                    in.Store.Value, // NOTE maybe Store_id
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanUserCreate(
	db db.Connector,
	in *event.Event) (out clean.User_create, err Errors) {

	out = clean.User_create{
		Api_Build:          in.Api_build,
		Api_Version:        in.Api,
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Birthyear:          in.BYear.Value64,
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Email:              in.Email,
		User_gender:        in.GetGender(db),
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Locale:             in.Locale,
		Name:               in.Name,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.User_id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanUserLogin(
	db db.Connector,
	in *event.Event) (out clean.User_login, err Errors) {

	out = clean.User_login{
		Api_Build:          in.Api_build,
		Api_Version:        in.Api,
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Email:              in.Email,
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Provider:           in.Provider,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Session_type:       in.Typeof.Value,
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.User_id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanUserLogout(
	db db.Connector,
	in *event.Event) (out clean.User_logout, err Errors) {

	out = clean.User_logout{
		Api_Build:          in.Api_build,
		Api_Version:        in.Api,
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Email:              in.Email,
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Session_type:       in.Typeof.Value,
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.User_id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanUserPwdReset(
	db db.Connector,
	in *event.Event) (out clean.User_pwd_reset, err Errors) {

	out = clean.User_pwd_reset{
		Api_Build:         in.Api_build,
		Api_Version:       in.Api,
		Api_env:           in.ApiEnv(),
		Client_addr:       in.Ipv4Ipv6(in.IP),
		Client_user_agent: in.User_agent,
		Email:             in.Email,
		Is_archive_user:   in.Is_archive_user.Value,
		Is_dealer_admin:   in.Is_Dealer_admin.Value,
		Query_string:      in.Query_string,
		Server_addr:       in.Ipv4Ipv6(in.Server_ip),
		Timestamp:         in.Times.Timestamp(),
		Type:              in.Type(),
		User_id:           in.User.User_id.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanUserVerify(
	db db.Connector,
	in *event.Event) (out clean.User_verify, err Errors) {

	out = clean.User_verify{
		Api_Build:         in.Api_build,
		Api_Version:       in.Api,
		Api_env:           in.ApiEnv(),
		Client_addr:       in.Ipv4Ipv6(in.IP),
		Client_user_agent: in.User_agent,
		Email:             in.Email,
		Is_archive_user:   in.Is_archive_user.Value,
		Is_dealer_admin:   in.Is_Dealer_admin.Value,
		Query_string:      in.Query_string,
		Server_addr:       in.Ipv4Ipv6(in.Server_ip),
		Timestamp:         in.Times.Timestamp(),
		Type:              in.Type(),
		User_id:           in.User.User_id.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogDownload(
	db db.Connector,
	in *event.Event) (out clean.Catalog_download, err Errors) {

	out = clean.Catalog_download{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Catalog_id:               in.Dealer.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_Dealer_admin:          in.Is_Dealer_admin.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogList(
	db db.Connector,
	in *event.Event) (out clean.Catalog_list, err Errors) {

	out = clean.Catalog_list{
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Catalogs:                 in.Catalogs.GetArray(),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value, Query_string: in.Query_string,
		Limit:             in.Limit,
		Offset:            in.Offset,
		Request_geohash:   in.Geohash(),
		Request_latitude:  in.Latitude(),
		Request_longitude: in.Longitude(),
		Request_radius:    in.Location.ReqRadius,
		Server_addr:       in.Ipv4Ipv6(in.Server_ip),
		Stores:            in.Stores.GetArray(),
		TimeRange:         in.Times.Range,
		Timestamp:         in.Times.Timestamp(),
		Type:              in.Type(),
		Typeof:            in.Typeof.Value,
		User_id:           in.User.Id.Value,
		Uuid:              in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogPages(
	db db.Connector,
	in *event.Event) (out clean.Catalog_pages, err Errors) {

	out = clean.Catalog_pages{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Catalog_id:               in.Catalog.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealer_id:                in.Dealer.Value,
		Expires:                  in.Expires(db),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Publish:                  in.PublishTimeStamp(db),
		Run_from:                 in.RunFrom(db),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogPageview(
	db db.Connector,
	in *event.Event) (out clean.Catalog_pageview, err Errors) {

	out = clean.Catalog_pageview{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Catalog_id:               in.Catalog.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealer_id:                in.Dealer.Value,
		Duration:                 in.UDuration(),
		Expires:                  in.Expires(db),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Orientation:              in.Orientation.Value,
		Pages:                    in.Pages.GetArray(),
		Query_string:             in.Query_string,
		Publish:                  in.PublishTimeStamp(db),
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Run_from:                 in.RunFrom(db),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
		View_session:             in.View_session.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogPushSend(
	db db.Connector,
	in *event.Event) (out clean.Catalog_push_send, err Errors) {

	out = clean.Catalog_push_send{
		Api_Build:   in.Api_build,
		Api_Version: in.Api,
		Api_env:     in.ApiEnv(),
		Catalogs:    in.Catalogs.GetArray(),
		Dealers:     in.GetDealers(),
		Endpoint_id: in.Endpoint_id.Value,
		Push_type:   in.Push_type,
		Timestamp:   in.Times.Timestamp(),
		Type:        in.Type(),
		User_id:     in.User.Id.Value,
		Uuid:        in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogPushFetch(
	db db.Connector,
	in *event.Event) (out clean.Catalog_push_fetch, err Errors) {

	out = clean.Catalog_push_fetch{
		Api_Build:          in.Api_build,
		Api_Version:        in.Api,
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Catalogs:           in.Catalogs.GetArray(),
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Dealers:            in.GetDealers(),
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Push_type:          in.Push_type,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.Id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanCatalogView(
	db db.Connector,
	in *event.Event) (out clean.Catalog_view, err Errors) {

	out = clean.Catalog_view{
		Api_Build:                in.Api_build,
		Api_Version:              in.Api,
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Catalog_id:               in.Catalog.Value,
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealer_id:                in.Dealer.Value,
		Expires:                  in.Expires(db),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Publish:                  in.PublishTimeStamp(db),
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Run_from:                 in.RunFrom(db),
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanDealerList(
	db db.Connector,
	in *event.Event) (out clean.Dealer_list, err Errors) {

	out = clean.Dealer_list{
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealers:                  in.GetDealers(),
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Limit:                    in.Limit,
		Offset:                   in.Offset,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanDealerFav(
	// NOTE recent events have dealer_id instead of dealer
	db db.Connector,
	in *event.Event) (out clean.Dealer_favorite, err Errors) {

	out = clean.Dealer_favorite{
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		Client_addr:        in.Ipv4Ipv6(in.IP),
		App_groups:         in.AppGroups(db),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Dealer_id:          in.Dealerid.Value,
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.Id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

func CleanDealerUnFav(
	// NOTE recent events have dealer_id instead of dealer
	db db.Connector,
	in *event.Event) (out clean.Dealer_unfavorite, err Errors) {
	out = clean.Dealer_unfavorite{
		Api_env:            in.ApiEnv(),
		App_id:             in.App_id.Value,
		App_groups:         in.AppGroups(db),
		Client_addr:        in.Ipv4Ipv6(in.IP),
		Client_app_version: in.ApiAppVersion(),
		Client_user_agent:  in.User_agent,
		Dealer_id:          in.Dealerid.Value,
		Is_archive_user:    in.Is_archive_user.Value,
		Is_dealer_admin:    in.Is_Dealer_admin.Value,
		Is_uuid_ephemeral:  in.Is_uuid_ephemeral.Value,
		Query_string:       in.Query_string,
		Server_addr:        in.Ipv4Ipv6(in.Server_ip),
		Timestamp:          in.Times.Timestamp(),
		Type:               in.Type(),
		User_id:            in.User.Id.Value,
		Uuid:               in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}

// NOTE recent events have dealer_id instead of dealer
func CleanDealerView(
	db db.Connector,
	in *event.Event) (out clean.Dealer_view, err Errors) {
	out = clean.Dealer_view{
		Api_env:                  in.ApiEnv(),
		App_id:                   in.App_id.Value,
		App_groups:               in.AppGroups(db),
		Client_addr:              in.Ipv4Ipv6(in.IP),
		Client_app_version:       in.ApiAppVersion(),
		Client_user_agent:        in.User_agent,
		Dealer_id:                in.Dealer.Value,
		Is_archive_user:          in.Is_archive_user.Value,
		Is_dealer_admin:          in.Is_Dealer_admin.Value,
		Is_user_defined_location: in.Location.GeoCoded(),
		Is_uuid_ephemeral:        in.Is_uuid_ephemeral.Value,
		Query_string:             in.Query_string,
		Request_geohash:          in.Geohash(),
		Request_latitude:         in.Latitude(),
		Request_longitude:        in.Longitude(),
		Request_radius:           in.Location.ReqRadius,
		Server_addr:              in.Ipv4Ipv6(in.Server_ip),
		Timestamp:                in.Times.Timestamp(),
		Type:                     in.Type(),
		User_id:                  in.User.Id.Value,
		Uuid:                     in.User.Uuid.Value,
	}
	err = eventChecker(in, out)
	return out, err
}
