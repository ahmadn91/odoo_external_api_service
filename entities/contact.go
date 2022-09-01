package entities

import "time"

type Contact struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Title       string `json:"title"`
	CreateDate  time.Time `json:"createDate"`
	Active		bool		`json:"active"`
	Email 		string		`json:"email"`
	Phone 		string 		`json:"phone"`

}

// id                           | integer                     |           | not null | nextval('res_partner_id_seq'::regclass)
// name                         | character varying           |           |          |
// company_id                   | integer                     |           |          |
// create_date                  | timestamp without time zone |           |          |
// display_name                 | character varying           |           |          |
// date                         | date                        |           |          |
// title                        | integer                     |           |          |
// parent_id                    | integer                     |           |          |
// ref                          | character varying           |           |          |
// lang                         | character varying           |           |          |
// tz                           | character varying           |           |          |
// user_id                      | integer                     |           |          |
// vat                          | character varying           |           |          |
// website                      | character varying           |           |          |
// comment                      | text                        |           |          |
// credit_limit                 | double precision            |           |          |
// active                       | boolean                     |           |          |
// employee                     | boolean                     |           |          |
// function                     | character varying           |           |          |
// type                         | character varying           |           |          |
// street                       | character varying           |           |          |
// street2                      | character varying           |           |          |
// zip                          | character varying           |           |          |
// city                         | character varying           |           |          |
// state_id                     | integer                     |           |          |
// country_id                   | integer                     |           |          |
// partner_latitude             | numeric                     |           |          |
// partner_longitude            | numeric                     |           |          |
// email                        | character varying           |           |          |
// phone                        | character varying           |           |          |
// mobile                       | character varying           |           |          |
// is_company                   | boolean                     |           |          |
// industry_id                  | integer                     |           |          |
// color                        | integer                     |           |          |
// partner_share                | boolean                     |           |          |
// commercial_partner_id        | integer                     |           |          |
// commercial_company_name      | character varying           |           |          |
// company_name                 | character varying           |           |          |
// create_uid                   | integer                     |           |          |
// write_uid                    | integer                     |           |          |
// write_date                   | timestamp without time zone |           |          |
// date_localization            | date                        |           |          |
// message_main_attachment_id   | integer                     |           |          |
// email_normalized             | character varying           |           |          |
// message_bounce               | integer                     |           |          |
// contact_address_complete     | character varying           |           |          |
// signup_token                 | character varying           |           |          |
// signup_type                  | character varying           |           |          |
// signup_expiration            | timestamp without time zone |           |          |
// plan_to_change_car           | boolean                     |           |          |
// plan_to_change_bike          | boolean                     |           |          |
// team_id                      | integer                     |           |          |
// partner_gid                  | integer                     |           |          |
// additional_info              | character varying           |           |          |
// phone_sanitized              | character varying           |           |          |
// debit_limit                  | numeric                     |           |          |
// last_time_entries_checked    | timestamp without time zone |           |          |
// invoice_warn                 | character varying           |           |          |
// invoice_warn_msg             | text                        |           |          |
// sale_warn                    | character varying           |           |          |
//  sale_warn_msg                | text                        |           |          |
//  zone_id                      | integer                     |           |          |
//  cluster_id                   | integer                     |           |          |
//  district_id                  | integer                     |           |          |
//  map_widget_field             | character varying           |           |          |
//  location_assigned_zone_false | boolean                     |           |          |
//  ocn_token                    | character varying           |           |          |
//  online_partner_information   | character varying           |           |          |