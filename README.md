### gobulenat

go bulenat is a login system

environment description:

docker-compose two instances: web app (gobulenat) and database (postgresql)

web app written in go and serves html templates (go package)

#### database

users table which will store attribute for authentication

id, first_name, last_name, email, password, session, auth_method, 

#### web app pages

1. login page `/` 

2. signup page `/signup`

3. logout button (button redirecting to `/`)

4. profile page `/profile/{id}` (redirecting by user id)

###### login

form based email and password (set cookie and authenticated)

###### signup

first and last name, email and password (redirecting back to `/`)

###### logout

will redirect to login page (set no cookie)

###### profile page

will just say hello to user after login, using user struct (`fname` and `lname`)

### roadmap

- login form on `/` and redirect from all pages unless have a token

- set auth methhod to read email\password from database `/`

- say hello to user from home page `/profile/${user}`

