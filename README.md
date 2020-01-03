# IMS API Specs
## Functions
- add piece of equipment to inventory
  - requires ADMIN privs
- remove piece of equipment to inventory
  - requires ADMIN privs
- set posessor of equipment
  - stored (in storage)
  - or UID of user that has it
- return equipment (set possesor to stored)
## Properties of things
- Equipment
  - UID
  - Name (Ex: Canon D700)
  - part # (Ex: #2)
    - combined w/name to form full item showed to user ("Canon D700 #2")
- User
  - UID
  - Name (first and/or last)
  - School Email
  - password hash
  - email used as username for login
  - password is hashed by js to be sent over 
