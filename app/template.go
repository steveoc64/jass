package main

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddTemplate("main-page", `<div class="container">
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sale-items", `Here is a big list of all the things

<ul>
	
{{range $key,$value := .}}
<li><img src="{{$value.Image}}"> {{$value.Name}} {{$value.Descr}} {{$value.SKU}} {{$value.Price}}</li>


{{end}}

</ul>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("slidemenu", `<!-- Slide in menu once logged in  -->
<nav class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-right" id="slidemenu">
  <a href="#" id="menu-shop"><i class="fa fa-facebook"></i> Shop</a>
  <a href="#" id="menu-discover"><i class="fa fa-facebook"></i> Discover</a>
  <a href="#" id="menu-merchandise"><i class="fa fa-facebook"></i> Merchandise</a>
  <a href="#" id="menu-facebook"><i class="fa fa-facebook"></i> Facebook</a>
  <a href="#" id="menu-instagram"><i class="fa fa-facebook"></i> Instagram</a>
  <a href="#" id="menu-blog"><i class="fa fa-facebook"></i> Blog</a>
  <a href="#" id="menu-contact"><i class="fa fa-facebook"></i> Contact</a>
</nav> 



`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
