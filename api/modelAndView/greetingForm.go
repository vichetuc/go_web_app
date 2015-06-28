package modelAndView

type GreetingAsForm struct {
	Author  string	  `form:"author" binding:"required"`
	Content string    `form:"content" binding:"required"`
}

