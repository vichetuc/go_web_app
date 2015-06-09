#Golang webapp example

It contains a few thoughts about my Golang studies, such as challenges and cool stuff I'm doing =)

##Challenges:

####1. GOPATH Conflicts 

The first version of the app was the implementation of the Google App Engine tutorial, it was a simple app with all the logic in the same file. The first challenge was separating it's logic in a kind of MVC structure. For some reason I didn't understand quite right yet, Google App Engine copies all source code to where app.yaml is and it generates a conflict.

######Problem:
``go-app-builder: Failed parsing input: app file <file_name>.go conflicts with same file imported from GOPATH``


######Solution:
The solution I found was to add the app.yaml file and an unic file that initializes the app in a single ``init()`` block. Then, separate the MVC structure in a /api directory

####2. HTML file parse

I took a lot of time to figure out how to parse an extracted HTML file to use in one of the routes/controller.
The solution was found in this Stack Overflow issue: http://stackoverflow.com/questions/17206467/go-how-to-render-multiple-templates-in-golang.


##Cool stuff:
nothing to put here yet =P
