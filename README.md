# EchoRestAPI

Rest Apu on Echo framework
Using User model with fields: ID and Name

Get all users:
Returns all users
```
GET /users
```
___
Get user by ID:
Returns user's body
```
GET /users/"id"
```
___
Create new user:
Returns user ID
```
POST /users/"id"
{name: "Ilya"}
```
___
Update info about user:
Returns updated user ID
```
PUT /users/"id"
{name: "newName"}
```
___
Delete user by ID
Returns removed user ID
```
DELETE /users/"id"
```
