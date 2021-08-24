package Entity

import "luigi.vanacore/go-snake/Core/Component"

type Entity struct {
	id         uint64
	tags       []string
	components []Component.Component
}
