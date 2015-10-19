// clean contains all the type information about the cleaned events.
// null fileds are  nil of each field's type.
// i.e. null ip ::ffff:0:0
// NOTE: api_build is mostly empty but we defined it on the spec
// probably meant to be some vaild git commmit identifier
package clean

type Catalog_download struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Catalog_id               int32
	Is_user_defined_location bool
	Is_archive_user          bool
	Is_Dealer_admin          bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Catalog_list struct {
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Catalogs                 []int32
	Offset                   int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Stores                   []int32
	Timestamp                float64
	TimeRange                string
	Type                     string
	Typeof                   string
	User_id                  int32
	Uuid                     string
}

type Catalog_pages struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Catalog_id               int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Expires                  float64
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Publish                  float64
	Query_string             string
	Run_from                 float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Catalog_pageview struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Catalog_id               int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Duration                 float64
	Expires                  float64
	Is_archive_user          bool
	Is_dealer_admin          bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Orientation              int32
	Pages                    []int32
	Publish                  float64
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Run_from                 float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
	View_session             string
}

// Catalog_push_send is triggered when a notification is sent from the amazon server
type Catalog_push_send struct {
	Api_Build   string
	Api_Version string
	Api_env     string
	Catalogs    []int32
	Dealers     []int32
	Endpoint_id int32
	Push_type   string
	Timestamp   float64
	Type        string
	User_id     int32
	Uuid        string
}

