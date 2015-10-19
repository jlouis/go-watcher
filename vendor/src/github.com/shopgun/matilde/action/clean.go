package action

import (
	"fmt"
	"time"

	"github.com/shopgun/matilde/db"
	"github.com/shopgun/matilde/events/clean"
)

// Clean dispatches the event pointer to the correct cleaner.
func (e *Event) Clean(db db.Connector) {
	in := &e.Raw
	if in.Type() != e.Type {
		err := fmt.Errorf("Event type mismatch")
		e.LogError(err)
		e.BadEvent = true
		return
	}
	var err Errors
	switch {
	case e.Type == "offer.view":
		var out clean.Offer_View
		out, err = CleanOfferView(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "offer.click":
		var out clean.Offer_Click
		out, err = CleanOfferClick(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "offer.list":
		var out clean.Offer_list
		out, err = CleanOfferList(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "offer.referral":
		var out clean.Offer_Referral
		out, err = CleanOfferReferral(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "offer.search":
		var out clean.Offer_Search
		out, err = CleanOfferSearch(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "offer.suggest":
		var out clean.Offer_Suggest
		out, err = CleanOfferSuggest(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingitem.add":
		var out clean.Shoppingitem_Add
		out, err = CleanShopItemAdd(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingitem.edit":
		var out clean.Shoppingitem_edit
		out, err = CleanShopItemEdit(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingitem.list":
		var out clean.Shoppingitem_list
		out, err = CleanShopItemList(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingitem.tick":
		var out clean.Shoppingitem_tick
		out, err = CleanShopItemTick(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppinglist.add":
		var out clean.Shoppinglist_add
		out, err = CleanShopListAdd(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingshare.accept":
		var out clean.Shoppingshare_accept
		out, err = CleanShopShareAccept(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "shoppingshare.add":
		var out clean.Shoppingshare_add
		out, err = CleanShopShareAdd(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "store.click":
		var out clean.Store_click
		out, err = CleanStoreClick(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "store.list":
		var out clean.Store_list
		out, err = CleanStoreList(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "store.view":
		var out clean.Store_view
		out, err = CleanStoreView(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "user.create":
		var out clean.User_create
		out, err = CleanUserCreate(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "user.login":
		var out clean.User_login
		out, err = CleanUserLogin(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "user.logout":
		var out clean.User_logout
		out, err = CleanUserLogout(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "user.password_reset":
		var out clean.User_pwd_reset
		out, err = CleanUserPwdReset(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "user.verify":
		var out clean.User_verify
		out, err = CleanUserVerify(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.download":
		var out clean.Catalog_download
		out, err = CleanCatalogDownload(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.list":
		var out clean.Catalog_list
		out, err = CleanCatalogList(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.pages":
		var out clean.Catalog_pages
		out, err = CleanCatalogPages(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.pageview":
		var out clean.Catalog_pageview
		out, err = CleanCatalogPageview(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.push_send":
		var out clean.Catalog_push_send
		out, err = CleanCatalogPushSend(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.push_fetch":
		var out clean.Catalog_push_fetch
		out, err = CleanCatalogPushFetch(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "catalog.view":
		var out clean.Catalog_view
		out, err = CleanCatalogView(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "dealer.list":
		var out clean.Dealer_list
		out, err = CleanDealerList(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "dealer.favorite":
		var out clean.Dealer_favorite
		out, err = CleanDealerFav(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "dealer.unfavorite":
		var out clean.Dealer_unfavorite
		out, err = CleanDealerUnFav(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	case e.Type == "dealer.view":
		var out clean.Dealer_view
		out, err = CleanDealerView(db, in)
		e.Cleaned = out
		e.Timestamp = time.Unix(0, int64(out.Timestamp*1e+9))
	default:
		if e.Type != "" {
			var err error
			err = fmt.Errorf(fmt.Sprintf("Type: %v is an unknown event type", e.Type))
			e.LogError(err)
		} else {
			err := fmt.Errorf(fmt.Sprintf("Danger: unpredicted fuckup!"))
			e.LogError(err)
		}
	}
	// now we append
	e.LogErrors(err)
	// finally we tag the event as bad if
	// we have at least one error
	if len(e.Errors) > 0 {
		e.BadEvent = true
	}
}
