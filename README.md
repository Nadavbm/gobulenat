### gobulenat

go bulenat is a login system I create during studies

lets see where it goes.. this readme file will be for planning

login system checklist:

this app will be composed from a db and a web app.

webapp written in go and database will be postgresql

#### database

users table which will store attribute for authentication

id, first_name, last_name, email, password, session, auth_method, 

#### web app pages

1. opening page 

2. signup page

3. login page

4. logout button

4. profile page

###### opening page

intro gobulenat - nothing but redirecting to login or signup pages

###### login

form based email and password

###### signup

first and last name, email and password

###### logout

will redirect to home page

###### profile page

will just say hello to user after login (kind of a home page)

#### specs

docker-compose file will launch web and posgresql

### roadmap

- set signup form and write to database `/signup` 

- write signup forms to database

- login form on `/` and redirect from all pages unless have a token

- set auth methhod to read email\password from database `/`

- say hello to user from home page `/profile/${user}`

- set style css for header (done)

### isues

solve db migration problem