package web

type ToSEO struct {
	Channel string `query:"channel" valid:"required,oneof=channel"`
	ID      string `query:"id" valid:"required,max=64"`
}
