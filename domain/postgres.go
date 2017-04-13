package domain

import (
	"strings"
	"time"
	"github.com/twinj/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        string  `sql:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Product struct {
	BaseModel
	Code 				string 				`sql:"unique;index" json:"code"`
	Name 				string 				`json:"name"`
	Version 			string 				`json:"version"`
	Routes	 			[]Route		       		`gorm:"ForeignKey:ReferrerID;AssociationForeignKey:ID" json:"routes"`
}

func (product *Product) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type Route struct {
	BaseModel
	Code				string 				`gorm:"unique_index:idx_route_code" json:"code"`
	ReferrerID			string				`gorm:"unique_index:idx_route_code"`
	ListenPath 			string 				`json:"listen_path"`
	Verb 				string 				`json:"verb"`
	ServiceName 			string 				`json:"service_name"`
	Handlers 			[]string 			`sql:"-" json:"handlers"`
	Roles	 			[]string 			`sql:"-" json:"roles"`
	Tags	 			[]string 			`sql:"-" json:"tags"`
	HandlersSerialized 		string
	RolesSerialized			string
	TagsSerialized			string
	InjectData			[]Inject			`gorm:"ForeignKey:ReferrerID;AssociationForeignKey:ID" json:"inject_data"`
	InjectGlobalData		bool				`json:"inject_global_data"`
	Asset				string				`json:"asset"`
	AssetId				string				`json:"asset_id"`
	Alias				string				`json:"alias"`
}

func (route *Route) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func (route *Route) BeforeSave() (err error) {
	if route.Handlers != nil {
		route.HandlersSerialized = strings.Join(route.Handlers, ";")
	}
	if route.Roles != nil {
		route.RolesSerialized = strings.Join(route.Roles, ";")
	}
	if route.Tags != nil {
		route.TagsSerialized = strings.Join(route.Tags, ";")
	}
	return
}

func (route *Route) AfterFind() (err error) {
	if route.HandlersSerialized != "" {
		route.Handlers = strings.Split(route.HandlersSerialized, ";")
	}
	if route.RolesSerialized != "" {
		route.Roles = strings.Split(route.RolesSerialized, ";")
	}
	if route.TagsSerialized != "" {
		route.Tags = strings.Split(route.TagsSerialized, ";")
	}
	return
}

type Inject struct {
	BaseModel
	ReferrerID			string				`gorm:"unique_index:idx_inject_code"`
	Code				string 				`gorm:"unique_index:idx_inject_code" json:"code"`
	Where 				string 				`json:"where"`
	Value				string 				`json:"value"`
}

func (inject *Inject) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type Client struct {
	BaseModel
	Code				string 				`sql:"unique;index" json:"code"`
	ApiPath				string 				`sql:"unique;index" json:"api_path"`
	Product				string				`json:"product"`
	ProductVersion			string				`json:"product_version"`
	Client				string 				`json:"client"`
	RemoveApiPath			bool				`json:"remove_api_path"`
	Version 			string				`json:"version"`
	InjectData			[]Inject			`gorm:"ForeignKey:ReferrerID;AssociationForeignKey:ID" json:"global_inject_data"`
	Routes	 			[]Route 			`gorm:"ForeignKey:ReferrerID;AssociationForeignKey:ID" json:"routes"`
}

func (client *Client) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("ApiPath", strings.Split(uuid.NewV4().String(), "-")[0])
	return nil
}
