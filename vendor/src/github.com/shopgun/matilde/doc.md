## Timestamps

ALL timestamps MUST be UNIX epoch timestamp as a float.

## Event sort order

Event blobs are sorted by timestamp ascending.


# Keys

## `timestamp` AND `timestampMs`

The timestamp of the event.

### Transform

    return float(timestampMs /1000 ) || float(timestamp)

new key is "timestamp"





## `ip`

The ip field is the ip of the device triggering the event.

Currently it is long encoded with the PHP ip2long.

### Transform

    return (ipv6 encoded ipv4 addresses) (as string)

the new key is "client_addr"




## `server_ip`

The ip of the server handling the request

### Transform

    return (ipv6 encoded ipv4 addresses) (as string)

The new key is "server_addr"




## `remote_addr`

String version of `ip` field.

Nuke.




## `publish`

Publish data of event content.
When the content is visible in apps.

### Transform

content_id  is sensitive:
i.e. catalog_id, offer_id etect

    if publish exist:
        return float(publish)
    else:
        return Database:Content(content_id).publish

new key is "publish"




## `date_*`

Deprecated. Nuke.


## `year` `quarter` `month` `week` `day` `hour` `minute` `second`

Deprecated. Nuke.



## `api_app_version`

The client application version identifier.

### Transform

    if api_app_version is undefined:
        return null
    else if len(api_app_version) < 1:
        return null
    else
        return api_app_version

The new key is "client_app_version"



## `run_from`

run_from data of event content.
When the content can be purchased in the stores.

### Transform

    if run_from exist:
        return float(run_from)
    else:
        return Database:Content(content_id).run_from

new key is "run_from"




## `expires`

expires data of event content.
When the content is no longer available

### Transform

    if expires exist:
        return float(expires)
    else:
        return Database:Content(content_id).expires

new key is "expires"

** for offers, need to save also the catalog, life time**


## `price`

The price of the content.
If currency field is undefined, currency is DKK.

### Transform

    if price exist:
        return float(price)
    else:
        return Database:Content(content_id).price

The new key is "price"




## `pre_price`

The pre_price of the content.
If currency field is undefined, currency is DKK.

### Transform

    if pre_price exist (can be null):
        return float(pre_price)
    else:
        return Database:Content(content_id).pre_price

The new key is "pre_price"



## `currency`

The currency field is the currency of price and pre_price
formatted as an ISO4217 string.
### strangeness
 2010 events have 1, as key to the table currency in the database
 and it is still dkk
 v2 events have null if is dkk
 vFuture will have strawberry if is dkk
### Transform

    if currency exists:
        return currency
    else:
        return "DKK"

The new key "currency"




## `latitude` + `longitude`

The desired location of the client device when the event happened.
Can be both string and number

### Transform

    // for lat
    return float(latitude) or null if not exist
    // for long
    return float(longitude) or null of not exist

The new keys are: "request_longitude" and "request_latitude"


## `geohash`

The full-precision geohash of events long + lat.

### Transform

    return geohash(longitude, latitude)

The new key is "request_geohash"



## `geocoded`

Int (bool) determines if the client is using a manual entered
address or a location service based address.

### Transform

    if 0
        return true
    else
        return false

new key "is_user_defined_location"



## `offer_view.catalog`

The internal catalog id. Can be null.

### Transform

    if exists:
        if null
            return null
        else:
            return int(catalog)
    else:
        return Database.Offer(offer).catalog

The new key is "catalog_id"



## `user_agent`

The user agent string of the client device.

### Transform

    if exists:
        return user_agent
    else:
        return null

new key is "client_user_agent"



## `scope`

Outer namespacing for event type.
See `action`


## `action`

sub-namespace for the event type.

### Transform

    if scope or actions doesnt exist:
        nuke
    else:
        return scope + "." + action

new key is "type"




## `device`

Nuke!



## `api_env`

The api environment name. eg. staging edge production


### Transform

    if api_env:
        return api_env
    else:
        return "production"

the new key is "api_env"



## `api`

The API external version identifier.

### Transform

    if api
        return api
    else:
        return null

the new key is "api_version"



## `api_build`

The API build identifier (e.g. "1.16.4")
This is NOT the same as `api` which is the external identifier

### Transform

    if api_build:
        return api_build
    else:
        return null

# new key is "api_build"




## `locationDetermined`

Timestamp of when the v1 client got the location.

Nuke


## `dealer`

The dealer id. it should be an interger.

### Transform

    return int(dealer)

The new key is "dealer_id"



## `is_dealer_admin`

Determines if the event was triggered by the contents admin.

### Transform

    if is_dealer_admin exists and is Truhty:
        return true
    else:
        return false

The new key is "is_dealer_admin"


## `is_archive_user`

Determines if the event was triggered by an archive user in an archive context.

### Transform

    if is_archive_user exists and is Truhty:
        return true
    else:
        return false

the new key is "is_archive_user"



## `activity`

Nuke



## `is_uuid_ephemeral`

Determines if the `uuid` field is a persistent UUID across session,
or simply a session identifier.

### Transform

    if is_uuid_ephemeral:
        return true
    else:
        return false

new key is "is_uuid_ephemeral"



## `query_string`

The HTTP request query string.

### Transform

    if query_string
        return query_string
    else
        return null

new key is "query_string"


## `distance`

The max distance a user desired to travel for content!

### Transform:

    if distance:
        return float(distance)
    else:
        return null

new key is "request_radius"



## `accuracy`

Nuke



## `user`

The user id of a user is logged in. Or null if not exists.

### Transform

    if user is null:
        return null
    else:
        return int(user)

The new key is "user_id"



## `birthYear`

The user's birth year

### Transform

    if user:
        return int(Database.User(user).birthYear)
    else:
        return null

New key is "user_birth_year"


## `gender`

The gender of the user.

### Transform

    if user is not null:
        return string(Database.User(user).gender)
    else:
        return null

new key is "user_gender"



## `uuid`

This is NOT an actual UUID!!!
This is a string that represents uniqueness.
Combine with is_uuid_ephemeral for increased accuracy.

### Transform

    if uuid:
        return uuid
    else:
        Nuke

The new key is "uuid"



## `app`

The eta app id.

### Transform

    return int(app)

the new key is "app_id"



