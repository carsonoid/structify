module github.com/carsonoid/structify

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/sanity-io/litter v1.2.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.0.0-20191109101513-0171b7c15da1
	k8s.io/client-go v0.0.0-20190819141724-e14f31a72a77 // indirect
)

replace github.com/sanity-io/litter => ./litter