// NOTE geo information is missing as of the 4/2015  but will be added
// NOTE the location/time is highly  client depended, on iOS a fetch is a user interaction.
// on android is not
type Catalog_push_fetch struct {
	Api_Build          string
	Api_Version        string
	Api_env            string
	App_id             int32
	App_groups         []int32
	Catalogs           []int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Dealers            []int32
	Is_archive_user    bool
	Is_dealer_admin    bool
	Push_type          string
	Query_string       string
	Server_addr        string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

// NOTE this event is triggered when data is requested not only when a catalog is viewed.
// DON't USE FOR ANY STATISTIC
type Catalog_view struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Catalog_id               int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Expires                  float64
	Is_archive_user          bool
	Is_dealer_admin          bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Publish                  float64
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Run_from                 float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Dealer_list struct {
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offset                   int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Dealer_unfavorite struct {
	Api_env            string
	App_id             int32
	App_groups         []int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Dealer_id          int32
	Is_dealer_admin    bool
	Is_archive_user    bool
	Is_uuid_ephemeral  bool
	Query_string       string
	Server_addr        string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

type Dealer_favorite struct {
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Dealer_view struct {
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offset                   int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Offer_View struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Catalog_id               int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Currency                 string
	Dealer_id                int32
	Expires                  float64
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Offer_id                 int32
	PrePrice                 float64
	Price                    float64
	Publish                  float64
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Run_from                 float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_birth_year          int64
	User_gender              string
	User_id                  int32
	Uuid                     string
}

type Offer_Click struct {
	// NOTE should we keep this trhee ?
	// price preprice currecny
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Catalog_id               int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Currency                 string
	Dealer_id                int32
	Expires                  float64
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Offer_id                 int32
	PrePrice                 float64
	Price                    float64
	Publish                  float64
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Run_from                 float64
	Server_addr              string
	Timestamp                float64
	Type                     string
	User_birth_year          int64
	User_gender              string
	User_id                  int32
	Uuid                     string
}

type Offer_list struct {
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offers                   []int32
	Offset                   int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Stores                   []int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Offer_Referral struct {
	Api_Version        string
	Api_env            string
	App_env            int32
	App_id             int32
	App_groups         []int32
	Buy_url            string
	Catalog_id         int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Dealer_id          int32
	Is_dealer_admin    bool
	Is_archive_user    bool
	Is_uuid_ephemeral  bool
	Offer_id           int32
	Query_string       string
	Server_addr        string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

type Offer_Search struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_user_defined_location bool
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offers                   []int32
	Offset                   int32
	Query                    string
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Results                  int32
	Server_addr              string
	Stores                   []int32
	Timestamp                float64
	Type                     string
	User_birth_year          int64
	User_gender              string
	User_id                  int32
	Uuid                     string
}

type Offer_Suggest struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Is_user_defined_location bool
	Limit                    int32
	Offers                   []int32
	Offset                   int32
	Query                    string
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Results                  int32
	Server_addr              string
	Stores                   []int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppingitem_Add struct {
	Api_build                string
	Api_env                  string
	Api_version              string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Descrption               string
	Is_archive_user          bool
	Is_dealer_admin          bool
	Is_owner                 bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offer_id                 int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Shopping_item            int32
	Shopping_list            int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppingitem_edit struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Description              string
	Is_user_defined_location bool
	Is_Owner                 bool
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offer_id                 int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Shopping_item            int32
	Shopping_list            int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppingitem_list struct {
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Description              string
	Is_user_defined_location bool
	Is_Owner                 bool
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Shopping_list            int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppingitem_tick struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Description              string
	Is_user_defined_location bool
	Is_Owner                 bool
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offer_id                 int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Shopping_item            int32
	Shopping_list            int32
	Timestamp                float64
	Tick                     bool
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppinglist_add struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Is_user_defined_location bool
	Is_Owner                 bool
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Shopping_list            int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Shoppingshare_accept struct {
	Api_Build          string
	Api_Version        string
	Api_env            string
	App_id             int32
	App_groups         []int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Query_string       string
	Server_addr        string
	Shopping_list      int32
	Timestamp          float64
	Type               string
	User_id            int32
}

type Shoppingshare_add struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Shopping_list            int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Store_click struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Store                    int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Store_directions struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Is_Dealer_admin          bool
	Is_archive_user          bool
	Is_user_defined_location bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Server_addr              string
	Store                    int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Store_list struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealers                  []int32
	Is_user_defined_location bool
	Is_archive_user          bool
	Is_Dealer_admin          bool
	Is_uuid_ephemeral        bool
	Limit                    int32
	Offer_id                 int32
	Offset                   int32
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Stores                   []int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type Store_view struct {
	Api_Build                string
	Api_Version              string
	Api_env                  string
	App_id                   int32
	App_groups               []int32
	Client_addr              string
	Client_app_version       string
	Client_user_agent        string
	Dealer_id                int32
	Is_user_defined_location bool
	Is_archive_user          bool
	Is_Dealer_admin          bool
	Is_uuid_ephemeral        bool
	Query_string             string
	Request_geohash          string
	Request_latitude         float64
	Request_longitude        float64
	Request_radius           float64
	Server_addr              string
	Store                    int32
	Timestamp                float64
	Type                     string
	User_id                  int32
	Uuid                     string
}

type User_create struct {
	Api_Build          string
	Api_Version        string
	Api_env            string
	App_id             int32
	App_groups         []int32
	Birthyear          int64
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Email              string
	User_gender        string
	Is_archive_user    bool
	Is_dealer_admin    bool
	Is_uuid_ephemeral  bool
	Locale             string
	Name               string
	Query_string       string
	Server_addr        string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

type User_login struct {
	Api_Build          string
	Api_Version        string
	Api_env            string
	App_id             int32
	App_groups         []int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Email              string
	Is_archive_user    bool
	Is_dealer_admin    bool
	Is_uuid_ephemeral  bool
	Provider           string
	Query_string       string
	Server_addr        string
	Session_type       string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

type User_logout struct {
	Api_Build          string
	Api_Version        string
	Api_env            string
	App_id             int32
	App_groups         []int32
	Client_addr        string
	Client_app_version string
	Client_user_agent  string
	Email              string
	Is_archive_user    bool
	Is_dealer_admin    bool
	Is_uuid_ephemeral  bool
	Query_string       string
	Server_addr        string
	Session_type       string
	Timestamp          float64
	Type               string
	User_id            int32
	Uuid               string
}

type User_pwd_reset struct {
	Api_Build         string
	Api_Version       string
	Api_env           string
	Client_addr       string
	Client_user_agent string
	Email             string
	Is_archive_user   bool
	Is_dealer_admin   bool
	Query_string      string
	Server_addr       string
	Timestamp         float64
	Type              string
	User_id           int32
}

type User_verify struct {
	Api_Build         string
	Api_Version       string
	Api_env           string
	Client_addr       string
	Client_user_agent string
	Email             string
	Is_archive_user   bool
	Is_dealer_admin   bool
	Query_string      string
	Server_addr       string
	Timestamp         float64
	Type              string
	User_id           int32
}
