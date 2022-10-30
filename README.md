# final-project-backend

## Project Description

A RESTful API of House Rental using Golang and PostgreSQL that allow certain actions such as:
- Register: User can register into the system which requires several input. The system will then validate the user's email if it has been used.
- Login: User can login into the system which requires the user's email and password.
- Get All Cities: Fetching all cities
- Get User Profile by User ID: Fetching a certain user profile by user ID
- Get All Houses: Fetching all houses
- Get House by House ID: Fetching a certain house by house ID
- Get All House Statuses: Fetching all house statuses
- Get House's Reservation by House ID: Fetching all reservation of a house by house ID 
- Get All Pickup Statuses: Fetching all pickup statuses
- Get All Pickup Details: Fetching all pickup details

Features below require authorization where user must login first:
- Get User Profile: Fetching the current logged in user's profile
- Update User Profile: User can update their profile
- Update User Role to Host: User can become a host
- Get Reservations by User: Fetching all user's reservations
- Top Up: User can top up their wallet
- Get Host Houses: Fetching all current logged in host's houses
- Get House Reservations by Host: Fetching all host's reservations
- Add New House: Host can add new house
- Update House: Host can update their houses
- Delete House: Host can delete their houses
- Add House Photo: Host can add photos of their houses
- Delete House Photo: Host can delete photos of their houses, but each house are required to have one photo
- Get Reservation Details by Reservation ID: Host and guest of the reservation are able to view the details of the reservation
- Add Reservation: User are able to make a reservation on a house
- Get Pickup Details by Reservation ID: Fetching pickup details by reservation ID
- Request Pickup: User are able to request pickup for an upcoming reservation
- Get Game Data: Fetching a user's game data
- Play Game: Play a guessing coin flip game where user will get Rp 50.000 if guessed correctly, and get Rp 10.000 if guessed incorrectly
- Update Pickup Status: Admin is able to update user's pickup requests to the appropriate status

## Entity Relationship Diagram (ERD)

<img src="./assets/ERD House Rental.jpg" alt="erd-house-rental" width="1000"/>

## Test coverage screenshots
<img src="./assets/Test Coverage.png" alt="test-coverage-overall" width="1000"/>
