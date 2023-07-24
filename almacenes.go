package baseapp

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Almacen struct {
	Id                 int       `json:"_" db:"id"`
	UUID               uuid.UUID `json:"id_uuid" db:"id_uuid" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	NameAlmacen        string    `json:"nombre_stop" db:"nombre_stop" binding:"required"`
	Abrevia            string    `json:"abrevia" db:"abrevia" binding:"required"`
	Ubication          string    `json:"ubication" db:"ubication" binding:"required"`
	Phone              string    `json:"phone" db:"phone" binding:"required"`
	EncargadoPrincipal string    `json:"encargado" db:"encargado_principal" binding:"required"`
	PhoneEncargado     string    `json:"phone_encargado" db:"phone_encargado" binding:"required"`
	UriImg             string    `json:"uri_img" db:"uri_img" binding:"required"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	Status             string    `json:"status" db:"status"`
}

type ResponseAlmacen struct {
	UUID               uuid.UUID `json:"id_uuid" db:"id_uuid" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	NameAlmacen        string    `json:"nombre_stop" db:"nombre_stop"`
	Abrevia            string    `json:"abrevia" db:"abrevia"`
	Ubication          string    `json:"ubication" db:"ubication" `
	Phone              string    `json:"phone" db:"phone" `
	EncargadoPrincipal string    `json:"encargado" db:"encargado_principal" `
	PhoneEncargado     string    `json:"phone_encargado" db:"phone_encargado"`
	UriImg             string    `json:"uri_img" db:"uri_img" `
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	Status             string    `json:"status" db:"status"`
}

func (i Almacen) UpperCase() *Almacen {
	return &Almacen{
		NameAlmacen:        strings.ToUpper(i.NameAlmacen),
		Abrevia:            strings.ToUpper(i.Abrevia),
		Ubication:          strings.ToUpper(i.Ubication),
		Phone:              strings.ToUpper(i.Phone),
		EncargadoPrincipal: strings.ToUpper(i.EncargadoPrincipal),
		PhoneEncargado:     strings.ToUpper(i.PhoneEncargado),
		UriImg:             i.UriImg,
	}
}
